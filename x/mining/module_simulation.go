package mining

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/stafihub/stafihub/testutil/sample"
	miningsimulation "github.com/stafihub/stafihub/x/mining/simulation"
	"github.com/stafihub/stafihub/x/mining/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = miningsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgToggleProviderSwitch = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgToggleProviderSwitch int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	miningGenesis := types.GenesisState{
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&miningGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgToggleProviderSwitch int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgToggleProviderSwitch, &weightMsgToggleProviderSwitch, nil,
		func(_ *rand.Rand) {
			weightMsgToggleProviderSwitch = defaultWeightMsgToggleProviderSwitch
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgToggleProviderSwitch,
		miningsimulation.SimulateMsgToggleProviderSwitch(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
