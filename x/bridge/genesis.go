package bridge

import (
	"encoding/hex"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/bridge/keeper"
	"github.com/stafihub/stafihub/x/bridge/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)

	for _, depositCount := range genState.DepositCountList {
		k.SetDepositCount(ctx, uint8(depositCount.ChainId), depositCount.Count)
	}

	for _, proposal := range genState.ProposalList {
		resourceBts, err := hex.DecodeString(proposal.ResourceId)
		if err != nil {
			panic(err)
		}
		var resourceId [32]byte
		copy(resourceId[:], resourceBts)

		k.SetProposal(ctx, uint8(proposal.ChainId), proposal.DepositNonce, resourceId, proposal.Proposal)
	}

	for _, relayFee := range genState.RelayFeeList {
		k.SetRelayFee(ctx, uint8(relayFee.ChainId), relayFee.Value)
	}

	for _, chainId := range genState.ChainIdList {
		k.AddChainId(ctx, uint8(chainId))
	}

	if len(genState.RelayFeeReceiver) != 0 {
		addr, err := sdk.AccAddressFromBech32(genState.RelayFeeReceiver)
		if err != nil {
			panic(err)
		}
		k.SetRelayFeeReceiver(ctx, addr)
	}

	for _, resourceIdToDenom := range genState.ResourceIdToDenom {
		strs := strings.Split(resourceIdToDenom, ":")
		if len(strs) != 2 {
			panic("not valid resourceIdToDenom")
		}

		resourceBts, err := hex.DecodeString(strs[0])
		if err != nil {
			panic(err)
		}
		var resourceId [32]byte
		copy(resourceId[:], resourceBts)
		k.SetResourceIdToDenom(ctx, resourceId, strs[1])
	}

	for _, resourceIdToDenomType := range genState.ResourceIdToDenomType {
		strs := strings.Split(resourceIdToDenomType, ":")
		if len(strs) != 2 {
			panic("not valid resourceIdToDenomType")
		}
		resourceBts, err := hex.DecodeString(strs[0])
		if err != nil {
			panic(err)
		}
		var resourceId [32]byte
		copy(resourceId[:], resourceBts)
		denomType := types.ResourceIdTypeForeign
		if strs[1] == "1" {
			denomType = types.ResourceIdTypeNative
		}
		k.SetResourceIdType(ctx, resourceId, denomType)
	}

}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.DepositCountList = k.GetDepositCountList(ctx)
	genesis.ProposalList = k.GetProposalList(ctx)
	genesis.RelayFeeList = k.GetRelayFeeList(ctx)
	genesis.ChainIdList = k.GetChainIdList(ctx)
	relayFeeReceiver, found := k.GetRelayFeeReceiver(ctx)
	if found {
		genesis.RelayFeeReceiver = relayFeeReceiver.String()
	}

	genesis.ResourceIdToDenom = k.GetAllResourceIdToDenom(ctx)
	genesis.ResourceIdToDenomType = k.GetAllResourceIdDenomTypes(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
