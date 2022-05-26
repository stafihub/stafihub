package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/mining/types"
)

func (k msgServer) EmergencyWithdraw(goCtx context.Context, msg *types.MsgEmergencyWithdraw) (*types.MsgEmergencyWithdrawResponse, error) {
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
	if !stakePool.EmergencySwitch {
		return nil, types.ErrEmergencySwitchClose
	}

	stakePool.TotalStakedAmount = stakePool.TotalStakedAmount.Sub(userStakeRecord.StakedAmount)
	if stakePool.TotalStakedAmount.IsNegative() {
		stakePool.TotalStakedAmount = sdk.ZeroInt()
	}
	stakePool.TotalStakedPower = stakePool.TotalStakedPower.Sub(userStakeRecord.StakedPower)
	if stakePool.TotalStakedPower.IsNegative() {
		stakePool.TotalStakedPower = sdk.ZeroInt()
	}

	userStakeRecord.StakedAmount = sdk.ZeroInt()
	userStakeRecord.StakedPower = sdk.ZeroInt()

	withdrawTokens := sdk.NewCoins(sdk.NewCoin(stakePool.StakeTokenDenom, userStakeRecord.StakedAmount))
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recipientAddr, withdrawTokens); err != nil {
		return nil, err
	}

	k.SetStakePool(ctx, stakePool)
	k.SetUserStakeRecord(ctx, userStakeRecord)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeEmergencyWithdraw,
			sdk.NewAttribute(types.AttributeKeyAccount, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyStakePoolIndex, fmt.Sprintf("%d", msg.StakePoolIndex)),
			sdk.NewAttribute(types.AttributeKeyStakeRecordIndex, fmt.Sprintf("%d", msg.StakeRecordIndex)),
			sdk.NewAttribute(types.AttributeKeyWithdrawAmount, userStakeRecord.StakedAmount.String()),
		),
	)

	return &types.MsgEmergencyWithdrawResponse{}, nil
}
