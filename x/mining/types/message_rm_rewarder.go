package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRmRewarder = "rm_rewarder"

var _ sdk.Msg = &MsgRmRewarder{}

func NewMsgRmRewarder(creator string, userAddress string) *MsgRmRewarder {
	return &MsgRmRewarder{
		Creator:     creator,
		UserAddress: userAddress,
	}
}

func (msg *MsgRmRewarder) Route() string {
	return RouterKey
}

func (msg *MsgRmRewarder) Type() string {
	return TypeMsgRmRewarder
}

func (msg *MsgRmRewarder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRmRewarder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRmRewarder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.UserAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user address (%s)", err)
	}
	return nil
}
