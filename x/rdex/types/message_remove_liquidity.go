package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRemoveLiquidity = "remove_liquidity"

var _ sdk.Msg = &MsgRemoveLiquidity{}

func NewMsgRemoveLiquidity(creator string, rmUnit, swapUnit sdk.Int, minOutTokens sdk.Coins, inputTokenDenom string) *MsgRemoveLiquidity {
	return &MsgRemoveLiquidity{
		Creator:         creator,
		RmUnit:          rmUnit,
		SwapUnit:        swapUnit,
		MinOutTokens:    minOutTokens,
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
	if len(msg.MinOutTokens) != 2 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid tokens length(%d)", len(msg.MinOutTokens))
	}

	if !msg.RmUnit.IsPositive() {
		return fmt.Errorf("invalid rm unit amount")
	}
	if msg.RmUnit.LT(msg.SwapUnit) {
		return fmt.Errorf("rm unit must bigger or equal to swap unit")
	}

	return nil
}
