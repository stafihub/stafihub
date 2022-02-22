package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetResourceidType = "set_resourceid_type"

var _ sdk.Msg = &MsgSetResourceidType{}

func NewMsgSetResourceidType(creator string, resourceId string, idType string) *MsgSetResourceidType {
	return &MsgSetResourceidType{
		Creator:    creator,
		ResourceId: resourceId,
		IdType:     idType,
	}
}

func (msg *MsgSetResourceidType) Route() string {
	return RouterKey
}

func (msg *MsgSetResourceidType) Type() string {
	return TypeMsgSetResourceidType
}

func (msg *MsgSetResourceidType) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetResourceidType) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetResourceidType) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
