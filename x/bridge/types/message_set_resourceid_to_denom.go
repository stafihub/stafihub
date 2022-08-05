package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetResourceidToDenom = "set_resourceid_to_denom"

var _ sdk.Msg = &MsgSetResourceidToDenom{}

func NewMsgSetResourceidToDenom(creator string, resourceId string, denom string, denomType DenomType) *MsgSetResourceidToDenom {
	return &MsgSetResourceidToDenom{
		Creator:    creator,
		ResourceId: resourceId,
		Denom:      denom,
		DenomType:  denomType,
	}
}

func (msg *MsgSetResourceidToDenom) Route() string {
	return RouterKey
}

func (msg *MsgSetResourceidToDenom) Type() string {
	return TypeMsgSetResourceidToDenom
}

func (msg *MsgSetResourceidToDenom) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetResourceidToDenom) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetResourceidToDenom) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
