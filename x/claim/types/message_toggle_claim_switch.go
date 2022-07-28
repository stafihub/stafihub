package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgToggleClaimSwitch = "toggle_claim_switch"

var _ sdk.Msg = &MsgToggleClaimSwitch{}

func NewMsgToggleClaimSwitch(creator string, round uint64) *MsgToggleClaimSwitch {
	return &MsgToggleClaimSwitch{
		Creator: creator,
		Round:   round,
	}
}

func (msg *MsgToggleClaimSwitch) Route() string {
	return RouterKey
}

func (msg *MsgToggleClaimSwitch) Type() string {
	return TypeMsgToggleClaimSwitch
}

func (msg *MsgToggleClaimSwitch) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgToggleClaimSwitch) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgToggleClaimSwitch) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
