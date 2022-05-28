package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRemoveLiquidity = "remove_liquidity"

var _ sdk.Msg = &MsgRemoveLiquidity{}

func NewMsgRemoveLiquidity(creator string, swapPoolIndex uint32, rmUnit, swapUnit sdk.Int, minOutToken0, minOutToken1 sdk.Coin, inputTokenDenom string) *MsgRemoveLiquidity {
	return &MsgRemoveLiquidity{
		Creator:         creator,
		SwapPoolIndex:   swapPoolIndex,
		RmUnit:          rmUnit,
		SwapUnit:        swapUnit,
		MinOutToken0:    minOutToken0,
		MinOutToken1:    minOutToken1,
		InputTokenDenom: inputTokenDenom,
	}
}

func (msg *MsgRemoveLiquidity) Route() string {
	return RouterKey
}

func (msg *MsgRemoveLiquidity) Type() string {
	return TypeMsgRemoveLiquidity
}

func (msg *MsgRemoveLiquidity) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemoveLiquidity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveLiquidity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if !msg.RmUnit.IsPositive() {
		return fmt.Errorf("invalid rm unit amount")
	}
	if msg.SwapUnit.IsNegative() {
		return fmt.Errorf("invalid swap unit amount")
	}

	if msg.RmUnit.LT(msg.SwapUnit) {
		return fmt.Errorf("rm unit must bigger or equal to swap unit")
	}

	if !msg.MinOutToken0.IsValid() || !msg.MinOutToken1.IsValid() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid token minOutToken0(%s) minOutToken1(%s)", msg.MinOutToken0, msg.MinOutToken1)
	}

	if strings.EqualFold(msg.MinOutToken0.Denom, msg.MinOutToken1.Denom) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid token denom minOutToken0(%s) minOutToken1(%s)", msg.MinOutToken0, msg.MinOutToken1)
	}

	if !msg.SwapUnit.IsZero() {
		if msg.InputTokenDenom != msg.MinOutToken0.Denom && msg.InputTokenDenom != msg.MinOutToken1.Denom {
			return fmt.Errorf("inputTokenDenom err")
		}
	}
	return nil
}
