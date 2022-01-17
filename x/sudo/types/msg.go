package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

var (
	_ sdk.Msg = &MsgUpdateAdmin{}
	_ sdk.Msg = &MsgAddDenom{}
)

func NewMsgUpdateAdmin(creator sdk.AccAddress, address sdk.AccAddress) *MsgUpdateAdmin {
  return &MsgUpdateAdmin{
		Creator: creator.String(),
    	Address: address.String(),
	}
}

func (msg *MsgUpdateAdmin) Route() string {
  return RouterKey
}

func (msg *MsgUpdateAdmin) Type() string {
  return "UpdateAdmin"
}

func (msg *MsgUpdateAdmin) GetSigners() []sdk.AccAddress {
  creator, _ := sdk.AccAddressFromBech32(msg.Creator)
  return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateAdmin) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateAdmin) ValidateBasic() error {
  if msg.Creator == "" || msg.Address == "" {
	  return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator (%s) or address (%s)", msg.Creator, msg.Address)
  }
  return nil
}

func NewMsgAddDenom(creator sdk.AccAddress, metadata banktypes.Metadata) *MsgAddDenom {
	return &MsgAddDenom{
		Creator: creator.String(),
		Metadata: metadata,
	}
}

func (msg *MsgAddDenom) Route() string {
	return RouterKey
}

func (msg *MsgAddDenom) Type() string {
	return "AddDenom"
}

func (msg *MsgAddDenom) GetSigners() []sdk.AccAddress {
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddDenom) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddDenom) ValidateBasic() error {
	if msg.Creator == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator (%s)", msg.Creator)
	}
	return nil
}

