package rvalidator

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rvalidator/keeper"
	"github.com/stafihub/stafihub/x/rvalidator/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
	for _, cycleSeconds := range genState.CycleSecondsList {
		k.SetCycleSeconds(ctx, cycleSeconds)
	}
	for _, latestDealedCycle := range genState.LatestDealedCycleList {
		k.SetLatestDealedCycle(ctx, latestDealedCycle)
	}

	for _, latestVotedCycle := range genState.LatestVotedCycleList {
		k.SetLatestVotedCycle(ctx, latestVotedCycle)
	}
	for _, selectedRValidator := range genState.SelectedRValidatorList {
		if err := k.RBankKeeper.CheckValAddress(ctx, selectedRValidator.Denom, selectedRValidator.ValAddress); err != nil {
			panic(err)
		}
		if err := k.RBankKeeper.CheckAccAddress(ctx, selectedRValidator.Denom, selectedRValidator.PoolAddress); err != nil {
			panic(err)
		}
		k.AddSelectedRValidator(ctx, selectedRValidator)
	}

	for _, shuffleSeconds := range genState.ShuffleSecondsList {
		k.SetShuffleSeconds(ctx, shuffleSeconds)
	}

	for _, dealingRValidator := range genState.DealingRValidatorList {
		if err := k.RBankKeeper.CheckValAddress(ctx, dealingRValidator.Denom, dealingRValidator.NewValAddress); err != nil {
			panic(err)
		}
		if err := k.RBankKeeper.CheckValAddress(ctx, dealingRValidator.Denom, dealingRValidator.OldValAddress); err != nil {
			panic(err)
		}
		if err := k.RBankKeeper.CheckAccAddress(ctx, dealingRValidator.Denom, dealingRValidator.PoolAddress); err != nil {
			panic(err)
		}
		k.SetDealingRValidator(ctx, dealingRValidator)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.CycleSecondsList = k.GetAllCycleSeconds(ctx)
	genesis.LatestDealedCycleList = k.GetAllLatestDealedCycle(ctx)
	genesis.LatestVotedCycleList = k.GetAllLatestVotedCycle(ctx)
	genesis.SelectedRValidatorList = k.GetSelectedRValidatorList(ctx)
	genesis.ShuffleSecondsList = k.GetAllShuffleSeconds(ctx)
	genesis.DealingRValidatorList = k.GetAllDealingRvalidators(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
