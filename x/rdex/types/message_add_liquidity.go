package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddLiquidity = "add_liquidity"

var _ sdk.Msg = &MsgAddLiquidity{}

func NewMsgAddLiquidity(creator string, denom string, rTokenAmount, fisAmount sdk.Int) *MsgAddLiquidity {
	return &MsgAddLiquidity{
		Creator:      creator,
		Denom:        denom,
		RTokenAmount: rTokenAmount,
		FisAmount:    fisAmount,
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
	if msg.RTokenAmount.LT(sdk.ZeroInt()) || msg.FisAmount.LT(sdk.ZeroInt()) {
		return fmt.Errorf("invalid token amount")
	}

	if msg.RTokenAmount.Equal(sdk.ZeroInt()) && msg.FisAmount.Equal(sdk.ZeroInt()) {
		return fmt.Errorf("token amount all zero error")
	}
	if len(msg.Denom) == 0 {
		return fmt.Errorf("invalid denom")
	}
	return nil
}
