package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgToggleUnbondSwitch = "toggle_unbond_switch"

var _ sdk.Msg = &MsgToggleUnbondSwitch{}

func NewMsgToggleUnbondSwitch(creator, denom string) *MsgToggleUnbondSwitch {
	return &MsgToggleUnbondSwitch{
		Creator: creator,
		Denom:   denom,
	}
}

func (msg *MsgToggleUnbondSwitch) Route() string {
	return RouterKey
}

func (msg *MsgToggleUnbondSwitch) Type() string {
	return TypeMsgToggleUnbondSwitch
}

func (msg *MsgToggleUnbondSwitch) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgToggleUnbondSwitch) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgToggleUnbondSwitch) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if len(msg.Denom) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "denom cannot be empty")
	}
	return nil
}
