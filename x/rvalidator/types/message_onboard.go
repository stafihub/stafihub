package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgOnboard{}

func NewMsgOnboard(creator sdk.AccAddress, denom string, address string, locked sdk.Coin) *MsgOnboard {
	return &MsgOnboard{
		Creator: creator.String(),
		Denom:   denom,
		Address: address,
		Locked:  locked,
	}
}

func (msg *MsgOnboard) Route() string {
	return RouterKey
}

func (msg *MsgOnboard) Type() string {
	return "Onboard"
}

func (msg *MsgOnboard) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgOnboard) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgOnboard) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
