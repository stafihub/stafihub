package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddDelegatorToWhitelist = "add_delegator_to_whitelist"

var _ sdk.Msg = &MsgAddDelegatorToWhitelist{}

func NewMsgAddDelegatorToWhitelist(creator string, delAddress string) *MsgAddDelegatorToWhitelist {
	return &MsgAddDelegatorToWhitelist{
		Creator:    creator,
		DelAddress: delAddress,
	}
}

func (msg *MsgAddDelegatorToWhitelist) Route() string {
	return RouterKey
}

func (msg *MsgAddDelegatorToWhitelist) Type() string {
	return TypeMsgAddDelegatorToWhitelist
}

func (msg *MsgAddDelegatorToWhitelist) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddDelegatorToWhitelist) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddDelegatorToWhitelist) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
