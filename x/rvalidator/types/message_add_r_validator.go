package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddRValidator = "add_r_validator"

var _ sdk.Msg = &MsgAddRValidator{}

func NewMsgAddRValidator(creator string, denom string, addressList []string) *MsgAddRValidator {
	return &MsgAddRValidator{
		Creator:     creator,
		Denom:       denom,
		AddressList: addressList,
	}
}

func (msg *MsgAddRValidator) Route() string {
	return RouterKey
}

func (msg *MsgAddRValidator) Type() string {
	return TypeMsgAddRValidator
}

func (msg *MsgAddRValidator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddRValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddRValidator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if err := sdk.ValidateDenom(msg.Denom); err != nil {
		return err
	}

	if len(msg.AddressList) == 0 {
		return fmt.Errorf("address list is empty")
	}
	return nil
}
