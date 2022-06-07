package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetShuffleSeconds = "set_shuffle_seconds"

var _ sdk.Msg = &MsgSetShuffleSeconds{}

func NewMsgSetShuffleSeconds(creator string, denom string, seconds uint64) *MsgSetShuffleSeconds {
	return &MsgSetShuffleSeconds{
		Creator: creator,
		Denom:   denom,
		Seconds: seconds,
	}
}

func (msg *MsgSetShuffleSeconds) Route() string {
	return RouterKey
}

func (msg *MsgSetShuffleSeconds) Type() string {
	return TypeMsgSetShuffleSeconds
}

func (msg *MsgSetShuffleSeconds) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetShuffleSeconds) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetShuffleSeconds) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
