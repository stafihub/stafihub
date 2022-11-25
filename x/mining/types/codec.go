package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
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
	cdc.RegisterConcrete(&MsgAddStakeToken{}, "mining/AddStakeToken", nil)
	cdc.RegisterConcrete(&MsgRmRewardToken{}, "mining/RmRewardToken", nil)
	cdc.RegisterConcrete(&MsgRmStakeToken{}, "mining/RmStakeToken", nil)
	cdc.RegisterConcrete(&MsgSetStakeItemLimit{}, "mining/SetStakeItemLimit", nil)
	cdc.RegisterConcrete(&MsgWithdrawRewardToken{}, "mining/WithdrawRewardToken", nil)
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
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddStakeToken{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRmRewardToken{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRmStakeToken{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetStakeItemLimit{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgWithdrawRewardToken{},
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
