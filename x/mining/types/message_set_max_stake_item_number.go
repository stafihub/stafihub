package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetMaxStakeItemNumber = "set_max_stake_item_number"

var _ sdk.Msg = &MsgSetMaxStakeItemNumber{}

func NewMsgSetMaxStakeItemNumber(creator string, number uint32) *MsgSetMaxStakeItemNumber {
	return &MsgSetMaxStakeItemNumber{
		Creator: creator,
		Number:  number,
	}
}

func (msg *MsgSetMaxStakeItemNumber) Route() string {
	return RouterKey
}

func (msg *MsgSetMaxStakeItemNumber) Type() string {
	return TypeMsgSetMaxStakeItemNumber
}

func (msg *MsgSetMaxStakeItemNumber) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetMaxStakeItemNumber) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetMaxStakeItemNumber) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
