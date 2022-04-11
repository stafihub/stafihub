package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgAddRelayer{}
	_ sdk.Msg = &MsgDeleteRelayer{}
	_ sdk.Msg = &MsgSetThreshold{}
)

func NewMsgCreateRelayer(creator sdk.AccAddress, arena, denom string, addresses []string) *MsgAddRelayer {
	return &MsgAddRelayer{
		Creator:   creator.String(),
		Arena:     arena,
		Denom:     denom,
		Addresses: addresses,
	}
}

func (msg *MsgAddRelayer) Route() string {
	return RouterKey
}

func (msg *MsgAddRelayer) Type() string {
	return "AddRelayer"
}

func (msg *MsgAddRelayer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic("invalid creator address")
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddRelayer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddRelayer) ValidateBasic() error {
	if msg.Creator == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}

	if len(msg.Addresses) == 0 {
		return fmt.Errorf("Addresses should not be empty")
	}
	if len(msg.Arena) == 0 {
		return fmt.Errorf("Arena should not be empty")
	}

	for _, addr := range msg.Addresses {
		if _, err := sdk.AccAddressFromBech32(addr); err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid relayer address %s", addr)
		}
	}
	return nil
}

func NewMsgDeleteRelayer(creator sdk.AccAddress, arena, denom string, address sdk.AccAddress) *MsgDeleteRelayer {
	return &MsgDeleteRelayer{
		Creator: creator.String(),
		Arena:   arena,
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
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic("invalid creator address")
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteRelayer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteRelayer) ValidateBasic() error {
	if len(msg.Arena) == 0 {
		return fmt.Errorf("Arena should not be empty")
	}
	if msg.Creator == "" || msg.Address == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator (%s) or address (%s)", msg.Creator, msg.Address)
	}
	return nil
}

func NewMsgSetThreshold(creator sdk.AccAddress, arena, denom string, value uint32) *MsgSetThreshold {
	return &MsgSetThreshold{
		Creator: creator.String(),
		Arena:   arena,
		Denom:   denom,
		Value:   value,
	}
}

func (msg *MsgSetThreshold) Route() string {
	return RouterKey
}

func (msg *MsgSetThreshold) Type() string {
	return "SetThreshold"
}

func (msg *MsgSetThreshold) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic("invalid creator address")
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetThreshold) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetThreshold) ValidateBasic() error {
	if msg.Creator == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}
	if len(msg.Arena) == 0 {
		return fmt.Errorf("Arena should not be empty")
	}
	if msg.Value <= 0 {
		return fmt.Errorf("threshold should be greater than 0")
	}

	return nil
}
