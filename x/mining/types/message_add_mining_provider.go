package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddMiningProvider = "add_mining_provider"

var _ sdk.Msg = &MsgAddMiningProvider{}

func NewMsgAddMiningProvider(creator string, userAddress string) *MsgAddMiningProvider {
	return &MsgAddMiningProvider{
		Creator:     creator,
		UserAddress: userAddress,
	}
}

func (msg *MsgAddMiningProvider) Route() string {
	return RouterKey
}

func (msg *MsgAddMiningProvider) Type() string {
	return TypeMsgAddMiningProvider
}

func (msg *MsgAddMiningProvider) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddMiningProvider) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddMiningProvider) ValidateBasic() error {
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
