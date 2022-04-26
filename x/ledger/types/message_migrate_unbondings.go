package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMigrateUnbondings = "migrate_unbondings"

var _ sdk.Msg = &MsgMigrateUnbondings{}

func NewMsgMigrateUnbondings(creator, denom, pool string, era uint32, unbondings []*Unbonding) *MsgMigrateUnbondings {
	return &MsgMigrateUnbondings{
		Creator:    creator,
		Denom:      denom,
		Pool:       pool,
		Era:        era,
		Unbondings: unbondings,
	}
}

func (msg *MsgMigrateUnbondings) Route() string {
	return RouterKey
}

func (msg *MsgMigrateUnbondings) Type() string {
	return TypeMsgMigrateUnbondings
}

func (msg *MsgMigrateUnbondings) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMigrateUnbondings) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMigrateUnbondings) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
