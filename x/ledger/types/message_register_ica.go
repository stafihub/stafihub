package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRegisterIca = "register_ica"

var _ sdk.Msg = &MsgRegisterIca{}

func NewMsgRegisterIca(creator string, owner string, connectionId string) *MsgRegisterIca {
	return &MsgRegisterIca{
		Creator:      creator,
		Owner:        owner,
		ConnectionId: connectionId,
	}
}

func (msg *MsgRegisterIca) Route() string {
	return RouterKey
}

func (msg *MsgRegisterIca) Type() string {
	return TypeMsgRegisterIca
}

func (msg *MsgRegisterIca) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterIca) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegisterIca) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
