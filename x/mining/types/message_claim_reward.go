package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgClaimReward = "claim_reward"

var _ sdk.Msg = &MsgClaimReward{}

func NewMsgClaimReward(creator string, stakeTokenDenom string, index uint32) *MsgClaimReward {
	return &MsgClaimReward{
		Creator:          creator,
		StakeTokenDenom:  stakeTokenDenom,
		StakeRecordIndex: index,
	}
}

func (msg *MsgClaimReward) Route() string {
	return RouterKey
}

func (msg *MsgClaimReward) Type() string {
	return TypeMsgClaimReward
}

func (msg *MsgClaimReward) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgClaimReward) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgClaimReward) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	err = sdk.ValidateDenom(msg.StakeTokenDenom)
	if err != nil {
		return err
	}
	return nil
}
