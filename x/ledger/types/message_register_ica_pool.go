package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRegisterIcaPool = "register_ica_pool"

var _ sdk.Msg = &MsgRegisterIcaPool{}

func NewMsgRegisterIcaPool(creator string, denom string, connectionId string) *MsgRegisterIcaPool {
	return &MsgRegisterIcaPool{
		Creator:      creator,
		Denom:        denom,
		ConnectionId: connectionId,
	}
}

func (msg *MsgRegisterIcaPool) Route() string {
	return RouterKey
}

func (msg *MsgRegisterIcaPool) Type() string {
	return TypeMsgRegisterIcaPool
}

func (msg *MsgRegisterIcaPool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterIcaPool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegisterIcaPool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
