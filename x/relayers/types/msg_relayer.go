package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgCreateRelayer{}
	_ sdk.Msg = &MsgDeleteRelayer{}
)

func NewMsgCreateRelayer(creator sdk.AccAddress, denom string, address sdk.AccAddress) *MsgCreateRelayer {
  return &MsgCreateRelayer{
		Creator : creator.String(),
		Denom: denom,
        Address: address.String(),
	}
}

func (msg *MsgCreateRelayer) Route() string {
  return RouterKey
}

func (msg *MsgCreateRelayer) Type() string {
  return "CreateRelayer"
}

func (msg *MsgCreateRelayer) GetSigners() []sdk.AccAddress {
  creator, _ := sdk.AccAddressFromBech32(msg.Creator)
  return []sdk.AccAddress{creator}
}

func (msg *MsgCreateRelayer) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateRelayer) ValidateBasic() error {
	if msg.Creator == "" || msg.Address == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator (%s) or address (%s)", msg.Creator, msg.Address)
	}
  return nil
}

func NewMsgDeleteRelayer(creator sdk.AccAddress, denom string, address sdk.AccAddress) *MsgDeleteRelayer {
  return &MsgDeleteRelayer{
		Creator: creator.String(),
		Denom: denom,
		Address: address.String(),
	}
}
func (msg *MsgDeleteRelayer) Route() string {
  return RouterKey
}

func (msg *MsgDeleteRelayer) Type() string {
  return "DeleteRelayer"
}

func (msg *MsgDeleteRelayer) GetSigners() []sdk.AccAddress {
  creator, _ := sdk.AccAddressFromBech32(msg.Creator)
  return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteRelayer) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteRelayer) ValidateBasic() error {
	if msg.Creator == "" || msg.Address == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator (%s) or address (%s)", msg.Creator, msg.Address)
	}
	return nil

  return nil
}