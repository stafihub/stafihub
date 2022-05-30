package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddRewardToken = "add_reward_token"

var _ sdk.Msg = &MsgAddRewardToken{}

func NewMsgAddRewardToken(creator string, denom string, minTotalRewardAmount sdk.Int) *MsgAddRewardToken {
	return &MsgAddRewardToken{
		Creator:              creator,
		Denom:                denom,
		MinTotalRewardAmount: minTotalRewardAmount,
	}
}

func (msg *MsgAddRewardToken) Route() string {
	return RouterKey
}

func (msg *MsgAddRewardToken) Type() string {
	return TypeMsgAddRewardToken
}

func (msg *MsgAddRewardToken) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddRewardToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddRewardToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if err := sdk.ValidateDenom(msg.Denom); err != nil {
		return err
	}

	if !msg.MinTotalRewardAmount.IsPositive() {
		return fmt.Errorf("minTotalRewardAmount is not positive")
	}
	return nil
}
