package ledger

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/ledger/keeper"
	"github.com/stafihub/stafihub/x/ledger/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	for _, exchangeRate := range genState.ExchangeRateList {
		k.MigrateExchangeRate(ctx, exchangeRate.Denom, exchangeRate.Value)
	}

	for _, eraExchangeRate := range genState.EraExchangeRateList {
		k.SetEraExchangeRate(ctx, eraExchangeRate.Denom, eraExchangeRate.Era, eraExchangeRate.Value)
	}

	for _, totalProtocolFee := range genState.TotalProtocolFeeList {
		k.SetTotalProtocolFee(ctx, totalProtocolFee.Denom, totalProtocolFee.Value)
	}

	for _, unbondSwitch := range genState.UnbondSwitchList {
		k.SetUnbondSwitch(ctx, unbondSwitch.Denom, unbondSwitch.Switch)
	}

	for _, bondedPool := range genState.BondedPoolList {
		k.SetBondedPool(ctx, bondedPool)
	}

	for _, bondPipeline := range genState.BondPipelineList {
		k.SetBondPipeline(ctx, *bondPipeline)
	}

	for _, eraUnbondLimit := range genState.EraUnbondLimitList {
		k.SetEraUnbondLimit(ctx, eraUnbondLimit.Denom, eraUnbondLimit.Limit)
	}

	for _, poolDetail := range genState.PoolDetailList {
		k.SetPoolDetail(ctx, poolDetail)
	}

	for _, currentEraSnapshot := range genState.CurrentEraSnapshotList {
		k.SetCurrentEraSnapshot(ctx, *currentEraSnapshot)
	}

	for _, snapshot := range genState.SnapshotList {
		k.SetSnapshot(ctx, snapshot.ShotId, *snapshot.Snapshot)
	}

	for _, eraSnapshot := range genState.EraSnapshotList {
		k.SetEraSnapshot(ctx, eraSnapshot.Era, types.EraSnapshot{
			Denom:   eraSnapshot.Denom,
			ShotIds: eraSnapshot.ShotIds,
		})
	}

	for _, chainEra := range genState.ChainEraList {
		k.SetChainEra(ctx, chainEra.Denom, chainEra.Era)
	}

	for _, stakingRewardCommission := range genState.StakingRewardCommissionList {
		k.SetStakingRewardCommission(ctx, stakingRewardCommission.Denom, stakingRewardCommission.Value)
	}

	if len(genState.ProtocolFeeReceiver) != 0 {
		protocolFeeReceiver, err := sdk.AccAddressFromBech32(genState.ProtocolFeeReceiver)
		if err != nil {
			panic(err)
		}
		k.SetProtocolFeeReceiver(ctx, protocolFeeReceiver)
	}

	for _, relayFeeReceiver := range genState.RelayFeeReceiverList {
		addr, err := sdk.AccAddressFromBech32(relayFeeReceiver.Address)
		if err != nil {
			panic(err)
		}
		k.SetRelayFeeReceiver(ctx, relayFeeReceiver.Denom, addr)
	}

	for _, totalExpectedActive := range genState.TotalExpectedActiveList {
		k.SetTotalExpectedActive(ctx, totalExpectedActive.Denom, totalExpectedActive.Era, totalExpectedActive.Value)
	}

	for _, poolUnbonding := range genState.PoolUnbondingList {
		k.SetPoolUnbonding(ctx, poolUnbonding.Denom, poolUnbonding.Pool, poolUnbonding.Era, poolUnbonding.Sequence, poolUnbonding.Unbonding)
		k.SetPoolUnbondSequence(ctx, poolUnbonding.Denom, poolUnbonding.Pool, poolUnbonding.Era, poolUnbonding.Sequence)
	}

	for _, unbondRelayFee := range genState.UnbondRelayFeeList {
		k.SetUnbondRelayFee(ctx, unbondRelayFee.Denom, unbondRelayFee.Value)
	}

	for _, unbondCommission := range genState.UnbondCommissionList {
		k.SetUnbondCommission(ctx, unbondCommission.Denom, unbondCommission.Value)
	}

	for _, bondRecord := range genState.BondRecordList {
		k.SetBondRecord(ctx, *bondRecord)
	}

	for _, rparams := range genState.RparamsList {
		k.SetRParams(ctx, *rparams)
	}

	for _, signature := range genState.SignatureList {
		k.SetSignature(ctx, *signature)
	}

	if genState.MigrateIsSealed {
		k.SealMigrateInit(ctx)
	} else {
		k.UnSealMigrateInit(ctx)
	}

	for i, icaPool := range k.GetAllIcaPoolDetailList(ctx) {
		k.SetIcaPoolDetail(ctx, icaPool)
		k.SetIcaPoolIndex(ctx, icaPool.Denom, uint32(i))
		k.SetIcaPoolDelegationAddrIndex(ctx, icaPool)
	}

	for _, interchainTxProposal := range k.GetInterchainTxProposalInfoList(ctx) {
		k.SetInterchainTxProposalStatus(ctx, interchainTxProposal.ProposalId, interchainTxProposal.Status)
		k.SetInterchainTxProposalSequenceIndex(ctx, interchainTxProposal.CtrlPortId, interchainTxProposal.CtrlChannelId,
			interchainTxProposal.Sequence, interchainTxProposal.ProposalId)
	}

	for _, totalExpectedFee := range genState.TotalExpectedFeeList {
		k.SetTotalExpectedFee(ctx, totalExpectedFee.Denom, totalExpectedFee.Era, totalExpectedFee.Value)
	}
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
	genesis.MigrateIsSealed = k.MigrateInitIsSealed(ctx)
	genesis.IcaPoolDetailList = k.GetAllIcaPoolDetailList(ctx)
	genesis.InterchainTxProposalInfoList = k.GetInterchainTxProposalInfoList(ctx)
	genesis.TotalExpectedFeeList = k.TotalExpectedFeeList(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
