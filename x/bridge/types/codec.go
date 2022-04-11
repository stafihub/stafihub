package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSetResourceidToDenom{}, "bridge/SetResourceidToDenom", nil)
	cdc.RegisterConcrete(&MsgDeposit{}, "bridge/Deposit", nil)
	cdc.RegisterConcrete(&MsgAddChainId{}, "bridge/AddChainId", nil)
	cdc.RegisterConcrete(&MsgSetDenomType{}, "bridge/SetDenomType", nil)
	cdc.RegisterConcrete(&MsgVoteProposal{}, "bridge/VoteProposal", nil)
	cdc.RegisterConcrete(&MsgRmChainId{}, "bridge/RmChainId", nil)
	cdc.RegisterConcrete(&MsgSetRelayFeeReceiver{}, "bridge/SetRelayFeeReceiver", nil)
	cdc.RegisterConcrete(&MsgSetRelayFee{}, "bridge/SetRelayFee", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetResourceidToDenom{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeposit{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddChainId{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetDenomType{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgVoteProposal{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRmChainId{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetRelayFeeReceiver{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetRelayFee{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
