package rate

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafiprotocol/stafihub/x/rate/keeper"
	"github.com/stafiprotocol/stafihub/x/rate/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
    // Set all the exchangeRate
for _, elem := range genState.ExchangeRateList {
	k.SetExchangeRate(ctx, elem.Denom, elem.Value)
}
// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

    genesis.ExchangeRateList = k.GetAllExchangeRate(ctx)
// this line is used by starport scaffolding # genesis/module/export

    return genesis
}
