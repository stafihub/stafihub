package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSwap = "swap"

var _ sdk.Msg = &MsgSwap{}

func NewMsgSwap(creator string, inputToken, minOutToken sdk.Coin) *MsgSwap {
	return &MsgSwap{
		Creator:     creator,
		InputToken:  inputToken,
		MinOutToken: minOutToken,
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
	if msg.InputToken.Amount.LTE(sdk.ZeroInt()) {
		sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid input token amount (%s)", msg.InputToken.Amount)
	}
	if msg.MinOutToken.Amount.LTE(sdk.ZeroInt()) {
		sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid input token amount (%s)", msg.InputToken.Amount)
	}
	return nil
}
