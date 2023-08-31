package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgForceUpdateBondedPool = "force_update_bonded_pool"

var _ sdk.Msg = &MsgForceUpdateBondedPool{}

func NewMsgForceUpdateBondedPool(creator string, denom string, address string, active, bond, unbond sdk.Int) *MsgForceUpdateBondedPool {
	return &MsgForceUpdateBondedPool{
		Creator: creator,
		Denom:   denom,
		Address: address,
		Active:  active,
		Bond:    bond,
		Unbond:  unbond,
	}
}

func (msg *MsgForceUpdateBondedPool) Route() string {
	return RouterKey
}

func (msg *MsgForceUpdateBondedPool) Type() string {
	return TypeMsgForceUpdateBondedPool
}

func (msg *MsgForceUpdateBondedPool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgForceUpdateBondedPool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgForceUpdateBondedPool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
