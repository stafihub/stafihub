package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/mining/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) AddRewardPool(goCtx context.Context, msg *types.MsgAddRewardPool) (*types.MsgAddRewardPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}
	stakePool, found := k.Keeper.GetStakePool(ctx, msg.StakeTokenDenom)
	if !found {
		return nil, types.ErrStakePoolAlreadyExist
	}

	if len(stakePool.RewardPools) >= int(stakePool.MaxRewardPools) {
		return nil, types.ErrRewardPoolNumberReachLimit
	}
	curBlockTime := uint64(ctx.BlockTime().Unix())

	willUseIndex := k.Keeper.GetRewardPoolNextIndex(ctx, msg.StakeTokenDenom)
	willUseLastRewardTimestamp := msg.StartTimestamp
	if msg.StartTimestamp < curBlockTime {
		willUseLastRewardTimestamp = curBlockTime
	}

	stakePool.RewardPools = append(stakePool.RewardPools, &types.RewardPool{
		Index:               willUseIndex,
		RewardTokenDenom:    msg.RewardTokenDenom,
		TotalRewardAmount:   msg.TotalRewardAmount,
		LeftRewardAmount:    msg.TotalRewardAmount,
		RewardPerSecond:     msg.RewardPerSecond,
		StartTimestamp:      msg.StartTimestamp,
		RewardPerPower:      sdk.ZeroInt(),
		LastRewardTimestamp: willUseLastRewardTimestamp,
	})

	k.Keeper.SetRewardPoolIndex(ctx, msg.StakeTokenDenom, willUseIndex)
	k.Keeper.SetStakePool(ctx, stakePool)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAddRewardPool,
			sdk.NewAttribute(types.AttributeKeyAccount, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyRewardTokenDenom, msg.RewardTokenDenom),
			sdk.NewAttribute(types.AttributeKeyTotalRewardAmount, msg.TotalRewardAmount.String()),
			sdk.NewAttribute(types.AttributeKeyRewardPerSecond, msg.RewardPerSecond.String()),
			sdk.NewAttribute(types.AttributeKeyStartTimestamp, fmt.Sprintf("%d", msg.StartTimestamp)),
			sdk.NewAttribute(types.AttributeKeyLastRewardTimestamp, fmt.Sprintf("%d", willUseLastRewardTimestamp)),
		),
	)

	return &types.MsgAddRewardPoolResponse{}, nil
}
