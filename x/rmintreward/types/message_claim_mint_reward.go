package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgClaimMintReward = "claim_mint_reward"

var _ sdk.Msg = &MsgClaimMintReward{}

func NewMsgClaimMintReward(creator string, denom string, cycle uint64, mintIndex uint64) *MsgClaimMintReward {
	return &MsgClaimMintReward{
		Creator:   creator,
		Denom:     denom,
		Cycle:     cycle,
		MintIndex: mintIndex,
	}
}

func (msg *MsgClaimMintReward) Route() string {
	return RouterKey
}

func (msg *MsgClaimMintReward) Type() string {
	return TypeMsgClaimMintReward
}

func (msg *MsgClaimMintReward) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgClaimMintReward) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgClaimMintReward) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
