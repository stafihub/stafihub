package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgWithdraw = "withdraw"

var _ sdk.Msg = &MsgWithdraw{}

func NewMsgWithdraw(creator string, stakePoolIndex uint32, stakeToken sdk.Coin, stakeRecordIndex uint32) *MsgWithdraw {
	return &MsgWithdraw{
		Creator:          creator,
		StakePoolIndex:   stakePoolIndex,
		StakeToken:       stakeToken,
		StakeRecordIndex: stakeRecordIndex,
	}
}

func (msg *MsgWithdraw) Route() string {
	return RouterKey
}

func (msg *MsgWithdraw) Type() string {
	return TypeMsgWithdraw
}

func (msg *MsgWithdraw) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgWithdraw) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgWithdraw) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	err = msg.StakeToken.Validate()
	if err != nil {
		return err
	}
	return nil
}
