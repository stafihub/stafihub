package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetWithdrawAddr = "set_withdraw_addr"

var _ sdk.Msg = &MsgSetWithdrawAddr{}

func NewMsgSetWithdrawAddr(creator string, delegationAddr string) *MsgSetWithdrawAddr {
	return &MsgSetWithdrawAddr{
		Creator:        creator,
		DelegationAddr: delegationAddr,
	}
}

func (msg *MsgSetWithdrawAddr) Route() string {
	return RouterKey
}

func (msg *MsgSetWithdrawAddr) Type() string {
	return TypeMsgSetWithdrawAddr
}

func (msg *MsgSetWithdrawAddr) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetWithdrawAddr) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetWithdrawAddr) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
