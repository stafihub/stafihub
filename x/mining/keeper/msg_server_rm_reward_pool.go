package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/mining/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) RmRewardPool(goCtx context.Context, msg *types.MsgRmRewardPool) (*types.MsgRmRewardPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}
	stakePool, found := k.Keeper.GetStakePool(ctx, msg.StakePoolIndex)
	if !found {
		return nil, types.ErrStakePoolNotExist
	}

	var willUseRewardPool *types.RewardPool
	newRewardPool := make([]*types.RewardPool, 0)
	for _, rewardPool := range stakePool.RewardPools {
		if rewardPool.Index == msg.RewardPoolIndex {
			willUseRewardPool = rewardPool
		} else {
			newRewardPool = append(newRewardPool, rewardPool)
		}
	}
	if willUseRewardPool == nil {
		return nil, types.ErrRewardPoolNotExist
	}
	if willUseRewardPool.LeftRewardAmount.IsPositive() {
		return nil, types.ErrRewardPoolLeftRewardAmountNotZero
	}

	stakePool.RewardPools = newRewardPool
	k.Keeper.SetStakePool(ctx, stakePool)

	return &types.MsgRmRewardPoolResponse{}, nil
}
