package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
    cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	rvotetypes "github.com/stafiprotocol/stafihub/x/rvote/types"
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
	cdc.RegisterConcrete(&MsgSetCommission{}, "ledger/SetCommission", nil)
	cdc.RegisterConcrete(&MsgSetReceiver{}, "ledger/SetReceiver", nil)

	cdc.RegisterConcrete(&SetChainEraProposal{}, "ledger/SetChainEraProposal", nil)
	cdc.RegisterConcrete(&BondReportProposal{}, "ledger/BondReportProposal", nil)
	cdc.RegisterConcrete(&BondAndReportActiveProposal{}, "ledger/BondAndReportActiveProposal", nil)
	cdc.RegisterConcrete(&ActiveReportProposal{}, "ledger/ActiveReportProposal", nil)
	cdc.RegisterConcrete(&WithdrawReportProposal{}, "ledger/WithdrawReportProposal", nil)
	cdc.RegisterConcrete(&TransferReportProposal{}, "ledger/TransferReportProposal", nil)
// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddNewPool{},
		&MsgRemovePool{},
		&MsgSetEraUnbondLimit{},
		&MsgSetInitBond{},
		&MsgSetChainBondingDuration{},
		&MsgSetPoolDetail{},
		&MsgSetLeastBond{},
		&MsgClearCurrentEraSnapShots{},
		&MsgSetCommission{},
		&MsgSetReceiver{},
	)

	registry.RegisterImplementations(
		(*rvotetypes.Content)(nil),
		&SetChainEraProposal{},
		&BondReportProposal{},
		&BondAndReportActiveProposal{},
		&ActiveReportProposal{},
		&WithdrawReportProposal{},
		&TransferReportProposal{},
	)

// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
