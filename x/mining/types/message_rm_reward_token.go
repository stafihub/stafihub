package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRmRewardToken = "rm_reward_token"

var _ sdk.Msg = &MsgRmRewardToken{}

func NewMsgRmRewardToken(creator string, denom string) *MsgRmRewardToken {
	return &MsgRmRewardToken{
		Creator: creator,
		Denom:   denom,
	}
}

func (msg *MsgRmRewardToken) Route() string {
	return RouterKey
}

func (msg *MsgRmRewardToken) Type() string {
	return TypeMsgRmRewardToken
}

func (msg *MsgRmRewardToken) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRmRewardToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRmRewardToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if err := sdk.ValidateDenom(msg.Denom); err != nil {
		return err
	}
	return nil
}
