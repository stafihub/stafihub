package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/mining/types"
)

func (k msgServer) ClaimReward(goCtx context.Context, msg *types.MsgClaimReward) (*types.MsgClaimRewardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	recipientAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	userStakeRecord, found := k.Keeper.GetUserStakeRecord(ctx, msg.Creator, msg.StakePoolIndex, msg.StakeRecordIndex)
	if !found {
		return nil, types.ErrUserStakeRecordNotExist
	}
	stakePool, found := k.Keeper.GetStakePool(ctx, msg.StakePoolIndex)
	if !found {
		return nil, types.ErrStakePoolNotExist
	}
	if stakePool.EmergencySwitch {
		return nil, types.ErrEmergencySwitchOpen
	}
	curBlockTime := uint64(ctx.BlockTime().Unix())

	updateStakePool(stakePool, curBlockTime)

	willClaimCoins := calRewardTokens(stakePool, userStakeRecord)
	setNewRewardDebt(stakePool, userStakeRecord)

	if willClaimCoins.IsAllPositive() {
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recipientAddr, willClaimCoins); err != nil {
			return nil, err
		}
	}

	k.SetStakePool(ctx, stakePool)
	k.SetUserStakeRecord(ctx, userStakeRecord)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeClaimReward,
			sdk.NewAttribute(types.AttributeKeyAccount, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyStakePoolIndex, fmt.Sprintf("%d", msg.StakePoolIndex)),
			sdk.NewAttribute(types.AttributeKeyClaimedTokens, willClaimCoins.String()),
			sdk.NewAttribute(types.AttributeKeyStakeRecordIndex, fmt.Sprintf("%d", msg.StakeRecordIndex)),
		),
	)

	return &types.MsgClaimRewardResponse{}, nil
}
