package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateThreshold{}

func NewMsgUpdateThreshold(creator sdk.AccAddress, denom string, value uint32) *MsgUpdateThreshold {
  return &MsgUpdateThreshold{
		Creator: creator.String(),
        Denom: denom,
        Value: value,
	}
}

func (msg *MsgUpdateThreshold) Route() string {
  return RouterKey
}

func (msg *MsgUpdateThreshold) Type() string {
  return "UpdateThreshold"
}

func (msg *MsgUpdateThreshold) GetSigners() []sdk.AccAddress {
  creator, _ := sdk.AccAddressFromBech32(msg.Creator)
  return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateThreshold) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateThreshold) ValidateBasic() error {
  if msg.Creator == "" {
    return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address")
  }

  return nil
}