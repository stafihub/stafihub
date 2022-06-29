package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/mining/types"
)

// will update {stakePool.RewardPools}
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

// will update {userStakeRecord.RewardInfos[n].ClaimedAmount} {stakePool.RewardPools[n].TotalClaimedAmount}, call after updateStakePool()
func calRewardTokens(stakePool *types.StakePool, userStakeRecord *types.UserStakeRecord) sdk.Coins {
	userRewardInfoMap := make(map[uint32]*types.UserRewardInfo)
	for _, rewardInfo := range userStakeRecord.UserRewardInfos {
		userRewardInfoMap[rewardInfo.RewardPoolIndex] = rewardInfo
	}

	rewardCoins := sdk.NewCoins()
	for _, rewardPool := range stakePool.RewardPools {
		var willUseRewardInfo *types.UserRewardInfo
		if rewardInfo, exist := userRewardInfoMap[rewardPool.Index]; exist {
			willUseRewardInfo = rewardInfo
		} else {
			willUseRewardInfo = &types.UserRewardInfo{
				RewardPoolIndex:  rewardPool.Index,
				RewardTokenDenom: rewardPool.RewardTokenDenom,
				RewardDebt:       sdk.ZeroInt(),
				ClaimedAmount:    sdk.ZeroInt(),
			}
			userStakeRecord.UserRewardInfos = append(userStakeRecord.UserRewardInfos, willUseRewardInfo)
		}

		rewardAmount := userStakeRecord.StakedPower.Mul(rewardPool.RewardPerPower).Quo(types.RewardFactor).Sub(willUseRewardInfo.RewardDebt)
		if rewardAmount.IsPositive() {
			rewardCoins = rewardCoins.Add(sdk.NewCoin(rewardPool.RewardTokenDenom, rewardAmount))
			willUseRewardInfo.ClaimedAmount = willUseRewardInfo.ClaimedAmount.Add(rewardAmount)
			rewardPool.TotalClaimedAmount = rewardPool.TotalClaimedAmount.Add(rewardAmount)
		}
	}
	return rewardCoins
}

// will update {userStakeRecord.RewardInfos[n].RewardDebt}, call after calRewardTokens()
func setNewRewardDebt(stakePool *types.StakePool, userStakeRecord *types.UserStakeRecord) {
	userRewardInfoMap := make(map[uint32]*types.UserRewardInfo)
	for _, rewardInfo := range userStakeRecord.UserRewardInfos {
		userRewardInfoMap[rewardInfo.RewardPoolIndex] = rewardInfo
	}

	for _, rewardPool := range stakePool.RewardPools {

		if rewardInfo, exist := userRewardInfoMap[rewardPool.Index]; exist {
			rewardInfo.RewardDebt = userStakeRecord.StakedPower.Mul(rewardPool.RewardPerPower).Quo(types.RewardFactor)
		} else {
			newUserRewardInfo := &types.UserRewardInfo{
				RewardPoolIndex:  rewardPool.Index,
				RewardTokenDenom: rewardPool.RewardTokenDenom,
				RewardDebt:       userStakeRecord.StakedPower.Mul(rewardPool.RewardPerPower).Quo(types.RewardFactor),
				ClaimedAmount:    sdk.ZeroInt(),
			}
			userStakeRecord.UserRewardInfos = append(userStakeRecord.UserRewardInfos, newUserRewardInfo)
		}
	}
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
