package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetRelayFee = "set_relay_fee"

var _ sdk.Msg = &MsgSetRelayFee{}

func NewMsgSetRelayFee(creator string, chainId uint32, value sdk.Coin) *MsgSetRelayFee {
	return &MsgSetRelayFee{
		Creator: creator,
		ChainId: chainId,
		Value:   value,
	}
}

func (msg *MsgSetRelayFee) Route() string {
	return RouterKey
}

func (msg *MsgSetRelayFee) Type() string {
	return TypeMsgSetRelayFee
}

func (msg *MsgSetRelayFee) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetRelayFee) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetRelayFee) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
