package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgToggleWhitelistSwitch = "toggle_whitelist_switch"

var _ sdk.Msg = &MsgToggleWhitelistSwitch{}

func NewMsgToggleWhitelistSwitch(creator string) *MsgToggleWhitelistSwitch {
	return &MsgToggleWhitelistSwitch{
		Creator: creator,
	}
}

func (msg *MsgToggleWhitelistSwitch) Route() string {
	return RouterKey
}

func (msg *MsgToggleWhitelistSwitch) Type() string {
	return TypeMsgToggleWhitelistSwitch
}

func (msg *MsgToggleWhitelistSwitch) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgToggleWhitelistSwitch) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgToggleWhitelistSwitch) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
