package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/mining/types"
)

func updateStakePool(stakePool *types.StakePool, curBlockTime uint64) {
	for _, rewardPool := range stakePool.RewardPools {
		if rewardPool.LastRewardTimestamp >= curBlockTime {
			continue
		}
		if stakePool.TotalStakedPower.IsZero() {
			rewardPool.LastRewardTimestamp = curBlockTime
			continue
		}

		reward := getPoolReward(rewardPool.LastRewardTimestamp, curBlockTime, rewardPool.RewardPerSecond, rewardPool.LeftRewardAmount)
		if reward.IsPositive() {
			rewardPool.LeftRewardAmount = rewardPool.LeftRewardAmount.Sub(reward)
			if rewardPool.LeftRewardAmount.IsNegative() {
				rewardPool.LeftRewardAmount = sdk.ZeroInt()
			}
			willAddRewardPerPower := reward.Mul(types.RewardFactor).Quo(stakePool.TotalStakedPower)
			rewardPool.RewardPerPower = rewardPool.RewardPerPower.Add(willAddRewardPerPower)
		}
		rewardPool.LastRewardTimestamp = curBlockTime
	}
}

func calRewardTokens(stakePool *types.StakePool, userStakeRecord *types.UserStakeRecord) sdk.Coins {
	rewardCoins := sdk.NewCoins()
	for _, rewardPool := range stakePool.RewardPools {
		rewardDebt := sdk.ZeroInt()
		existInRewardInfos := false
		for _, rewardInfo := range userStakeRecord.RewardInfos {
			if rewardPool.Index == rewardInfo.RewardPoolIndex {
				rewardDebt = rewardInfo.RewardDebt
				rewardInfo.RewardDebt = userStakeRecord.StakedPower.Mul(rewardPool.RewardPerPower).Quo(types.RewardFactor)
				existInRewardInfos = true
				break
			}
		}
		if !existInRewardInfos {
			userStakeRecord.RewardInfos = append(userStakeRecord.RewardInfos, &types.RewardInfo{
				RewardPoolIndex:  rewardPool.Index,
				RewardTokenDenom: rewardPool.RewardTokenDenom,
				RewardDebt:       userStakeRecord.StakedPower.Mul(rewardPool.RewardPerPower).Quo(types.RewardFactor),
			})
		}

		rewardAmount := userStakeRecord.StakedPower.Mul(rewardPool.RewardPerPower).Quo(types.RewardFactor).Sub(rewardDebt)
		if rewardAmount.IsPositive() {
			rewardCoins = rewardCoins.Add(sdk.NewCoin(rewardPool.RewardTokenDenom, rewardAmount))
		}
	}
	return rewardCoins
}

func getPoolReward(from, to uint64, rewardPerSecond, leftRewardAmount sdk.Int) sdk.Int {
	duration := uint64(0)
	if from < to {
		duration = to - from
	}
	reward := rewardPerSecond.Mul(sdk.NewIntFromUint64(duration))
	if reward.GT(leftRewardAmount) {
		reward = leftRewardAmount
	}
	return reward
}
