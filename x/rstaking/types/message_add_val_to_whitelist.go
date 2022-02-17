package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddValToWhitelist = "add_val_to_whitelist"

var _ sdk.Msg = &MsgAddValToWhitelist{}

func NewMsgAddValToWhitelist(creator string, valAddress string) *MsgAddValToWhitelist {
	return &MsgAddValToWhitelist{
		Creator:    creator,
		ValAddress: valAddress,
	}
}

func (msg *MsgAddValToWhitelist) Route() string {
	return RouterKey
}

func (msg *MsgAddValToWhitelist) Type() string {
	return TypeMsgAddValToWhitelist
}

func (msg *MsgAddValToWhitelist) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddValToWhitelist) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddValToWhitelist) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
