package rmintreward

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rmintreward/keeper"
	"github.com/stafihub/stafihub/x/rmintreward/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)

	for _, actLatestCycle := range genState.ActLatestCycleList {
		k.SetActLatestCycle(ctx, actLatestCycle.Denom, actLatestCycle.Cycle)
	}

	for _, actCurrentCycle := range genState.ActCurrentCycleList {
		k.SetActCurrentCycle(ctx, actCurrentCycle.Denom, actCurrentCycle.Cycle)
	}

	for _, mintRewardAct := range genState.MintrewardActList {
		k.SetMintRewardAct(ctx, mintRewardAct.Denom, mintRewardAct.Cycle, mintRewardAct.MintRewardAct)
	}

	for _, userClaimInfo := range genState.UserClaimInfoList {
		account, err := sdk.AccAddressFromBech32(userClaimInfo.Account)
		if err != nil {
			panic(err)
		}
		k.SetUserClaimInfo(ctx, account, userClaimInfo.Denom, userClaimInfo.Cycle, userClaimInfo.MintIndex, userClaimInfo.UserClaimInfo)
	}

	for _, userActs := range genState.UserActList {
		account, err := sdk.AccAddressFromBech32(userActs.Account)
		if err != nil {
			panic(err)
		}
		k.SetUserActs(ctx, account, userActs.Denom, userActs.Acts)
	}

	for _, userMintCount := range genState.UserMintCountList {
		account, err := sdk.AccAddressFromBech32(userMintCount.Account)
		if err != nil {
			panic(err)
		}
		k.SetUserMintCount(ctx, account, userMintCount.Denom, userMintCount.Cycle, userMintCount.Count)
	}

}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.ActLatestCycleList = k.GetActLatestCycleList(ctx)
	genesis.ActCurrentCycleList = k.GetActCurrentCycleList(ctx)
	genesis.MintrewardActList = k.GetMintRewardActList(ctx)
	genesis.UserClaimInfoList = k.GetUserClaimInfoList(ctx)
	genesis.UserActList = k.GetUserActsList(ctx)
	genesis.UserMintCountList = k.GetUserMintCountList(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
