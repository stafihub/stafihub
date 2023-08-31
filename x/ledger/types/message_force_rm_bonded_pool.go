package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgForceRmBondedPool = "force_rm_bonded_pool"

var _ sdk.Msg = &MsgForceRmBondedPool{}

func NewMsgForceRmBondedPool(creator string, denom string, address string) *MsgForceRmBondedPool {
	return &MsgForceRmBondedPool{
		Creator: creator,
		Denom:   denom,
		Address: address,
	}
}

func (msg *MsgForceRmBondedPool) Route() string {
	return RouterKey
}

func (msg *MsgForceRmBondedPool) Type() string {
	return TypeMsgForceRmBondedPool
}

func (msg *MsgForceRmBondedPool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgForceRmBondedPool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgForceRmBondedPool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
