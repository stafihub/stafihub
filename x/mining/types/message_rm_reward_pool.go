package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRmRewardPool = "rm_reward_pool"

var _ sdk.Msg = &MsgRmRewardPool{}

func NewMsgRmRewardPool(creator string, stakePoolIndex, rewardPoolIndex uint32) *MsgRmRewardPool {
	return &MsgRmRewardPool{
		Creator:         creator,
		StakePoolIndex:  stakePoolIndex,
		RewardPoolIndex: rewardPoolIndex,
	}
}

func (msg *MsgRmRewardPool) Route() string {
	return RouterKey
}

func (msg *MsgRmRewardPool) Type() string {
	return TypeMsgRmRewardPool
}

func (msg *MsgRmRewardPool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRmRewardPool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRmRewardPool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
