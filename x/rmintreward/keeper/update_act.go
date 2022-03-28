package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rmintreward/types"
)

func (k Keeper) UpdateActLatestCycle(ctx sdk.Context) {
	now := ctx.BlockHeight()
	denoms := k.GetActDenoms(ctx)
	for _, denom := range denoms {
		latestCycle, found := k.GetActLatestCycle(ctx, denom)
		if !found {
			continue
		}

		lastCurrentCycle, found := k.GetActCurrenttCycle(ctx, denom)
		if found && lastCurrentCycle == latestCycle {
			continue
		}

		begin := lastCurrentCycle + 1
		if !found {
			begin = 0
		}
		for i := begin; i <= latestCycle; i++ {
			act, found := k.GetMintRewardAct(ctx, denom, i)
			if !found {
				continue
			}
			if now < act.Begin {
				break
			}
			if act.Begin <= now && act.End >= now {
				if i != lastCurrentCycle {
					k.SetActCurrentCycle(ctx, denom, i)
				}
				break
			}
		}

	}
}

func (k Keeper) UpdateUserClaimInfo(ctx sdk.Context, user sdk.AccAddress, denom string, mintRTokenAmount, nativeTokenAmount sdk.Int) {
	currentCycle, found := k.GetActCurrenttCycle(ctx, denom)
	if !found {
		return
	}
	act, found := k.GetMintRewardAct(ctx, denom, currentCycle)
	if !found {
		return
	}

	now := ctx.BlockHeight()
	if act.Begin > now || act.End < now {
		return
	}

	userClaimInfo := types.UserClaimInfo{
		MintRTokenAmount:   mintRTokenAmount,
		NativeTokenAmount:  nativeTokenAmount,
		MintBLock:          now,
		LatestClaimedBlock: now,
		TokenClaimInfos:    []*types.TokenClaimInfo{},
	}

	for _, rewardInfo := range act.TokenRewardInfos {
		if rewardInfo.LeftAmount.LTE(sdk.ZeroInt()) {
			continue
		}
		shouldRewardAmount := rewardInfo.RewardRate.MulInt(mintRTokenAmount).RoundInt()

		if shouldRewardAmount.GT(rewardInfo.LeftAmount) {
			shouldRewardAmount = rewardInfo.LeftAmount
		}
		if rewardInfo.UserLimit.GT(sdk.ZeroInt()) && shouldRewardAmount.GT(rewardInfo.UserLimit) {
			shouldRewardAmount = rewardInfo.UserLimit
		}

		rewardInfo.LeftAmount = rewardInfo.LeftAmount.Sub(shouldRewardAmount)
		userClaimInfo.TokenClaimInfos = append(userClaimInfo.TokenClaimInfos, &types.TokenClaimInfo{
			Denom:              denom,
			TotalRewardAmount:  shouldRewardAmount,
			TotalClaimedAmount: sdk.ZeroInt(),
		})
	}

	act.TotalRTokenAmount = act.TotalRTokenAmount.Add(mintRTokenAmount)
	act.TotalNativeTokenAmount = act.TotalNativeTokenAmount.Add(nativeTokenAmount)

	count, found := k.GetUserMintCount(ctx, user, denom, currentCycle)
	if !found {
		count = 0
	}

	userAct, found := k.GetUserActs(ctx, user, denom)
	if !found {
		userAct = &types.Acts{
			Acts: []uint64{},
		}
	}
	userAct.Acts = append(userAct.Acts, currentCycle)

	k.SetUserActs(ctx, user, denom, userAct)
	k.SetUserClaimInfo(ctx, user, denom, currentCycle, count, &userClaimInfo)
	k.SetUserMintCount(ctx, user, denom, currentCycle, count+1)
	k.SetMintRewardAct(ctx, denom, currentCycle, act)

}
