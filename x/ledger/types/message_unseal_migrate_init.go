package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUnsealMigrateInit = "unseal_migrate_init"

var _ sdk.Msg = &MsgUnsealMigrateInit{}

func NewMsgUnsealMigrateInit(creator string) *MsgUnsealMigrateInit {
	return &MsgUnsealMigrateInit{
		Creator: creator,
	}
}

func (msg *MsgUnsealMigrateInit) Route() string {
	return RouterKey
}

func (msg *MsgUnsealMigrateInit) Type() string {
	return TypeMsgUnsealMigrateInit
}

func (msg *MsgUnsealMigrateInit) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUnsealMigrateInit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUnsealMigrateInit) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
