package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetChainEra{}

func NewMsgSetChainEra(creator string, denom string, newEra string) *MsgSetChainEra {
  return &MsgSetChainEra{
		Creator: creator,
    Denom: denom,
    NewEra: newEra,
	}
}

func (msg *MsgSetChainEra) Route() string {
  return RouterKey
}

func (msg *MsgSetChainEra) Type() string {
  return "SetChainEra"
}

func (msg *MsgSetChainEra) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgSetChainEra) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgSetChainEra) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  	if err != nil {
  		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  	}
  return nil
}

