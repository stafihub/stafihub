package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgEmergencyWithdraw = "emergency_withdraw"

var _ sdk.Msg = &MsgEmergencyWithdraw{}

func NewMsgEmergencyWithdraw(creator string, stakePoolIndex, stakeRecordIndex uint32) *MsgEmergencyWithdraw {
	return &MsgEmergencyWithdraw{
		Creator:          creator,
		StakePoolIndex:   stakePoolIndex,
		StakeRecordIndex: stakeRecordIndex,
	}
}

func (msg *MsgEmergencyWithdraw) Route() string {
	return RouterKey
}

func (msg *MsgEmergencyWithdraw) Type() string {
	return TypeMsgEmergencyWithdraw
}

func (msg *MsgEmergencyWithdraw) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEmergencyWithdraw) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEmergencyWithdraw) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
