package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateThreshold{}

func NewMsgUpdateThreshold(
    creator string,
    denom string,
    value uint32,
) *MsgUpdateThreshold {
  return &MsgUpdateThreshold{
		Creator: creator,
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
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateThreshold) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateThreshold) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  }
   return nil
}