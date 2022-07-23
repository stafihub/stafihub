package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetWithdrawalAddr = "set_withdrawal_addr"

var _ sdk.Msg = &MsgSetWithdrawalAddr{}

func NewMsgSetWithdrawalAddr(creator string, delegationAddr string) *MsgSetWithdrawalAddr {
	return &MsgSetWithdrawalAddr{
		Creator:        creator,
		DelegationAddr: delegationAddr,
	}
}

func (msg *MsgSetWithdrawalAddr) Route() string {
	return RouterKey
}

func (msg *MsgSetWithdrawalAddr) Type() string {
	return TypeMsgSetWithdrawalAddr
}

func (msg *MsgSetWithdrawalAddr) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetWithdrawalAddr) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetWithdrawalAddr) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
