package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stafihub/stafihub/utils"
)

const TypeMsgAddStakeItem = "add_stake_item"

var _ sdk.Msg = &MsgAddStakeItem{}

func NewMsgAddStakeItem(creator string, index uint32, lockSecond uint64, powerRewardRate utils.Dec) *MsgAddStakeItem {
	return &MsgAddStakeItem{
		Creator:         creator,
		Index:           index,
		LockSecond:      lockSecond,
		PowerRewardRate: powerRewardRate,
	}
}

func (msg *MsgAddStakeItem) Route() string {
	return RouterKey
}

func (msg *MsgAddStakeItem) Type() string {
	return TypeMsgAddStakeItem
}

func (msg *MsgAddStakeItem) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddStakeItem) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddStakeItem) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
