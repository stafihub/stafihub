package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddPoolCreator = "add_pool_creator"

var _ sdk.Msg = &MsgAddPoolCreator{}

func NewMsgAddPoolCreator(creator string, userAddress string) *MsgAddPoolCreator {
	return &MsgAddPoolCreator{
		Creator:     creator,
		UserAddress: userAddress,
	}
}

func (msg *MsgAddPoolCreator) Route() string {
	return RouterKey
}

func (msg *MsgAddPoolCreator) Type() string {
	return TypeMsgAddPoolCreator
}

func (msg *MsgAddPoolCreator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddPoolCreator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddPoolCreator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
