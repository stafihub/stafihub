package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgWithdrawRewardToken = "withdraw_reward_token"

var _ sdk.Msg = &MsgWithdrawRewardToken{}

func NewMsgWithdrawRewardToken(creator string, stakePoolIndex, rewardPoolIndex uint32, withdrawAmount sdk.Int) *MsgWithdrawRewardToken {
	return &MsgWithdrawRewardToken{
		Creator:         creator,
		StakePoolIndex:  stakePoolIndex,
		RewardPoolIndex: rewardPoolIndex,
		WithdrawAmount:  withdrawAmount,
	}
}

func (msg *MsgWithdrawRewardToken) Route() string {
	return RouterKey
}

func (msg *MsgWithdrawRewardToken) Type() string {
	return TypeMsgWithdrawRewardToken
}

func (msg *MsgWithdrawRewardToken) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgWithdrawRewardToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgWithdrawRewardToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if !msg.WithdrawAmount.IsPositive() {
		return fmt.Errorf("withdrawAmount not positive")
	}
	return nil
}
