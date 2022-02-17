package stakextra

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stafihub/stafihub/x/stakextra/keeper"
	"github.com/stafihub/stafihub/x/stakextra/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, mintKeeper types.MintKeeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
	k.SetInflationBase(ctx, genState.InflationBase)
	moduleAddress := authTypes.NewModuleAddress(types.ModuleName)
	params := mintKeeper.GetParams(ctx)
	if params.MintDenom != genState.CoinToBeBurned.Denom {
		panic("mint denom not equal coinToBeBurned denom")
	}

	balance := k.GetBankKeeper().GetBalance(ctx, moduleAddress, params.MintDenom)

	if balance.Amount.GT(sdk.ZeroInt()) {
		k.GetBankKeeper().BurnCoins(ctx, types.ModuleName, sdk.NewCoins(balance))
	}
	k.GetBankKeeper().MintCoins(ctx, types.ModuleName, sdk.NewCoins(genState.CoinToBeBurned))
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper, mintKeeper types.MintKeeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.InflationBase = k.GetInflationBase(ctx)
	params := mintKeeper.GetParams(ctx)
	moduleAddress := authTypes.NewModuleAddress(types.ModuleName)
	balance := k.GetBankKeeper().GetBalance(ctx, moduleAddress, params.MintDenom)
	genesis.CoinToBeBurned = balance
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
