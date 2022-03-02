package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRmChainId = "rm_chain_id"

var _ sdk.Msg = &MsgRmChainId{}

func NewMsgRmChainId(creator string, chainId uint32) *MsgRmChainId {
	return &MsgRmChainId{
		Creator: creator,
		ChainId: chainId,
	}
}

func (msg *MsgRmChainId) Route() string {
	return RouterKey
}

func (msg *MsgRmChainId) Type() string {
	return TypeMsgRmChainId
}

func (msg *MsgRmChainId) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRmChainId) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRmChainId) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
