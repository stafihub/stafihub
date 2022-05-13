package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/mining/types"
)

func (k msgServer) AddRewardPool(goCtx context.Context, msg *types.MsgAddRewardPool) (*types.MsgAddRewardPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	stakePool, found := k.Keeper.GetStakePool(ctx, msg.StakeTokenDenom)
	if !found {
		return nil, types.ErrStakePoolAlreadyExist
	}

	stakePool.RewardPools = append(stakePool.RewardPools, &types.RewardPool{
		RewardTokenDenom:  msg.RewardTokenDenom,
		TotalRewardAmount: msg.TotalRewardAmount,
		LeftRewardAmount:  msg.TotalRewardAmount,
		RewardPerSecond:   msg.RewardPerSecond,
		StartTimestamp:    msg.StartTimestamp,
		RewardPerPower:    sdk.ZeroInt(),
	})

	k.Keeper.SetStakePool(ctx, msg.StakeTokenDenom, stakePool)

	return &types.MsgAddRewardPoolResponse{}, nil
}
