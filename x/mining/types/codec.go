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
	cdc.RegisterConcrete(&MsgUpdateStakeItem{}, "mining/UpdateStakeItem", nil)
	cdc.RegisterConcrete(&MsgAddMiningProvider{}, "mining/AddMiningProvider", nil)
	cdc.RegisterConcrete(&MsgRmMiningProvider{}, "mining/RmMiningProvider", nil)
	cdc.RegisterConcrete(&MsgAddRewardToken{}, "mining/AddRewardToken", nil)
	cdc.RegisterConcrete(&MsgSetMaxRewardPoolNumber{}, "mining/SetMaxRewardPoolNumber", nil)
	cdc.RegisterConcrete(&MsgUpdateRewardPool{}, "mining/UpdateRewardPool", nil)
	cdc.RegisterConcrete(&MsgToggleProviderSwitch{}, "mining/ToggleProviderSwitch", nil)
	cdc.RegisterConcrete(&MsgSetMaxStakeItemNumber{}, "mining/SetMaxStakeItemNumber", nil)
	cdc.RegisterConcrete(&MsgAddReward{}, "mining/AddReward", nil)
	cdc.RegisterConcrete(&MsgToggleEmergencySwitch{}, "mining/ToggleEmergencySwitch", nil)
	cdc.RegisterConcrete(&MsgEmergencyWithdraw{}, "mining/EmergencyWithdraw", nil)
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
		&MsgUpdateStakeItem{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddMiningProvider{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRmMiningProvider{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddRewardToken{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetMaxRewardPoolNumber{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateRewardPool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgToggleProviderSwitch{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetMaxStakeItemNumber{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddReward{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgToggleEmergencySwitch{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgEmergencyWithdraw{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
