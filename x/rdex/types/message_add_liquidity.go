package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddLiquidity = "add_liquidity"

var _ sdk.Msg = &MsgAddLiquidity{}

func NewMsgAddLiquidity(creator string, swapPoolIndex uint32, token0, token1 sdk.Coin) *MsgAddLiquidity {
	return &MsgAddLiquidity{
		Creator:       creator,
		SwapPoolIndex: swapPoolIndex,
		Token0:        token0,
		Token1:        token1,
	}
}

func (msg *MsgAddLiquidity) Route() string {
	return RouterKey
}

func (msg *MsgAddLiquidity) Type() string {
	return TypeMsgAddLiquidity
}

func (msg *MsgAddLiquidity) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddLiquidity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddLiquidity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if !msg.Token0.IsValid() || !msg.Token1.IsValid() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid token token0(%s) token1(%s)", msg.Token0, msg.Token1)
	}
	if strings.EqualFold(msg.Token0.Denom, msg.Token1.Denom) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid token denom token0(%s) token1(%s)", msg.Token0, msg.Token1)
	}

	if !msg.Token0.IsPositive() && !msg.Token1.IsPositive() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid token amount token0(%s) token1(%s)", msg.Token0, msg.Token1)
	}

	return nil
}
