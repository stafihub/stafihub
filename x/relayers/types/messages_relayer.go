package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateRelayer{}

func NewMsgCreateRelayer(
    creator string,
    denom string,
    address sdk.AccAddress,
) *MsgCreateRelayer {
  return &MsgCreateRelayer{
		Creator : creator,
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
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgCreateRelayer) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateRelayer) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  	if err != nil {
  		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  	}

	if msg.Address == "" {
		return ErrEmptyRelayerAddr
	}
  return nil
}

var _ sdk.Msg = &MsgDeleteRelayer{}

func NewMsgDeleteRelayer(
    creator string,
    denom string,
    address sdk.AccAddress,
) *MsgDeleteRelayer {
  return &MsgDeleteRelayer{
		Creator: creator,
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
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteRelayer) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteRelayer) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  }

  if msg.Address == "" {
	  return ErrEmptyRelayerAddr
  }

  return nil
}