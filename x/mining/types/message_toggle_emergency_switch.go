package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgToggleEmergencySwitch = "toggle_emergency_switch"

var _ sdk.Msg = &MsgToggleEmergencySwitch{}

func NewMsgToggleEmergencySwitch(creator string) *MsgToggleEmergencySwitch {
	return &MsgToggleEmergencySwitch{
		Creator: creator,
	}
}

func (msg *MsgToggleEmergencySwitch) Route() string {
	return RouterKey
}

func (msg *MsgToggleEmergencySwitch) Type() string {
	return TypeMsgToggleEmergencySwitch
}

func (msg *MsgToggleEmergencySwitch) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgToggleEmergencySwitch) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgToggleEmergencySwitch) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
