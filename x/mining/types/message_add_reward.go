package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddReward = "add_reward"

var _ sdk.Msg = &MsgAddReward{}

func NewMsgAddReward(creator string, stakePoolIndex, rewardPoolIndex uint32, addAmount sdk.Int, startTimestamp uint64, rewardPerSecond sdk.Int) *MsgAddReward {
	return &MsgAddReward{
		Creator:         creator,
		StakePoolIndex:  stakePoolIndex,
		RewardPoolIndex: rewardPoolIndex,
		AddAmount:       addAmount,
		StartTimestamp:  startTimestamp,
		RewardPerSecond: rewardPerSecond,
	}
}

func (msg *MsgAddReward) Route() string {
	return RouterKey
}

func (msg *MsgAddReward) Type() string {
	return TypeMsgAddReward
}

func (msg *MsgAddReward) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddReward) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddReward) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if !msg.AddAmount.IsPositive() {
		return fmt.Errorf("addAmount is not positive")
	}

	if msg.RewardPerSecond.IsNegative() {
		return fmt.Errorf("rewardPerSecond is negative")
	}
	return nil
}
