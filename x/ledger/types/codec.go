package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
    cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	relayertypes "github.com/stafiprotocol/stafihub/x/relayers/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgAddNewPool{}, "ledger/AddNewPool", nil)
	cdc.RegisterConcrete(&MsgRemovePool{}, "ledger/RemovePool", nil)
	cdc.RegisterConcrete(&MsgSetEraUnbondLimit{}, "ledger/SetEraUnbondLimit", nil)
	cdc.RegisterConcrete(&MsgSetInitBond{}, "ledger/SetInitBond", nil)
	cdc.RegisterConcrete(&MsgSetChainBondingDuration{}, "ledger/SetChainBondingDuration", nil)
	cdc.RegisterConcrete(&MsgSetPoolDetail{}, "ledger/SetPoolDetail", nil)
	cdc.RegisterConcrete(&MsgSetLeastBond{}, "ledger/SetLeastBond", nil)
	cdc.RegisterConcrete(&MsgClearCurrentEraSnapShots{}, "ledger/ClearCurrentEraSnapShots", nil)
	cdc.RegisterConcrete(&SetChainEraProposal{}, "ledger/SetChainEraProposal", nil)
	cdc.RegisterConcrete(&MsgSetCommission{}, "ledger/SetCommission", nil)
	cdc.RegisterConcrete(&MsgSetReceiver{}, "ledger/SetReceiver", nil)
	cdc.RegisterConcrete(&ActiveReportProposal{}, "ledger/ActiveReportProposal", nil)

	cdc.RegisterConcrete(&BondReportProposal{}, "ledger/BondReportProposal", nil)
	cdc.RegisterConcrete(&BondAndReportActiveProposal{}, "ledger/BondAndReportActiveProposal", nil)
	cdc.RegisterConcrete(&WithdrawReportProposal{}, "ledger/WithdrawReportProposal", nil)
	cdc.RegisterConcrete(&TransferReportProposal{}, "ledger/TransferReportProposal", nil)
// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgAddNewPool{},
)
registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgRemovePool{},
)
registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgSetEraUnbondLimit{},
)
registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgSetInitBond{},
)
registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgSetChainBondingDuration{},
)
registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgSetPoolDetail{},
)
registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgSetLeastBond{},
)
registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgClearCurrentEraSnapShots{},
)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetCommission{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetReceiver{},
	)
registry.RegisterImplementations((*relayertypes.Content)(nil),
	&SetChainEraProposal{},
)
registry.RegisterImplementations((*relayertypes.Content)(nil),
	&ActiveReportProposal{},
)

registry.RegisterImplementations((*relayertypes.Content)(nil),
	&BondReportProposal{},
)
registry.RegisterImplementations((*relayertypes.Content)(nil),
	&BondAndReportActiveProposal{},
)
registry.RegisterImplementations((*relayertypes.Content)(nil),
	&WithdrawReportProposal{},
)
registry.RegisterImplementations((*relayertypes.Content)(nil),
	&TransferReportProposal{},
)
// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
