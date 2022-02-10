package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetRParams = "set_r_params"

var _ sdk.Msg = &MsgSetRParams{}

func NewMsgSetRParams(creator string, denom string, chainId string, nativeDenom string, gasPrice string, eraSeconds string, validators []string) *MsgSetRParams {
	return &MsgSetRParams{
		Creator:     creator,
		Denom:       denom,
		ChainId:     chainId,
		NativeDenom: nativeDenom,
		GasPrice:    gasPrice,
		EraSeconds:  eraSeconds,
		Validators:  validators,
	}
}

func (msg *MsgSetRParams) Route() string {
	return RouterKey
}

func (msg *MsgSetRParams) Type() string {
	return TypeMsgSetRParams
}

func (msg *MsgSetRParams) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetRParams) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetRParams) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
