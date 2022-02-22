package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgAddRelayer{}, "bridge/AddRelayer", nil)
	cdc.RegisterConcrete(&MsgSetThreshold{}, "bridge/SetThreshold", nil)
	cdc.RegisterConcrete(&MsgSetResourceidToDenom{}, "bridge/SetResourceidToDenom", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddRelayer{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetThreshold{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetResourceidToDenom{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
