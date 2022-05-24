package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgToggleProviderSwitch = "toggle_provider_switch"

var _ sdk.Msg = &MsgToggleProviderSwitch{}

func NewMsgToggleProviderSwitch(creator string) *MsgToggleProviderSwitch {
	return &MsgToggleProviderSwitch{
		Creator: creator,
	}
}

func (msg *MsgToggleProviderSwitch) Route() string {
	return RouterKey
}

func (msg *MsgToggleProviderSwitch) Type() string {
	return TypeMsgToggleProviderSwitch
}

func (msg *MsgToggleProviderSwitch) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgToggleProviderSwitch) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgToggleProviderSwitch) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
