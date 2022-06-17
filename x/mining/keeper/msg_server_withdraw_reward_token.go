package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/mining/types"
)

func (k msgServer) WithdrawRewardToken(goCtx context.Context, msg *types.MsgWithdrawRewardToken) (*types.MsgWithdrawRewardTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	user, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	stakePool, found := k.Keeper.GetStakePool(ctx, msg.StakePoolIndex)
	if !found {
		return nil, types.ErrStakePoolNotExist
	}

	// find reward pool
	var willUseRewardPool *types.RewardPool
	for _, rewardPool := range stakePool.RewardPools {
		if rewardPool.Index == msg.RewardPoolIndex {
			willUseRewardPool = rewardPool
		}
	}
	if willUseRewardPool == nil {
		return nil, types.ErrRewardPoolNotExist
	}
	// check permission
	if willUseRewardPool.Creator != msg.Creator {
		return nil, types.ErrWithdrawRewardTokenPermissionDeny
	}

	//check amount
	curBlockTime := uint64(ctx.BlockTime().Unix())
	updateStakePool(stakePool, curBlockTime)

	rewardToken, found := k.Keeper.GetRewardToken(ctx, willUseRewardPool.RewardTokenDenom)
	if !found {
		return nil, types.ErrRewardTokenNotSupport
	}

	if willUseRewardPool.LeftRewardAmount.LT(msg.WithdrawAmount) ||
		willUseRewardPool.TotalRewardAmount.Sub(msg.WithdrawAmount).LT(rewardToken.MinTotalRewardAmount) {
		return nil, types.ErrWithdrawRewardTokenAmountTooLarge
	}

	willUseRewardPool.TotalRewardAmount = willUseRewardPool.TotalRewardAmount.Sub(msg.WithdrawAmount)
	willUseRewardPool.LeftRewardAmount = willUseRewardPool.LeftRewardAmount.Sub(msg.WithdrawAmount)

	willWithdrawCoin := sdk.NewCoins(sdk.NewCoin(willUseRewardPool.RewardTokenDenom, msg.WithdrawAmount))
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, user, willWithdrawCoin); err != nil {
		return nil, err
	}

	k.Keeper.SetStakePool(ctx, stakePool)

	return &types.MsgWithdrawRewardTokenResponse{}, nil
}
