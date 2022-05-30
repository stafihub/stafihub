package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRmStakeToken = "rm_stake_token"

var _ sdk.Msg = &MsgRmStakeToken{}

func NewMsgRmStakeToken(creator string, denom string) *MsgRmStakeToken {
	return &MsgRmStakeToken{
		Creator: creator,
		Denom:   denom,
	}
}

func (msg *MsgRmStakeToken) Route() string {
	return RouterKey
}

func (msg *MsgRmStakeToken) Type() string {
	return TypeMsgRmStakeToken
}

func (msg *MsgRmStakeToken) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRmStakeToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRmStakeToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if err := sdk.ValidateDenom(msg.Denom); err != nil {
		return err
	}
	return nil
}
