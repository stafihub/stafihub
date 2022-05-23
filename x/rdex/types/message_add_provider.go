package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddProvider = "add_provider"

var _ sdk.Msg = &MsgAddProvider{}

func NewMsgAddProvider(creator string, userAddress string) *MsgAddProvider {
	return &MsgAddProvider{
		Creator:     creator,
		UserAddress: userAddress,
	}
}

func (msg *MsgAddProvider) Route() string {
	return RouterKey
}

func (msg *MsgAddProvider) Type() string {
	return TypeMsgAddProvider
}

func (msg *MsgAddProvider) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddProvider) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddProvider) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.UserAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid provider address (%s)", err)
	}

	return nil
}
