package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddChainId = "add_chain_id"

var _ sdk.Msg = &MsgAddChainId{}

func NewMsgAddChainId(creator string, chainId uint32) *MsgAddChainId {
	return &MsgAddChainId{
		Creator: creator,
		ChainId: chainId,
	}
}

func (msg *MsgAddChainId) Route() string {
	return RouterKey
}

func (msg *MsgAddChainId) Type() string {
	return TypeMsgAddChainId
}

func (msg *MsgAddChainId) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddChainId) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddChainId) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
