package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetMaxRewardPoolNumber = "set_max_reward_pool_number"

var _ sdk.Msg = &MsgSetMaxRewardPoolNumber{}

func NewMsgSetMaxRewardPoolNumber(creator string, number uint32) *MsgSetMaxRewardPoolNumber {
	return &MsgSetMaxRewardPoolNumber{
		Creator: creator,
		Number:  number,
	}
}

func (msg *MsgSetMaxRewardPoolNumber) Route() string {
	return RouterKey
}

func (msg *MsgSetMaxRewardPoolNumber) Type() string {
	return TypeMsgSetMaxRewardPoolNumber
}

func (msg *MsgSetMaxRewardPoolNumber) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetMaxRewardPoolNumber) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetMaxRewardPoolNumber) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
