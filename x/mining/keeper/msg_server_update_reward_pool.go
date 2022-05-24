package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/mining/types"
)

func (k msgServer) UpdateRewardPool(goCtx context.Context, msg *types.MsgUpdateRewardPool) (*types.MsgUpdateRewardPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	user, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	stakePool, found := k.Keeper.GetStakePool(ctx, msg.StakePoolIndex)
	if !found {
		return nil, types.ErrStakePoolNotExist
	}

	var willUseRewardPool *types.RewardPool
	for _, rewardPool := range stakePool.RewardPools {
		if rewardPool.Index == msg.RewardPoolIndex {
			willUseRewardPool = rewardPool
		}
	}
	if willUseRewardPool == nil {
		return nil, types.ErrRewardPoolNotExist
	}

	maxRewardPoolNumber := k.Keeper.GetMaxRewardPoolNumber(ctx)
	if len(stakePool.RewardPools) >= int(maxRewardPoolNumber) {
		return nil, types.ErrRewardPoolNumberReachLimit
	}

	rewardToken, found := k.Keeper.GetRewardToken(ctx, willUseRewardPool.RewardTokenDenom)
	if !found {
		return nil, types.ErrRewardTokenNotSupport
	}
	if msg.NewRewardAmount.LT(rewardToken.MinTotalRewardAmount) {
		return nil, types.ErrTotalRewardAmountLessThanLimit
	}

	rewardTokens := sdk.NewCoins(sdk.NewCoin(willUseRewardPool.RewardTokenDenom, msg.NewRewardAmount))
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, user, types.ModuleName, rewardTokens); err != nil {
		return nil, err
	}

	willUseRewardPool.TotalRewardAmount = willUseRewardPool.TotalRewardAmount.Add(msg.NewRewardAmount)
	if willUseRewardPool.LeftRewardAmount.IsZero() {
		willUseRewardPool.LastRewardTimestamp = uint64(ctx.BlockTime().Unix())
	}
	willUseRewardPool.LeftRewardAmount = willUseRewardPool.LeftRewardAmount.Add(msg.NewRewardAmount)

	k.Keeper.SetStakePool(ctx, stakePool)

	return &types.MsgUpdateRewardPoolResponse{}, nil
}
