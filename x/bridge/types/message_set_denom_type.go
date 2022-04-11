package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetDenomType = "set_denom_type"

var _ sdk.Msg = &MsgSetDenomType{}

func NewMsgSetDenomType(creator string, denom string, idType string) *MsgSetDenomType {
	return &MsgSetDenomType{
		Creator: creator,
		Denom:   denom,
		IdType:  idType,
	}
}

func (msg *MsgSetDenomType) Route() string {
	return RouterKey
}

func (msg *MsgSetDenomType) Type() string {
	return TypeMsgSetDenomType
}

func (msg *MsgSetDenomType) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetDenomType) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetDenomType) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
