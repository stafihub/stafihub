package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSetMerkleRoot{}, "claim/SetMerkleRoot", nil)
	cdc.RegisterConcrete(&MsgClaim{}, "claim/Claim", nil)
	cdc.RegisterConcrete(&MsgToggleClaimSwitch{}, "claim/ToggleClaimSwitch", nil)
	cdc.RegisterConcrete(&MsgProvideToken{}, "claim/ProvideToken", nil)
	cdc.RegisterConcrete(&MsgWithdrawToken{}, "claim/WithdrawToken", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetMerkleRoot{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgClaim{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgToggleClaimSwitch{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgProvideToken{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgWithdrawToken{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
