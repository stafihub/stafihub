package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddRelayer = "add_relayer"

var _ sdk.Msg = &MsgAddRelayer{}

func NewMsgAddRelayer(creator string, address string) *MsgAddRelayer {
	return &MsgAddRelayer{
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgAddRelayer) Route() string {
	return RouterKey
}

func (msg *MsgAddRelayer) Type() string {
	return TypeMsgAddRelayer
}

func (msg *MsgAddRelayer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddRelayer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddRelayer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
