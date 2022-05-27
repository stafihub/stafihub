package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgStake = "stake"

var _ sdk.Msg = &MsgStake{}

func NewMsgStake(creator string, stakePoolIndex uint32, stakeAmount sdk.Int, stakeItemIndex uint32) *MsgStake {
	return &MsgStake{
		Creator:        creator,
		StakePoolIndex: stakePoolIndex,
		StakeAmount:    stakeAmount,
		StakeItemIndex: stakeItemIndex,
	}
}

func (msg *MsgStake) Route() string {
	return RouterKey
}

func (msg *MsgStake) Type() string {
	return TypeMsgStake
}

func (msg *MsgStake) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgStake) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgStake) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if !msg.StakeAmount.IsPositive() {
		return fmt.Errorf("stake token amount not positive")
	}

	return nil
}
