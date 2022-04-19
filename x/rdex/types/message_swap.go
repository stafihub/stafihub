package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSwap = "swap"

var _ sdk.Msg = &MsgSwap{}

func NewMsgSwap(creator string, denom string, inputAmount, minOutAmount sdk.Int, inputIsFis bool) *MsgSwap {
	return &MsgSwap{
		Creator:      creator,
		Denom:        denom,
		InputAmount:  inputAmount,
		MinOutAmount: minOutAmount,
		InputIsFis:   inputIsFis,
	}
}

func (msg *MsgSwap) Route() string {
	return RouterKey
}

func (msg *MsgSwap) Type() string {
	return TypeMsgSwap
}

func (msg *MsgSwap) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSwap) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSwap) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.InputAmount.LTE(sdk.ZeroInt()) || msg.MinOutAmount.LTE(sdk.ZeroInt()) {
		return fmt.Errorf("invalid amount")
	}
	if len(msg.Denom) == 0 {
		return fmt.Errorf("invalid denom")
	}
	return nil
}
