package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddRewarder = "add_rewarder"

var _ sdk.Msg = &MsgAddRewarder{}

func NewMsgAddRewarder(creator string, userAddress string) *MsgAddRewarder {
	return &MsgAddRewarder{
		Creator:     creator,
		UserAddress: userAddress,
	}
}

func (msg *MsgAddRewarder) Route() string {
	return RouterKey
}

func (msg *MsgAddRewarder) Type() string {
	return TypeMsgAddRewarder
}

func (msg *MsgAddRewarder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddRewarder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddRewarder) ValidateBasic() error {
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
