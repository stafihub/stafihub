package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/mining/types"
)

func (k msgServer) AddStakePool(goCtx context.Context, msg *types.MsgAddStakePool) (*types.MsgAddStakePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	user, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	if k.GetMiningProviderSwitch(ctx) && !k.HasMiningProvider(ctx, user) {
		return nil, types.ErrUserNotAdminOrMiningProvider
	}
	curBlockTime := uint64(ctx.BlockTime().Unix())
	willUseStakePoolIndex := k.Keeper.GetStakePoolNextIndex(ctx)

	rewardPools := make([]*types.RewardPool, 0)
	for i, rewardPool := range msg.RewardPoolInfoList {
		rewardToken, found := k.Keeper.GetRewardToken(ctx, rewardPool.RewardTokenDenom)
		if !found {
			return nil, types.ErrRewardTokenNotSupport
		}
		if rewardPool.TotalRewardAmount.LT(rewardToken.MinTotalRewardAmount) {
			return nil, types.ErrTotalRewardAmountLessThanLimit
		}

		willUseLastRewardTimestamp := rewardPool.StartTimestamp
		if rewardPool.StartTimestamp < curBlockTime {
			willUseLastRewardTimestamp = curBlockTime
		}

		rewardTokens := sdk.NewCoins(sdk.NewCoin(rewardPool.RewardTokenDenom, rewardPool.TotalRewardAmount))
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, user, types.ModuleName, rewardTokens); err != nil {
			return nil, err
		}

		rewardPool := types.RewardPool{
			Index:               uint32(i),
			RewardTokenDenom:    rewardPool.RewardTokenDenom,
			TotalRewardAmount:   rewardPool.TotalRewardAmount,
			LeftRewardAmount:    rewardPool.TotalRewardAmount,
			RewardPerSecond:     rewardPool.RewardPerSecond,
			StartTimestamp:      rewardPool.StartTimestamp,
			RewardPerPower:      sdk.ZeroInt(),
			LastRewardTimestamp: willUseLastRewardTimestamp,
		}

		rewardPools = append(rewardPools, &rewardPool)
	}

	stakePool := types.StakePool{
		Index:             willUseStakePoolIndex,
		StakeTokenDenom:   msg.StakeTokenDenom,
		RewardPools:       rewardPools,
		TotalStakedAmount: sdk.ZeroInt(),
		TotalStakedPower:  sdk.ZeroInt(),
	}

	for i, stakeItemInfo := range msg.StakeItemInfoList {
		stakeItem := types.StakeItem{
			Index:           uint32(i),
			StakePoolIndex:  willUseStakePoolIndex,
			LockSecond:      stakeItemInfo.LockSecond,
			PowerRewardRate: stakeItemInfo.PowerRewardRate,
			Enable:          true,
		}
		k.Keeper.SetStakeItem(ctx, &stakeItem)
	}
	k.Keeper.SetStakeItemIndex(ctx, willUseStakePoolIndex, uint32(len(msg.StakeItemInfoList)-1))

	k.SetStakePool(ctx, &stakePool)
	k.Keeper.SetStakePoolIndex(ctx, willUseStakePoolIndex)
	k.Keeper.SetRewardPoolIndex(ctx, willUseStakePoolIndex, uint32(len(rewardPools)-1))

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAddStakePool,
			sdk.NewAttribute(types.AttributeKeyAccount, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyStakeTokenDenom, msg.StakeTokenDenom),
			sdk.NewAttribute(types.AttributeKeyStakePoolIndex, fmt.Sprintf("%d", willUseStakePoolIndex)),
		),
	)

	return &types.MsgAddStakePoolResponse{}, nil
}
