package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgStake = "stake"

var _ sdk.Msg = &MsgStake{}

func NewMsgStake(creator string, stakeToken sdk.Coin, stakeItemIndex uint32) *MsgStake {
	return &MsgStake{
		Creator:        creator,
		StakeToken:     stakeToken,
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
	err = msg.StakeToken.Validate()
	if err != nil {
		return err
	}

	return nil
}
