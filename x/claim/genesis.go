package claim

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/claim/keeper"
	"github.com/stafihub/stafihub/x/claim/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)

	for _, claimBitMap := range genState.ClaimBitMapList {
		k.SetClaimBitMap(ctx, claimBitMap.Round, claimBitMap.WordIndex, claimBitMap.Bits)
	}

	for _, claimSwitch := range genState.ClaimSwitchList {
		k.SetClaimSwitch(ctx, claimSwitch.Round, claimSwitch.IsOpen)
	}

	for _, merkleRoot := range genState.MerkleRootList {
		nodeHash, err := keeper.NodeHashFromHexString(merkleRoot.RootHash)
		if err != nil {
			panic(err)
		}
		k.SetMerkleRoot(ctx, merkleRoot.Round, nodeHash)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.ClaimBitMapList = k.GetClaimBitMapList(ctx)
	genesis.ClaimSwitchList = k.GetClaimSwitchList(ctx)
	genesis.MerkleRootList = k.GetMerkleRootList(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
