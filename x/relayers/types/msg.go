package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgCreateRelayer{}
	_ sdk.Msg = &MsgDeleteRelayer{}
	_ sdk.Msg = &MsgUpdateThreshold{}
)

func NewMsgCreateRelayer(creator sdk.AccAddress, taipe, denom string, addresses []string) *MsgCreateRelayer {
	return &MsgCreateRelayer{
		Creator:   creator.String(),
		Taipe: taipe,
		Denom:     denom,
		Addresses: addresses,
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
	if msg.Creator == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}

	if len(msg.Addresses) == 0 {
		return fmt.Errorf("Addresses should not be empty")
	}

	for _, addr := range msg.Addresses {
		if _, err := sdk.AccAddressFromBech32(addr); err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid relayer address %s", addr)
		}
	}
	return nil
}

func NewMsgDeleteRelayer(creator sdk.AccAddress, taipe, denom string, address sdk.AccAddress) *MsgDeleteRelayer {
	return &MsgDeleteRelayer{
		Creator: creator.String(),
		Taipe:taipe,
		Denom:   denom,
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
}

func NewMsgUpdateThreshold(creator sdk.AccAddress, taipe, denom string, value uint32) *MsgUpdateThreshold {
	return &MsgUpdateThreshold{
		Creator: creator.String(),
		Taipe: taipe,
		Denom:   denom,
		Value:   value,
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
