package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgInitRValidator = "init_r_validator"

var _ sdk.Msg = &MsgInitRValidator{}

func NewMsgInitRValidator(creator string, denom, poolAddress string, addressList []string) *MsgInitRValidator {
	return &MsgInitRValidator{
		Creator:        creator,
		Denom:          denom,
		PoolAddress:    poolAddress,
		ValAddressList: addressList,
	}
}

func (msg *MsgInitRValidator) Route() string {
	return RouterKey
}

func (msg *MsgInitRValidator) Type() string {
	return TypeMsgInitRValidator
}

func (msg *MsgInitRValidator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgInitRValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInitRValidator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if err := sdk.ValidateDenom(msg.Denom); err != nil {
		return err
	}

	if len(msg.ValAddressList) == 0 {
		return fmt.Errorf("address list is empty")
	}
	return nil
}
