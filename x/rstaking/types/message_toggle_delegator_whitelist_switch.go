package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgToggleDelegatorWhitelistSwitch = "toggle_delegator_whitelist_switch"

var _ sdk.Msg = &MsgToggleDelegatorWhitelistSwitch{}

func NewMsgToggleDelegatorWhitelistSwitch(creator string) *MsgToggleDelegatorWhitelistSwitch {
	return &MsgToggleDelegatorWhitelistSwitch{
		Creator: creator,
	}
}

func (msg *MsgToggleDelegatorWhitelistSwitch) Route() string {
	return RouterKey
}

func (msg *MsgToggleDelegatorWhitelistSwitch) Type() string {
	return TypeMsgToggleDelegatorWhitelistSwitch
}

func (msg *MsgToggleDelegatorWhitelistSwitch) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgToggleDelegatorWhitelistSwitch) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgToggleDelegatorWhitelistSwitch) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
