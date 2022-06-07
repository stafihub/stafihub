package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetCycleSeconds = "set_cycle_seconds"

var _ sdk.Msg = &MsgSetCycleSeconds{}

func NewMsgSetCycleSeconds(creator string, denom string, seconds uint64) *MsgSetCycleSeconds {
	return &MsgSetCycleSeconds{
		Creator: creator,
		Denom:   denom,
		Seconds: seconds,
	}
}

func (msg *MsgSetCycleSeconds) Route() string {
	return RouterKey
}

func (msg *MsgSetCycleSeconds) Type() string {
	return TypeMsgSetCycleSeconds
}

func (msg *MsgSetCycleSeconds) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetCycleSeconds) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetCycleSeconds) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
