package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgAddStakePool{}, "mining/AddStakePool", nil)
	cdc.RegisterConcrete(&MsgAddStakeItem{}, "mining/AddStakeItem", nil)
	cdc.RegisterConcrete(&MsgAddRewardPool{}, "mining/AddRewardPool", nil)
	cdc.RegisterConcrete(&MsgStake{}, "mining/Stake", nil)
	cdc.RegisterConcrete(&MsgClaimReward{}, "mining/ClaimReward", nil)
	cdc.RegisterConcrete(&MsgWithdraw{}, "mining/Withdraw", nil)
	cdc.RegisterConcrete(&MsgProvideRewardToken{}, "mining/ProvideRewardToken", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddStakePool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddStakeItem{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddRewardPool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgStake{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgClaimReward{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgWithdraw{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgProvideRewardToken{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
