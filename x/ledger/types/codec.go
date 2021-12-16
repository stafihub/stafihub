package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
    cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
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
cdc.RegisterConcrete(&MsgSetChainEra{}, "ledger/SetChainEra", nil)
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
	&MsgSetChainEra{},
)
// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
