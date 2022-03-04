package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetRelayGasPrice = "set_relay_gas_price"

var _ sdk.Msg = &MsgSetRelayGasPrice{}

func NewMsgSetRelayGasPrice(creator string, denom string, gasPrice string) *MsgSetRelayGasPrice {
	return &MsgSetRelayGasPrice{
		Creator:  creator,
		Denom:    denom,
		GasPrice: gasPrice,
	}
}

func (msg *MsgSetRelayGasPrice) Route() string {
	return RouterKey
}

func (msg *MsgSetRelayGasPrice) Type() string {
	return TypeMsgSetRelayGasPrice
}

func (msg *MsgSetRelayGasPrice) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetRelayGasPrice) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetRelayGasPrice) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
