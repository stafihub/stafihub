package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddStakeToken = "add_stake_token"

var _ sdk.Msg = &MsgAddStakeToken{}

func NewMsgAddStakeToken(creator string, denom string) *MsgAddStakeToken {
	return &MsgAddStakeToken{
		Creator: creator,
		Denom:   denom,
	}
}

func (msg *MsgAddStakeToken) Route() string {
	return RouterKey
}

func (msg *MsgAddStakeToken) Type() string {
	return TypeMsgAddStakeToken
}

func (msg *MsgAddStakeToken) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddStakeToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddStakeToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if err := sdk.ValidateDenom(msg.Denom); err != nil {
		return err
	}
	return nil
}
