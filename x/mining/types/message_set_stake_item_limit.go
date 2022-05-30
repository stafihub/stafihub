package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stafihub/stafihub/utils"
)

const TypeMsgSetStakeItemLimit = "set_stake_item_limit"

var _ sdk.Msg = &MsgSetStakeItemLimit{}

func NewMsgSetStakeItemLimit(creator string, maxLockSecond uint64, maxPowerRewardRate utils.Dec) *MsgSetStakeItemLimit {
	return &MsgSetStakeItemLimit{
		Creator:            creator,
		MaxLockSecond:      maxLockSecond,
		MaxPowerRewardRate: maxPowerRewardRate,
	}
}

func (msg *MsgSetStakeItemLimit) Route() string {
	return RouterKey
}

func (msg *MsgSetStakeItemLimit) Type() string {
	return TypeMsgSetStakeItemLimit
}

func (msg *MsgSetStakeItemLimit) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetStakeItemLimit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetStakeItemLimit) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.MaxLockSecond == 0 {
		return fmt.Errorf("maxLockSecond zero")
	}

	if !msg.MaxPowerRewardRate.IsPositive() {
		return fmt.Errorf("maxPowerRewardRate is not positive")
	}

	return nil
}
