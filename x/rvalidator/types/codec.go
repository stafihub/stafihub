package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	rvotetypes "github.com/stafihub/stafihub/x/rvote/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgInitRValidator{}, "rvalidator/InitRValidator", nil)
	cdc.RegisterConcrete(&UpdateRValidatorProposal{}, "rvalidator/UpdateRValidator", nil)
	cdc.RegisterConcrete(&UpdateRValidatorReportProposal{}, "rvalidator/UpdateRValidatorReport", nil)
	cdc.RegisterConcrete(&MsgSetCycleSeconds{}, "rvalidator/SetCycleSeconds", nil)
	cdc.RegisterConcrete(&MsgSetShuffleSeconds{}, "rvalidator/SetShuffleSeconds", nil)
	cdc.RegisterConcrete(&MsgAddRValidator{}, "rvalidator/AddRValidator", nil)
	cdc.RegisterConcrete(&MsgRmRValidator{}, "rvalidator/RmRValidator", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgInitRValidator{},
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
	registry.RegisterImplementations(
		(*rvotetypes.Content)(nil),
		&UpdateRValidatorProposal{},
	)
	registry.RegisterImplementations(
		(*rvotetypes.Content)(nil),
		&UpdateRValidatorReportProposal{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddRValidator{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRmRValidator{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(Amino)
)

func init() {
	RegisterCodec(Amino)
	cryptocodec.RegisterCrypto(Amino)
	sdk.RegisterLegacyAminoCodec(Amino)
	Amino.Seal()
}
