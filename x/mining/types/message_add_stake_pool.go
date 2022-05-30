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
	if len(msg.StakeItemInfoList) == 0 {
		return fmt.Errorf("stake item list empty")
	}
	denomMap := make(map[string]bool)
	for _, rewardPool := range msg.RewardPoolInfoList {
		err = sdk.ValidateDenom(rewardPool.RewardTokenDenom)
		if err != nil {
			return err
		}
		if denomMap[rewardPool.RewardTokenDenom] {
			return ErrRewardTokenDenomDuplicate
		}
		denomMap[rewardPool.RewardTokenDenom] = true

		if !rewardPool.TotalRewardAmount.IsPositive() {
			return fmt.Errorf("minTotalRewardAmount is not positive")
		}
		if !rewardPool.RewardPerSecond.IsPositive() {
			return fmt.Errorf("RewardPerSecond is not positive")
		}
	}

	for _, stakeItem := range msg.StakeItemInfoList {
		if !stakeItem.PowerRewardRate.IsPositive() {
			return fmt.Errorf("PowerRewardRate is not positive")
		}
	}

	return nil
}
