package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRmDelegatorFromWhitelist = "rm_delegator_from_whitelist"

var _ sdk.Msg = &MsgRmDelegatorFromWhitelist{}

func NewMsgRmDelegatorFromWhitelist(creator string, delAddress string) *MsgRmDelegatorFromWhitelist {
	return &MsgRmDelegatorFromWhitelist{
		Creator:    creator,
		DelAddress: delAddress,
	}
}

func (msg *MsgRmDelegatorFromWhitelist) Route() string {
	return RouterKey
}

func (msg *MsgRmDelegatorFromWhitelist) Type() string {
	return TypeMsgRmDelegatorFromWhitelist
}

func (msg *MsgRmDelegatorFromWhitelist) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRmDelegatorFromWhitelist) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRmDelegatorFromWhitelist) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
