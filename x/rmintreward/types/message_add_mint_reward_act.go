package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddMintRewardAct = "add_mint_reward_act"

var _ sdk.Msg = &MsgAddMintRewardAct{}

func NewMsgAddMintRewardAct(creator, denom string, act *MintRewardActPost) *MsgAddMintRewardAct {
	return &MsgAddMintRewardAct{
		Creator: creator,
		Denom:   denom,
		Act:     act,
	}
}

func (msg *MsgAddMintRewardAct) Route() string {
	return RouterKey
}

func (msg *MsgAddMintRewardAct) Type() string {
	return TypeMsgAddMintRewardAct
}

func (msg *MsgAddMintRewardAct) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddMintRewardAct) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddMintRewardAct) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
