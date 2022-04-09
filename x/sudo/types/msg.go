package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = &MsgUpdateAdmin{}
)

func NewMsgUpdateAdmin(creator sdk.AccAddress, address sdk.AccAddress) *MsgUpdateAdmin {
	return &MsgUpdateAdmin{
		Creator: creator.String(),
		Address: address.String(),
	}
}

func (msg *MsgUpdateAdmin) Route() string {
	return RouterKey
}

func (msg *MsgUpdateAdmin) Type() string {
	return "UpdateAdmin"
}

func (msg *MsgUpdateAdmin) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic("invalid creator address")
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateAdmin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateAdmin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return err
	}
	_, err = sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return err
	}
	return nil
}
