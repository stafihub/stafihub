package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgSetChainEra{}
	_ sdk.Msg = &MsgActiveReport{}
)

func NewMsgSetChainEra(creator string, denom string, era uint32) *MsgSetChainEra {
	return &MsgSetChainEra{
		Creator: creator,
		Denom: denom,
		Era: era,
	}
}

func (msg *MsgSetChainEra) Route() string {
	return RouterKey
}

func (msg *MsgSetChainEra) Type() string {
	return "SetChainEra"
}

func (msg *MsgSetChainEra) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetChainEra) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetChainEra) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func NewMsgActiveReport(creator string, denom string, shotId string, staked string, unstaked string) *MsgActiveReport {
	return &MsgActiveReport{
		Creator: creator,
		Denom: denom,
		ShotId: shotId,
		Staked: staked,
		Unstaked: unstaked,
	}
}

func (msg *MsgActiveReport) Route() string {
	return RouterKey
}

func (msg *MsgActiveReport) Type() string {
	return "ActiveReport"
}

func (msg *MsgActiveReport) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgActiveReport) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgActiveReport) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
