package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgAddRValidator{}, "rvalidator/AddRValidator", nil)
	cdc.RegisterConcrete(&MsgRmRValidator{}, "rvalidator/RmRValidator", nil)
	cdc.RegisterConcrete(&UpdateRValidatorProposal{}, "rvalidator/UpdateRValidator", nil)
	cdc.RegisterConcrete(&MsgSetCycleSeconds{}, "rvalidator/SetCycleSeconds", nil)
	cdc.RegisterConcrete(&MsgSetShuffleSeconds{}, "rvalidator/SetShuffleSeconds", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddRValidator{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRmRValidator{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&UpdateRValidatorProposal{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetCycleSeconds{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetShuffleSeconds{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
