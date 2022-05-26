package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/mining/types"
)

func (k msgServer) AddReward(goCtx context.Context, msg *types.MsgAddReward) (*types.MsgAddRewardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	user, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	curBlockTime := uint64(ctx.BlockTime().Unix())

	stakePool, found := k.Keeper.GetStakePool(ctx, msg.StakePoolIndex)
	if !found {
		return nil, types.ErrStakePoolNotExist
	}
	if stakePool.EmergencySwitch {
		return nil, types.ErrEmergencySwitchOpen
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

	updateStakePool(stakePool, curBlockTime)

	// can reset rewardPerSecond/startTimestamp if reward pool is end
	if willUseRewardPool.LeftRewardAmount.IsZero() {
		// check min limit
		rewardToken, found := k.Keeper.GetRewardToken(ctx, willUseRewardPool.RewardTokenDenom)
		if !found {
			return nil, types.ErrRewardTokenNotSupport
		}
		if msg.AddAmount.LT(rewardToken.MinTotalRewardAmount) {
			return nil, types.ErrTotalRewardAmountLessThanLimit
		}

		willUseRewardPool.StartTimestamp = msg.StartTimestamp
		willUseRewardPool.LastRewardTimestamp = msg.StartTimestamp

		if msg.StartTimestamp < curBlockTime {
			willUseRewardPool.StartTimestamp = curBlockTime
			willUseRewardPool.LastRewardTimestamp = curBlockTime
		}

		if msg.RewardPerSecond.IsPositive() {
			willUseRewardPool.RewardPerSecond = msg.RewardPerSecond
		}
	} else {
		if msg.StartTimestamp != 0 || !msg.RewardPerSecond.IsZero() {
			return nil, types.ErrStartTimestampAndRewardPerSecondNotZero
		}
	}
	// update total rewardAmount/leftRewardAmount
	willUseRewardPool.TotalRewardAmount = willUseRewardPool.TotalRewardAmount.Add(msg.AddAmount)
	willUseRewardPool.LeftRewardAmount = willUseRewardPool.LeftRewardAmount.Add(msg.AddAmount)

	rewardTokens := sdk.NewCoins(sdk.NewCoin(willUseRewardPool.RewardTokenDenom, msg.AddAmount))
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, user, types.ModuleName, rewardTokens); err != nil {
		return nil, err
	}

	k.Keeper.SetStakePool(ctx, stakePool)

	return &types.MsgAddRewardResponse{}, nil
}
