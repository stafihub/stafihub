package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRmBondedPool = "rm_bonded_pool"

var _ sdk.Msg = &MsgRmBondedPool{}

func NewMsgRmBondedPool(creator string, denom string, address string) *MsgRmBondedPool {
	return &MsgRmBondedPool{
		Creator: creator,
		Denom:   denom,
		Address: address,
	}
}

func (msg *MsgRmBondedPool) Route() string {
	return RouterKey
}

func (msg *MsgRmBondedPool) Type() string {
	return TypeMsgRmBondedPool
}

func (msg *MsgRmBondedPool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRmBondedPool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRmBondedPool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
