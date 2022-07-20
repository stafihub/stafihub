package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgInitPool = "init_pool"

var _ sdk.Msg = &MsgInitPool{}

func NewMsgInitPool(creator string, denom string, pool string) *MsgInitPool {
	return &MsgInitPool{
		Creator: creator,
		Denom:   denom,
		Pool:    pool,
	}
}

func (msg *MsgInitPool) Route() string {
	return RouterKey
}

func (msg *MsgInitPool) Type() string {
	return TypeMsgInitPool
}

func (msg *MsgInitPool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgInitPool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInitPool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
