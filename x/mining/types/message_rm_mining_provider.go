package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRmMiningProvider = "rm_mining_provider"

var _ sdk.Msg = &MsgRmMiningProvider{}

func NewMsgRmMiningProvider(creator string, userAddress string) *MsgRmMiningProvider {
	return &MsgRmMiningProvider{
		Creator:     creator,
		UserAddress: userAddress,
	}
}

func (msg *MsgRmMiningProvider) Route() string {
	return RouterKey
}

func (msg *MsgRmMiningProvider) Type() string {
	return TypeMsgRmMiningProvider
}

func (msg *MsgRmMiningProvider) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRmMiningProvider) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRmMiningProvider) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.UserAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user address (%s)", err)
	}
	return nil
}
