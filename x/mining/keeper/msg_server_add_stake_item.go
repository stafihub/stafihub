package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/mining/types"
)

func (k msgServer) AddStakeItem(goCtx context.Context, msg *types.MsgAddStakeItem) (*types.MsgAddStakeItemResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	user, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) && !k.Keeper.HasMiningProvider(ctx, user) {
		return nil, types.ErrUserNotAdminOrMiningProvider
	}
	willUseIndex := k.GetStakeItemNextIndex(ctx, msg.StakePoolIndex)

	maxStakeItemNumber := k.Keeper.GetMaxStakeItemNumber(ctx)
	if willUseIndex >= maxStakeItemNumber {
		return nil, types.ErrStakeItemNumberReachLimit
	}

	stakeItem := types.StakeItem{
		Enable:          msg.Enable,
		Index:           willUseIndex,
		LockSecond:      msg.LockSecond,
		PowerRewardRate: msg.PowerRewardRate,
	}

	k.Keeper.SetStakeItemIndex(ctx, msg.StakePoolIndex, willUseIndex)
	k.Keeper.SetStakeItem(ctx, &stakeItem)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAddStakeItem,
			sdk.NewAttribute(types.AttributeKeyAccount, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyStakePoolIndex, fmt.Sprintf("%d", msg.StakePoolIndex)),
			sdk.NewAttribute(types.AttributeKeyStakeItemIndex, fmt.Sprintf("%d", willUseIndex)),
			sdk.NewAttribute(types.AttributeKeyLockSecond, fmt.Sprintf("%d", msg.LockSecond)),
			sdk.NewAttribute(types.AttributeKeyPowerRewardRate, msg.PowerRewardRate.String()),
		),
	)
	return &types.MsgAddStakeItemResponse{}, nil
}
