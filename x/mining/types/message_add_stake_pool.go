package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddStakePool = "add_stake_pool"

var _ sdk.Msg = &MsgAddStakePool{}

func NewMsgAddStakePool(creator string, stakeTokenDenom string, rewardPoolInfoList []*CreateRewardPoolInfo, stakeItemList []*CreateStakeItemInfo) *MsgAddStakePool {
	return &MsgAddStakePool{
		Creator:            creator,
		StakeTokenDenom:    stakeTokenDenom,
		RewardPoolInfoList: rewardPoolInfoList,
		StakeItemInfoList:  stakeItemList,
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
	if len(msg.RewardPoolInfoList) == 0 {
		return fmt.Errorf("reward pool list empty")
	}

	for _, rewardPool := range msg.RewardPoolInfoList {
		err = sdk.ValidateDenom(rewardPool.RewardTokenDenom)
		if err != nil {
			return err
		}
		if rewardPool.TotalRewardAmount.IsNegative() {
			return fmt.Errorf("minTotalRewardAmount is negative")
		}
		if rewardPool.RewardPerSecond.IsNegative() {
			return fmt.Errorf("RewardPerSecond is negative")
		}
	}
	return nil
}
