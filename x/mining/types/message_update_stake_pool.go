package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateStakePool = "update_stake_pool"

var _ sdk.Msg = &MsgUpdateStakePool{}

func NewMsgUpdateStakePool(creator string, stakeTokenDenom string, maxRewardPools uint32, mintTotalRewardAmount sdk.Int) *MsgUpdateStakePool {
	return &MsgUpdateStakePool{
		Creator:              creator,
		StakeTokenDenom:      stakeTokenDenom,
		MaxRewardPools:       maxRewardPools,
		MinTotalRewardAmount: mintTotalRewardAmount,
	}
}

func (msg *MsgUpdateStakePool) Route() string {
	return RouterKey
}

func (msg *MsgUpdateStakePool) Type() string {
	return TypeMsgUpdateStakePool
}

func (msg *MsgUpdateStakePool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateStakePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateStakePool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	err = sdk.ValidateDenom(msg.StakeTokenDenom)
	if err != nil {
		return err
	}
	if msg.MinTotalRewardAmount.IsNegative() {
		return fmt.Errorf("minTotalRewardAmount is negative")
	}
	return nil
}
