package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgProvideToken = "provide_token"

var _ sdk.Msg = &MsgProvideToken{}

func NewMsgProvideToken(creator string, token sdk.Coin) *MsgProvideToken {
	return &MsgProvideToken{
		Creator: creator,
		Token:   token,
	}
}

func (msg *MsgProvideToken) Route() string {
	return RouterKey
}

func (msg *MsgProvideToken) Type() string {
	return TypeMsgProvideToken
}

func (msg *MsgProvideToken) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgProvideToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgProvideToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
