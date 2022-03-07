package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetRValidatorIndicator{}

func NewMsgSetRValidatorIndicator(creator sdk.AccAddress, denom string, commission sdk.Dec, uptime uint32, locked sdk.Coin) *MsgSetRValidatorIndicator {
	return &MsgSetRValidatorIndicator{
		Creator:    creator.String(),
		Denom:      denom,
		Commission: commission,
		Uptime:     uptime,
		Locked:     locked,
	}
}

func (msg *MsgSetRValidatorIndicator) Route() string {
	return RouterKey
}

func (msg *MsgSetRValidatorIndicator) Type() string {
	return "SetRValidatorIndicator"
}

func (msg *MsgSetRValidatorIndicator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetRValidatorIndicator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetRValidatorIndicator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
