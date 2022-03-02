package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRmRelayer = "rm_relayer"

var _ sdk.Msg = &MsgRmRelayer{}

func NewMsgRmRelayer(creator string, chainId uint32, address string) *MsgRmRelayer {
	return &MsgRmRelayer{
		Creator: creator,
		ChainId: chainId,
		Address: address,
	}
}

func (msg *MsgRmRelayer) Route() string {
	return RouterKey
}

func (msg *MsgRmRelayer) Type() string {
	return TypeMsgRmRelayer
}

func (msg *MsgRmRelayer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRmRelayer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRmRelayer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
