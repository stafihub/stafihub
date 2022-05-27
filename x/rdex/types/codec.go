package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreatePool{}, "rdex/CreatePool", nil)
	cdc.RegisterConcrete(&MsgAddLiquidity{}, "rdex/AddLiquidity", nil)
	cdc.RegisterConcrete(&MsgSwap{}, "rdex/Swap", nil)
	cdc.RegisterConcrete(&MsgRemoveLiquidity{}, "rdex/RemoveLiquidity", nil)
	cdc.RegisterConcrete(&MsgToggleProviderSwitch{}, "rdex/ToggleProviderSwitch", nil)
	cdc.RegisterConcrete(&MsgAddProvider{}, "rdex/AddProvider", nil)
	cdc.RegisterConcrete(&MsgRmProvider{}, "rdex/RmProvider", nil)
	cdc.RegisterConcrete(&MsgAddPoolCreator{}, "rdex/AddPoolCreator", nil)
	cdc.RegisterConcrete(&MsgRmPoolCreator{}, "rdex/RmPoolCreator", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddLiquidity{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSwap{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRemoveLiquidity{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgToggleProviderSwitch{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddProvider{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRmProvider{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddPoolCreator{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRmPoolCreator{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
