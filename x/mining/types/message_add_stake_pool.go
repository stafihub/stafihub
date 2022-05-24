package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddStakePool = "add_stake_pool"

var _ sdk.Msg = &MsgAddStakePool{}

func NewMsgAddStakePool(creator string, stakeTokenDenom, rewardTokenDenom string, totalRewardAmount, rewardPerSecond sdk.Int, startTimestamp uint64) *MsgAddStakePool {
	return &MsgAddStakePool{
		Creator:           creator,
		StakeTokenDenom:   stakeTokenDenom,
		RewardTokenDenom:  rewardTokenDenom,
		TotalRewardAmount: totalRewardAmount,
		RewardPerSecond:   rewardPerSecond,
		StartTimestamp:    startTimestamp,
	}
}

func (msg *MsgAddStakePool) Route() string {
	return RouterKey
}

func (msg *MsgAddStakePool) Type() string {
	return TypeMsgAddStakePool
}

func (msg *MsgAddStakePool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddStakePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddStakePool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	err = sdk.ValidateDenom(msg.StakeTokenDenom)
	if err != nil {
		return err
	}
	err = sdk.ValidateDenom(msg.RewardTokenDenom)
	if err != nil {
		return err
	}
	if msg.TotalRewardAmount.IsNegative() {
		return fmt.Errorf("minTotalRewardAmount is negative")
	}
	if msg.RewardPerSecond.IsNegative() {
		return fmt.Errorf("RewardPerSecond is negative")
	}
	return nil
}
