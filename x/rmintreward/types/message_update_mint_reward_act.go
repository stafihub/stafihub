package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateMintRewardAct = "update_mint_reward_act"

var _ sdk.Msg = &MsgUpdateMintRewardAct{}

func NewMsgUpdateMintRewardAct(creator string, denom string, cycle uint64, act *MintRewardActPost) *MsgUpdateMintRewardAct {
	return &MsgUpdateMintRewardAct{
		Creator: creator,
		Denom:   denom,
		Cycle:   cycle,
		Act:     act,
	}
}

func (msg *MsgUpdateMintRewardAct) Route() string {
	return RouterKey
}

func (msg *MsgUpdateMintRewardAct) Type() string {
	return TypeMsgUpdateMintRewardAct
}

func (msg *MsgUpdateMintRewardAct) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateMintRewardAct) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateMintRewardAct) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
