package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetMerkleRoot = "set_merkle_root"

var _ sdk.Msg = &MsgSetMerkleRoot{}

func NewMsgSetMerkleRoot(creator string, round uint64, merkleRoot string) *MsgSetMerkleRoot {
	return &MsgSetMerkleRoot{
		Creator:    creator,
		Round:      round,
		MerkleRoot: merkleRoot,
	}
}

func (msg *MsgSetMerkleRoot) Route() string {
	return RouterKey
}

func (msg *MsgSetMerkleRoot) Type() string {
	return TypeMsgSetMerkleRoot
}

func (msg *MsgSetMerkleRoot) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetMerkleRoot) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetMerkleRoot) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
