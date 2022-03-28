package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgProvideRewardToken = "provide_reward_token"

var _ sdk.Msg = &MsgProvideRewardToken{}

func NewMsgProvideRewardToken(creator string, amount sdk.Coin) *MsgProvideRewardToken {
	return &MsgProvideRewardToken{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgProvideRewardToken) Route() string {
	return RouterKey
}

func (msg *MsgProvideRewardToken) Type() string {
	return TypeMsgProvideRewardToken
}

func (msg *MsgProvideRewardToken) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgProvideRewardToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgProvideRewardToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
