package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetPoolStatus = "set_pool_status"

var _ sdk.Msg = &MsgSetPoolStatus{}

func NewMsgSetPoolStatus(creator string, denom string, pool string, status PoolStatus) *MsgSetPoolStatus {
	return &MsgSetPoolStatus{
		Creator: creator,
		Denom:   denom,
		Pool:    pool,
		Status:  status,
	}
}

func (msg *MsgSetPoolStatus) Route() string {
	return RouterKey
}

func (msg *MsgSetPoolStatus) Type() string {
	return TypeMsgSetPoolStatus
}

func (msg *MsgSetPoolStatus) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetPoolStatus) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetPoolStatus) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
