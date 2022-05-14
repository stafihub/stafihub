package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/mining/types"
)

func (k msgServer) ClaimReward(goCtx context.Context, msg *types.MsgClaimReward) (*types.MsgClaimRewardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	recipientAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	userStakeRecord, found := k.Keeper.GetUserStakeRecord(ctx, msg.Creator, msg.StakeTokenDenom, msg.Index)
	if !found {
		return nil, types.ErrUserStakeRecordNotExist
	}
	stakePool, found := k.Keeper.GetStakePool(ctx, msg.StakeTokenDenom)
	if !found {
		return nil, types.ErrStakePoolNotExist
	}
	curBlockTime := uint64(ctx.BlockTime().Unix())

	updateStakePool(stakePool, curBlockTime)

	willClaimCoins := sdk.NewCoins()
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
			willClaimCoins = append(willClaimCoins, sdk.NewCoin(rewardPool.RewardTokenDenom, rewardAmount))
		}
	}

	if willClaimCoins.Len() > 0 {
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recipientAddr, willClaimCoins); err != nil {
			return nil, err
		}
	}

	k.SetStakePool(ctx, stakePool)
	k.SetUserStakeRecord(ctx, userStakeRecord)

	return &types.MsgClaimRewardResponse{}, nil
}
