package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetRelayFeeReceiver = "set_relay_fee_receiver"

var _ sdk.Msg = &MsgSetRelayFeeReceiver{}

func NewMsgSetRelayFeeReceiver(creator string, denom, receiver string) *MsgSetRelayFeeReceiver {
	return &MsgSetRelayFeeReceiver{
		Creator:  creator,
		Denom:    denom,
		Receiver: receiver,
	}
}

func (msg *MsgSetRelayFeeReceiver) Route() string {
	return RouterKey
}

func (msg *MsgSetRelayFeeReceiver) Type() string {
	return TypeMsgSetRelayFeeReceiver
}

func (msg *MsgSetRelayFeeReceiver) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetRelayFeeReceiver) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetRelayFeeReceiver) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
