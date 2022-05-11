package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreatePool = "create_pool"

var _ sdk.Msg = &MsgCreatePool{}

func NewMsgCreatePool(creator string, tokens sdk.Coins) *MsgCreatePool {
	return &MsgCreatePool{
		Creator: creator,
		Tokens:  tokens,
	}
}

func (msg *MsgCreatePool) Route() string {
	return RouterKey
}

func (msg *MsgCreatePool) Type() string {
	return TypeMsgCreatePool
}

func (msg *MsgCreatePool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreatePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreatePool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if len(msg.Tokens) != 2 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid tokens length (%s)", len(msg.Tokens))
	}
	for _, token := range msg.Tokens {
		if token.Amount.LTE(sdk.ZeroInt()) {
			sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid token amount (%s)", token.Amount)
		}
	}

	return nil
}
