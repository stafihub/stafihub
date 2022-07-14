package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	rvotetypes "github.com/stafihub/stafihub/x/rvote/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSetEraUnbondLimit{}, "ledger/SetEraUnbondLimit", nil)
	cdc.RegisterConcrete(&MsgSetPoolDetail{}, "ledger/SetPoolDetail", nil)
	cdc.RegisterConcrete(&MsgSetLeastBond{}, "ledger/SetLeastBond", nil)
	cdc.RegisterConcrete(&MsgClearCurrentEraSnapShots{}, "ledger/ClearCurrentEraSnapShots", nil)
	cdc.RegisterConcrete(&MsgSetStakingRewardCommission{}, "ledger/SetStakingRewardCommission", nil)
	cdc.RegisterConcrete(&MsgSetProtocolFeeReceiver{}, "ledger/SetProtocolFeeReceiver", nil)
	cdc.RegisterConcrete(&MsgSetUnbondRelayFee{}, "ledger/SetUnbondRelayFee", nil)
	cdc.RegisterConcrete(&MsgLiquidityUnbond{}, "ledger/LiquidityUnbond", nil)
	cdc.RegisterConcrete(&MsgSetUnbondCommission{}, "ledger/SetUnbondCommission", nil)

	cdc.RegisterConcrete(&SetChainEraProposal{}, "ledger/SetChainEraProposal", nil)
	cdc.RegisterConcrete(&BondReportProposal{}, "ledger/BondReportProposal", nil)
	cdc.RegisterConcrete(&ActiveReportProposal{}, "ledger/ActiveReportProposal", nil)
	cdc.RegisterConcrete(&TransferReportProposal{}, "ledger/TransferReportProposal", nil)
	cdc.RegisterConcrete(&ExecuteBondProposal{}, "ledger/ExecuteBondProposal", nil)
	cdc.RegisterConcrete(&InterchainTxProposal{}, "ledger/InterchainTxProposal", nil)

	cdc.RegisterConcrete(&MsgSubmitSignature{}, "ledger/SubmitSignature", nil)
	cdc.RegisterConcrete(&MsgSetRParams{}, "ledger/SetRParams", nil)
	cdc.RegisterConcrete(&MsgSetRelayFeeReceiver{}, "ledger/SetRelayFeeReceiver", nil)
	cdc.RegisterConcrete(&MsgSetRelayGasPrice{}, "ledger/SetRelayGasPrice", nil)
	cdc.RegisterConcrete(&MsgSetEraSeconds{}, "ledger/SetEraSeconds", nil)
	cdc.RegisterConcrete(&MsgRmBondedPool{}, "ledger/RmBondedPool", nil)
	cdc.RegisterConcrete(&MsgMigrateInit{}, "ledger/MigrateInit", nil)
	cdc.RegisterConcrete(&MsgMigrateUnbondings{}, "ledger/MigrateUnbondings", nil)
	cdc.RegisterConcrete(&MsgToggleUnbondSwitch{}, "ledger/ToggleUnbondSwitch", nil)
	cdc.RegisterConcrete(&MsgUnsealMigrateInit{}, "ledger/UnsealMigrateInit", nil)
	cdc.RegisterConcrete(&MsgRegisterIcaPool{}, "ledger/RegisterIcaPool", nil)
	cdc.RegisterConcrete(&MsgSetWithdrawAddr{}, "ledger/SetWithdrawAddr", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetEraUnbondLimit{},
		&MsgSetPoolDetail{},
		&MsgSetLeastBond{},
		&MsgClearCurrentEraSnapShots{},
		&MsgSetStakingRewardCommission{},
		&MsgSetProtocolFeeReceiver{},
	)

	registry.RegisterImplementations(
		(*rvotetypes.Content)(nil),
		&SetChainEraProposal{},
		&BondReportProposal{},
		&ActiveReportProposal{},
		&TransferReportProposal{},
		&ExecuteBondProposal{},
		&InterchainTxProposal{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSubmitSignature{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetRParams{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetRelayFeeReceiver{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetRelayGasPrice{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetEraSeconds{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRmBondedPool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMigrateInit{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMigrateUnbondings{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgToggleUnbondSwitch{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUnsealMigrateInit{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterIcaPool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetWithdrawAddr{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
