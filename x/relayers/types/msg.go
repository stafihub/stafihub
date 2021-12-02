package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"
)

var (
	_ sdk.Msg = &MsgCreateRelayer{}
	_ sdk.Msg = &MsgDeleteRelayer{}
	_ sdk.Msg = &MsgSetThreshold{}
)


func NewMsgCreateRelayer(denom string, address string) *MsgCreateRelayer {
  return &MsgCreateRelayer{
    Denom: denom,
    Address: address,
	}
}

func (msg *MsgCreateRelayer) Route() string {
  return RouterKey
}

func (msg *MsgCreateRelayer) Type() string {
  return "CreateRelayer"
}

func (msg *MsgCreateRelayer) GetSigners() []sdk.AccAddress {
  gs := DefaultGenesis()
  admin, err := sdk.AccAddressFromBech32(gs.Admin)
  	if err != nil {
		panic(err)
	}
  return []sdk.AccAddress{admin}
}

func (msg *MsgCreateRelayer) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateRelayer) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Address)
  	if err != nil {
  		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid relayer address (%s)", err)
  	}
  return nil
}

func NewMsgDeleteRelayer(denom string, address string) *MsgDeleteRelayer {
	return &MsgDeleteRelayer{
		Denom: denom,
		Address: address,
	}
}

func (msg *MsgDeleteRelayer) Route() string {
	return RouterKey
}

func (msg *MsgDeleteRelayer) Type() string {
	return "DeleteRelayer"
}

func (msg *MsgDeleteRelayer) GetSigners() []sdk.AccAddress {
	gs := DefaultGenesis()
	admin, err := sdk.AccAddressFromBech32(gs.Admin)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{admin}
}

func (msg *MsgDeleteRelayer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteRelayer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid relayer address (%s)", err)
	}
	return nil
}

func NewMsgSetThreshold(denom string, value string) *MsgSetThreshold {
	return &MsgSetThreshold{
		Denom: denom,
		Value: value,
	}
}

func (msg *MsgSetThreshold) Route() string {
	return RouterKey
}

func (msg *MsgSetThreshold) Type() string {
	return "SetThreshold"
}

func (msg *MsgSetThreshold) GetSigners() []sdk.AccAddress {
	gs := DefaultGenesis()
	admin, err := sdk.AccAddressFromBech32(gs.Admin)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{admin}
}

func (msg *MsgSetThreshold) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetThreshold) ValidateBasic() error {
	_, err := strconv.ParseUint(msg.Value, 10, 64)
	if err != nil {
		return err
	}

	return nil
}

