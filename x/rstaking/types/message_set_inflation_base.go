package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetInflationBase = "set_inflation_base"

var _ sdk.Msg = &MsgSetInflationBase{}

func NewMsgSetInflationBase(creator string, inflationBase sdk.Int) *MsgSetInflationBase {
	return &MsgSetInflationBase{
		Creator:       creator,
		InflationBase: inflationBase,
	}
}

func (msg *MsgSetInflationBase) Route() string {
	return RouterKey
}

func (msg *MsgSetInflationBase) Type() string {
	return TypeMsgSetInflationBase
}

func (msg *MsgSetInflationBase) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetInflationBase) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetInflationBase) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
