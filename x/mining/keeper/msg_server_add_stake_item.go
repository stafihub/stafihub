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
	stakePool, found := k.Keeper.GetStakePool(ctx, msg.StakePoolIndex)
	if !found {
		return nil, types.ErrStakePoolNotExist
	}

	if !(k.sudoKeeper.IsAdmin(ctx, msg.Creator) || k.Keeper.HasMiningProvider(ctx, user) && msg.Creator == stakePool.Creator) {
		return nil, types.ErrUpdateStakeItemPermissionDeny
	}
	willUseIndex := k.GetStakeItemNextIndex(ctx, msg.StakePoolIndex)

	maxStakeItemNumber := k.Keeper.GetMaxStakeItemNumber(ctx)
	if willUseIndex >= maxStakeItemNumber {
		return nil, types.ErrStakeItemNumberReachLimit
	}

	stakeItemLimit := k.Keeper.GetStakeItemLimit(ctx)
	if msg.LockSecond > stakeItemLimit.MaxLockSecond {
		return nil, types.ErrStakeItemEraSecondExceedLimit
	}
	if msg.PowerRewardRate.GT(stakeItemLimit.MaxPowerRewardRate) {
		return nil, types.ErrStakeItemPowerRewardRateExceedLimit
	}

	// check reward second and lock second
	for _, rewardPool := range stakePool.RewardPools {
		if rewardPool.TotalRewardAmount.Quo(rewardPool.RewardPerSecond).LT(sdk.NewIntFromUint64(msg.LockSecond)) {
			return nil, types.ErrRewardSecondsLessThanMaxLockSeconds
		}
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
