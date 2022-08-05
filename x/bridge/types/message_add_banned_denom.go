package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddBannedDenom = "add_banned_denom"

var _ sdk.Msg = &MsgAddBannedDenom{}

func NewMsgAddBannedDenom(creator string, chainId uint32, denom string) *MsgAddBannedDenom {
	return &MsgAddBannedDenom{
		Creator: creator,
		ChainId: chainId,
		Denom:   denom,
	}
}

func (msg *MsgAddBannedDenom) Route() string {
	return RouterKey
}

func (msg *MsgAddBannedDenom) Type() string {
	return TypeMsgAddBannedDenom
}

func (msg *MsgAddBannedDenom) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddBannedDenom) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddBannedDenom) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
