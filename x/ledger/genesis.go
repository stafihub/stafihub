package ledger

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/ledger/keeper"
	"github.com/stafihub/stafihub/x/ledger/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.ExchangeRateList = k.GetAllExchangeRate(ctx)
	genesis.EraExchangeRateList = k.GetEraExchangeRateList(ctx)
	genesis.TotalProtocolFeeList = k.GetAllTotalProtocolFee(ctx)
	genesis.UnbondSwitchList = k.GetUnbondSwitchList(ctx)

	genesis.BondedPoolList = k.GetBondedPoolList(ctx)
	genesis.BondPipelineList = k.GetBondPipelineList(ctx)
	genesis.EraUnbondLimitList = k.GetEraUnbondLimitList(ctx)
	genesis.PoolDetailList = k.GetPoolDetailList(ctx)
	genesis.CurrentEraSnapshotList = k.CurrentEraSnapshotList(ctx)
	genesis.SnapshotList = k.SnapshotList(ctx)
	genesis.EraSnapshotList = k.EraSnapshotList(ctx)
	genesis.ChainEraList = k.GetChainEraList(ctx)
	genesis.StakingRewardCommissionList = k.GetStakingRewardCommissionList(ctx)
	feeReceiver, found := k.GetProtocolFeeReceiver(ctx)
	if found {
		genesis.ProtocolFeeReceiver = feeReceiver.String()
	}
	genesis.RelayFeeReceiverList = k.GetRelayFeeReceiverList(ctx)
	genesis.TotalExpectedActiveList = k.TotalExpectedActiveList(ctx)
	genesis.PoolUnbondingList = k.GetPoolUnbondingList(ctx)
	genesis.UnbondRelayFeeList = k.GetUnbondRelayFeeList(ctx)
	genesis.UnbondCommissionList = k.GetUnbondCommissionList(ctx)
	genesis.BondRecordList = k.GetBondRecordList(ctx)
	genesis.RparamsList = k.GetRParamsList(ctx)
	genesis.SignatureList = k.GetSignatureList(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
