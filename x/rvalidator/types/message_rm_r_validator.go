package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRmRValidator = "rm_r_validator"

var _ sdk.Msg = &MsgRmRValidator{}

func NewMsgRmRValidator(creator string, denom string, poolAddress string, oldAddress string, newAddress string) *MsgRmRValidator {
	return &MsgRmRValidator{
		Creator:     creator,
		Denom:       denom,
		PoolAddress: poolAddress,
		OldAddress:  oldAddress,
		NewAddress:  newAddress,
	}
}

func (msg *MsgRmRValidator) Route() string {
	return RouterKey
}

func (msg *MsgRmRValidator) Type() string {
	return TypeMsgRmRValidator
}

func (msg *MsgRmRValidator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRmRValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRmRValidator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
