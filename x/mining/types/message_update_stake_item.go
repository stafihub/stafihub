package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stafihub/stafihub/utils"
)

const TypeMsgUpdateStakeItem = "update_stake_item"

var _ sdk.Msg = &MsgUpdateStakeItem{}

func NewMsgUpdateStakeItem(creator string, index uint32, lockSecond uint64, powerRewardRate utils.Dec, enable bool) *MsgUpdateStakeItem {
	return &MsgUpdateStakeItem{
		Creator:         creator,
		Index:           index,
		LockSecond:      lockSecond,
		PowerRewardRate: powerRewardRate,
		Enable:          enable,
	}
}

func (msg *MsgUpdateStakeItem) Route() string {
	return RouterKey
}

func (msg *MsgUpdateStakeItem) Type() string {
	return TypeMsgUpdateStakeItem
}

func (msg *MsgUpdateStakeItem) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateStakeItem) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateStakeItem) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if !msg.PowerRewardRate.IsPositive() {
		return fmt.Errorf("powerRewardRate is not positive")
	}
	return nil
}
