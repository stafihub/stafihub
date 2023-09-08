package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgOpenIcaChannel = "open_ica_channel"

var _ sdk.Msg = &MsgOpenIcaChannel{}

func NewMsgOpenIcaChannel(creator string, poolAddress string, accountType AccountType) *MsgOpenIcaChannel {
	return &MsgOpenIcaChannel{
		Creator:     creator,
		PoolAddress: poolAddress,
		AccountType: accountType,
	}
}

func (msg *MsgOpenIcaChannel) Route() string {
	return RouterKey
}

func (msg *MsgOpenIcaChannel) Type() string {
	return TypeMsgOpenIcaChannel
}

func (msg *MsgOpenIcaChannel) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgOpenIcaChannel) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgOpenIcaChannel) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
