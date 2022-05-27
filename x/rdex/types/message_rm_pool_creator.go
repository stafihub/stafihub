package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRmPoolCreator = "rm_pool_creator"

var _ sdk.Msg = &MsgRmPoolCreator{}

func NewMsgRmPoolCreator(creator string, userAddress string) *MsgRmPoolCreator {
	return &MsgRmPoolCreator{
		Creator:     creator,
		UserAddress: userAddress,
	}
}

func (msg *MsgRmPoolCreator) Route() string {
	return RouterKey
}

func (msg *MsgRmPoolCreator) Type() string {
	return TypeMsgRmPoolCreator
}

func (msg *MsgRmPoolCreator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRmPoolCreator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRmPoolCreator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
