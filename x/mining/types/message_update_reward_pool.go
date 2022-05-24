package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateRewardPool = "update_reward_pool"

var _ sdk.Msg = &MsgUpdateRewardPool{}

func NewMsgUpdateRewardPool(creator string, stakePoolIndex, rewardPoolIndex uint32, newRewardAmount sdk.Int) *MsgUpdateRewardPool {
	return &MsgUpdateRewardPool{
		Creator:         creator,
		StakePoolIndex:  stakePoolIndex,
		RewardPoolIndex: rewardPoolIndex,
		NewRewardAmount: newRewardAmount,
	}
}

func (msg *MsgUpdateRewardPool) Route() string {
	return RouterKey
}

func (msg *MsgUpdateRewardPool) Type() string {
	return TypeMsgUpdateRewardPool
}

func (msg *MsgUpdateRewardPool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateRewardPool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateRewardPool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
