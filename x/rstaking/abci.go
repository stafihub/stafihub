package rstaking

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rstaking/keeper"
	"github.com/stafihub/stafihub/x/rstaking/types"
)

func BeginBlocker(ctx sdk.Context, k keeper.Keeper, mintKeeper types.MintKeeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	// fetch stored minter & params
	minter := mintKeeper.GetMinter(ctx)
	params := mintKeeper.GetParams(ctx)

	// recalculate inflation rate
	totalStakingSupply := k.StakingTokenSupply(ctx)
	bondedRatio := k.BondedRatio(ctx)
	minter.Inflation = minter.NextInflationRate(params, bondedRatio)
	minter.AnnualProvisions = minter.NextAnnualProvisions(params, totalStakingSupply)

	mintedCoin := minter.BlockProvision(params)
	mintedCoins := sdk.NewCoins(mintedCoin)

	err := k.GetBankKeeper().BurnCoins(ctx, types.ModuleName, mintedCoins)
	// inflation is zero in err case
	if err != nil {
		_ = k.GetBankKeeper().BurnCoins(ctx, k.GetFeeCollectorName(), mintedCoins)
	}
}
