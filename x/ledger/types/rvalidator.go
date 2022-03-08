package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	utils "github.com/stafihub/stafihub/utils"
)

var (
	_ sdk.Msg = &MsgOnboard{}
	_ sdk.Msg = &MsgSetRValidatorIndicator{}
)

func NewMsgOnboard(creator sdk.AccAddress, denom string, address, operatorAddress string, locked sdk.Coin) *MsgOnboard {
	return &MsgOnboard{
		Creator:         creator.String(),
		Denom:           denom,
		Address:         address,
		OperatorAddress: operatorAddress,
		Locked:          locked,
	}
}

func (msg *MsgOnboard) Route() string {
	return RouterKey
}

func (msg *MsgOnboard) Type() string {
	return "Onboard"
}

func (msg *MsgOnboard) GetSigners() []sdk.AccAddress {
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{creator}
}

func (msg *MsgOnboard) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgOnboard) ValidateBasic() error {
	if msg.Creator == "" || msg.Address == "" || msg.OperatorAddress == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator (%s) or Address (%s) or OperatorAddress(%s)", msg.Creator, msg.Address, msg.OperatorAddress)
	}
	return nil
}

func NewMsgSetRValidatorIndicator(creator sdk.AccAddress, denom string, commission utils.Dec, uptime uint32, votingPower int64, locked sdk.Coin) *MsgSetRValidatorIndicator {
	return &MsgSetRValidatorIndicator{
		Creator:     creator.String(),
		Denom:       denom,
		Commission:  commission,
		Uptime:      uptime,
		VotingPower: votingPower,
		Locked:      locked,
	}
}

func (msg *MsgSetRValidatorIndicator) Route() string {
	return RouterKey
}

func (msg *MsgSetRValidatorIndicator) Type() string {
	return "SetRValidatorIndicator"
}

func (msg *MsgSetRValidatorIndicator) GetSigners() []sdk.AccAddress {
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetRValidatorIndicator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetRValidatorIndicator) ValidateBasic() error {
	if msg.Creator == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}

	if msg.Commission.LT(utils.ZeroDec()) {
		return fmt.Errorf("commission %s less than zeroDec", msg.Commission.String())
	}

	if msg.Commission.GTE(utils.OneDec()) {
		return fmt.Errorf("commission %s bigger than oneDec", msg.Commission.String())
	}

	if msg.VotingPower < 0 {
		return fmt.Errorf("votingPower %d is less than 0", msg.VotingPower)
	}

	return nil
}
