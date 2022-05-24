package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddRewardPool = "add_reward_pool"

var _ sdk.Msg = &MsgAddRewardPool{}

func NewMsgAddRewardPool(creator string, stakePoolIndex uint32, rewardTokenDenom string, totalRewardAmount sdk.Int, rewardPerSecond sdk.Int, startTimestamp uint64) *MsgAddRewardPool {
	return &MsgAddRewardPool{
		Creator:           creator,
		StakePoolIndex:    stakePoolIndex,
		RewardTokenDenom:  rewardTokenDenom,
		TotalRewardAmount: totalRewardAmount,
		RewardPerSecond:   rewardPerSecond,
		StartTimestamp:    startTimestamp,
	}
}

func (msg *MsgAddRewardPool) Route() string {
	return RouterKey
}

func (msg *MsgAddRewardPool) Type() string {
	return TypeMsgAddRewardPool
}

func (msg *MsgAddRewardPool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddRewardPool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddRewardPool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	err = sdk.ValidateDenom(msg.RewardTokenDenom)
	if err != nil {
		return err
	}
	if !msg.TotalRewardAmount.IsPositive() {
		return fmt.Errorf("totalRewardAmount is not positive")
	}

	if !msg.RewardPerSecond.IsPositive() {
		return fmt.Errorf("rewardPerSecond is not positive")
	}

	return nil
}
