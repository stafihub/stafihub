package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetThreshold = "set_threshold"

var _ sdk.Msg = &MsgSetThreshold{}

func NewMsgSetThreshold(creator string, threshold uint32) *MsgSetThreshold {
	return &MsgSetThreshold{
		Creator:   creator,
		Threshold: threshold,
	}
}

func (msg *MsgSetThreshold) Route() string {
	return RouterKey
}

func (msg *MsgSetThreshold) Type() string {
	return TypeMsgSetThreshold
}

func (msg *MsgSetThreshold) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetThreshold) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetThreshold) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
