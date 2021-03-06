package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetEraSeconds = "set_era_seconds"

var _ sdk.Msg = &MsgSetEraSeconds{}

func NewMsgSetEraSeconds(creator string, denom string, eraSeconds uint32, bondingDuration uint32, offset int32) *MsgSetEraSeconds {
	return &MsgSetEraSeconds{
		Creator:         creator,
		Denom:           denom,
		EraSeconds:      eraSeconds,
		BondingDuration: bondingDuration,
		Offset:          offset,
	}
}

func (msg *MsgSetEraSeconds) Route() string {
	return RouterKey
}

func (msg *MsgSetEraSeconds) Type() string {
	return TypeMsgSetEraSeconds
}

func (msg *MsgSetEraSeconds) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetEraSeconds) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetEraSeconds) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.EraSeconds == 0 {
		return fmt.Errorf("eraSeconds cannot be zero")
	}
	return nil
}
