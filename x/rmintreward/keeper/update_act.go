package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
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
