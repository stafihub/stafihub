package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRmBannedDenom = "rm_banned_denom"

var _ sdk.Msg = &MsgRmBannedDenom{}

func NewMsgRmBannedDenom(creator string, chainId uint32, denom string) *MsgRmBannedDenom {
	return &MsgRmBannedDenom{
		Creator: creator,
		ChainId: chainId,
		Denom:   denom,
	}
}

func (msg *MsgRmBannedDenom) Route() string {
	return RouterKey
}

func (msg *MsgRmBannedDenom) Type() string {
	return TypeMsgRmBannedDenom
}

func (msg *MsgRmBannedDenom) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRmBannedDenom) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRmBannedDenom) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
