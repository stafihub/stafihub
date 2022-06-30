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
	// check provider
	if k.GetMiningProviderSwitch(ctx) && !k.HasMiningProvider(ctx, user) {
		return nil, types.ErrUserNotMiningProvider
	}

	// check stake token denom
	if !k.rDexKeeper.IsRDexLpToken(ctx, msg.StakeTokenDenom) && !k.Keeper.HasStakeToken(ctx, msg.StakeTokenDenom) {
		return nil, types.ErrStakeTokenPermissionDeny
	}

	// check rewrad pool number
	maxRewardPoolNumber := k.Keeper.GetMaxRewardPoolNumber(ctx)
	if len(msg.RewardPoolInfoList) > int(maxRewardPoolNumber) {
		return nil, types.ErrRewardPoolNumberReachLimit
	}
	// check stake item number
	maxStakeItemNumber := k.Keeper.GetMaxStakeItemNumber(ctx)
	if len(msg.StakeItemInfoList) > int(maxStakeItemNumber) {
		return nil, types.ErrStakeItemNumberReachLimit
	}

	curBlockTime := uint64(ctx.BlockTime().Unix())
	willUseStakePoolIndex := k.Keeper.GetStakePoolNextIndex(ctx)

	// check and store stake item
	stakeItemLimit := k.Keeper.GetStakeItemLimit(ctx)
	maxLockSecond := uint64(0)
	for i, stakeItemInfo := range msg.StakeItemInfoList {
		if stakeItemInfo.LockSecond > maxLockSecond {
			maxLockSecond = stakeItemInfo.LockSecond
		}

		if stakeItemInfo.LockSecond > stakeItemLimit.MaxLockSecond {
			return nil, types.ErrStakeItemEraSecondExceedLimit
		}
		if stakeItemInfo.PowerRewardRate.GT(stakeItemLimit.MaxPowerRewardRate) {
			return nil, types.ErrStakeItemPowerRewardRateExceedLimit
		}

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

	// check and store reward pool
	rewardPools := make([]*types.RewardPool, 0)
	rewardTokens := sdk.NewCoins()
	for i, rewardPool := range msg.RewardPoolInfoList {
		rewardToken, found := k.Keeper.GetRewardToken(ctx, rewardPool.RewardTokenDenom)
		if !found {
			return nil, types.ErrRewardTokenNotSupport
		}
		if rewardPool.TotalRewardAmount.LT(rewardToken.MinTotalRewardAmount) {
			return nil, types.ErrTotalRewardAmountLessThanLimit
		}
		if rewardPool.RewardPerSecond.LT(rewardToken.MinRewardPerSecond) {
			return nil, types.ErrRewardPerSecondLessThanLimit
		}

		// check reward second and max lock second
		if rewardPool.TotalRewardAmount.Quo(rewardPool.RewardPerSecond).LT(sdk.NewIntFromUint64(maxLockSecond)) {
			return nil, types.ErrRewardSecondsLessThanMaxLockSeconds
		}

		willUseLastRewardTimestamp := rewardPool.StartTimestamp
		if rewardPool.StartTimestamp < curBlockTime {
			willUseLastRewardTimestamp = curBlockTime
		}

		rewardTokens = rewardTokens.Add(sdk.NewCoin(rewardPool.RewardTokenDenom, rewardPool.TotalRewardAmount))

		rewardPool := types.RewardPool{
			Index:               uint32(i),
			RewardTokenDenom:    rewardPool.RewardTokenDenom,
			TotalRewardAmount:   rewardPool.TotalRewardAmount,
			LeftRewardAmount:    rewardPool.TotalRewardAmount,
			RewardPerSecond:     rewardPool.RewardPerSecond,
			StartTimestamp:      willUseLastRewardTimestamp,
			RewardPerPower:      sdk.ZeroInt(),
			LastRewardTimestamp: willUseLastRewardTimestamp,
			Creator:             msg.Creator,
			TotalClaimedAmount:  sdk.ZeroInt(),
		}
		rewardPools = append(rewardPools, &rewardPool)
	}
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, user, types.ModuleName, rewardTokens); err != nil {
		return nil, err
	}

	stakePool := types.StakePool{
		Index:             willUseStakePoolIndex,
		StakeTokenDenom:   msg.StakeTokenDenom,
		RewardPools:       rewardPools,
		TotalStakedAmount: sdk.ZeroInt(),
		TotalStakedPower:  sdk.ZeroInt(),
		EmergencySwitch:   false,
		Creator:           msg.Creator,
	}

	k.SetStakePool(ctx, &stakePool)
	k.Keeper.SetStakePoolIndex(ctx, willUseStakePoolIndex)

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
