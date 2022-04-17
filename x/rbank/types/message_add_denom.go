package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	xBankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

const TypeMsgAddDenom = "add_denom"

var _ sdk.Msg = &MsgAddDenom{}

func NewMsgAddDenom(creator string, accAddressPrefix, valAddressPrefix string, metadata xBankTypes.Metadata) *MsgAddDenom {
	return &MsgAddDenom{
		Creator:          creator,
		AccAddressPrefix: accAddressPrefix,
		ValAddressPrefix: valAddressPrefix,
		Metadata:         metadata,
	}
}

func (msg *MsgAddDenom) Route() string {
	return RouterKey
}

func (msg *MsgAddDenom) Type() string {
	return TypeMsgAddDenom
}

func (msg *MsgAddDenom) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddDenom) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddDenom) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
