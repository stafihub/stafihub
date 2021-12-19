package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgAddNewPool{}
	_ sdk.Msg = &MsgRemovePool{}
	_ sdk.Msg = &MsgSetEraUnbondLimit{}
	_ sdk.Msg = &MsgSetInitBond{}
	_ sdk.Msg = &MsgSetChainBondingDuration{}
	_ sdk.Msg = &MsgSetPoolDetail{}
	_ sdk.Msg = &MsgSetLeastBond{}
	_ sdk.Msg = &MsgClearCurrentEraSnapShots{}
	_ sdk.Msg = &MsgSetCommission{}
	_ sdk.Msg = &MsgSetReceiver{}
)

func NewMsgAddNewPool(creator sdk.AccAddress, denom string, addr string) *MsgAddNewPool {
  return &MsgAddNewPool{
		Creator: creator.String(),
    Denom: denom,
    Addr: addr,
	}
}

func (msg *MsgAddNewPool) Route() string {
  return RouterKey
}

func (msg *MsgAddNewPool) Type() string {
  return "AddNewPool"
}

func (msg *MsgAddNewPool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddNewPool) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgAddNewPool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func NewMsgRemovePool(creator sdk.AccAddress, denom string, addr string) *MsgRemovePool {
	return &MsgRemovePool{
		Creator: creator.String(),
		Denom: denom,
		Addr: addr,
	}
}

func (msg *MsgRemovePool) Route() string {
	return RouterKey
}

func (msg *MsgRemovePool) Type() string {
	return "RemovePool"
}

func (msg *MsgRemovePool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemovePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemovePool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func NewMsgSetEraUnbondLimit(creator sdk.AccAddress, denom string, limit uint32) *MsgSetEraUnbondLimit {
	return &MsgSetEraUnbondLimit{
		Creator: creator.String(),
		Denom: denom,
		Limit: limit,
	}
}

func (msg *MsgSetEraUnbondLimit) Route() string {
	return RouterKey
}

func (msg *MsgSetEraUnbondLimit) Type() string {
	return "SetEraUnbondLimit"
}

func (msg *MsgSetEraUnbondLimit) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetEraUnbondLimit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetEraUnbondLimit) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func NewMsgSetInitBond(creator sdk.AccAddress, denom string, pool string, amount sdk.Int, receiver sdk.AccAddress) *MsgSetInitBond {
	return &MsgSetInitBond{
		Creator: creator.String(),
		Denom: denom,
		Pool: pool,
		Amount: amount,
		Receiver: receiver.String(),
	}
}

func (msg *MsgSetInitBond) Route() string {
	return RouterKey
}

func (msg *MsgSetInitBond) Type() string {
	return "SetInitBond"
}

func (msg *MsgSetInitBond) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetInitBond) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetInitBond) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address (%s)", err)
	}
	return nil
}

func NewMsgSetChainBondingDuration(creator sdk.AccAddress, denom string, era uint32) *MsgSetChainBondingDuration {
	return &MsgSetChainBondingDuration{
		Creator: creator.String(),
		Denom: denom,
		Era: era,
	}
}

func (msg *MsgSetChainBondingDuration) Route() string {
	return RouterKey
}

func (msg *MsgSetChainBondingDuration) Type() string {
	return "SetChainBondingDuration"
}

func (msg *MsgSetChainBondingDuration) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetChainBondingDuration) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetChainBondingDuration) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func NewMsgSetPoolDetail(creator sdk.AccAddress, denom string, pool string, subAccounts []string, threshold uint32) *MsgSetPoolDetail {
	return &MsgSetPoolDetail{
		Creator: creator.String(),
		Denom: denom,
		Pool: pool,
		SubAccounts: subAccounts,
		Threshold: threshold,
	}
}

func (msg *MsgSetPoolDetail) Route() string {
	return RouterKey
}

func (msg *MsgSetPoolDetail) Type() string {
	return "SetPoolDetail"
}

func (msg *MsgSetPoolDetail) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetPoolDetail) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetPoolDetail) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if len(msg.SubAccounts) != 0 {
		accLen := len(msg.SubAccounts[0])
		for _, acc := range msg.SubAccounts {
			if len(acc) != accLen {
				return fmt.Errorf("subAccounts not same size")
			}
		}
	}
	return nil
}

func NewMsgSetLeastBond(creator sdk.AccAddress, denom string, amount sdk.Int) *MsgSetLeastBond {
	return &MsgSetLeastBond{
		Creator: creator.String(),
		Denom: denom,
		Amount: amount,
	}
}

func (msg *MsgSetLeastBond) Route() string {
	return RouterKey
}

func (msg *MsgSetLeastBond) Type() string {
	return "SetLeastBond"
}

func (msg *MsgSetLeastBond) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetLeastBond) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetLeastBond) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func NewMsgClearCurrentEraSnapShots(creator sdk.AccAddress, denom string) *MsgClearCurrentEraSnapShots {
	return &MsgClearCurrentEraSnapShots{
		Creator: creator.String(),
		Denom: denom,
	}
}

func (msg *MsgClearCurrentEraSnapShots) Route() string {
	return RouterKey
}

func (msg *MsgClearCurrentEraSnapShots) Type() string {
	return "ClearCurrentEraSnapShots"
}

func (msg *MsgClearCurrentEraSnapShots) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgClearCurrentEraSnapShots) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgClearCurrentEraSnapShots) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func NewMsgSetCommission(creator sdk.AccAddress, commission sdk.Dec) *MsgSetCommission {
	return &MsgSetCommission{
		Creator: creator.String(),
		Commission: commission,
	}
}

func (msg *MsgSetCommission) Route() string {
	return RouterKey
}

func (msg *MsgSetCommission) Type() string {
	return "SetCommission"
}

func (msg *MsgSetCommission) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetCommission) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetCommission) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Commission.GT(sdk.OneDec()) {
		return fmt.Errorf("rate %s bigger than oneDec", msg.Commission.String())
	}
	return nil
}

func NewMsgSetReceiver(creator sdk.AccAddress, receiver sdk.AccAddress) *MsgSetReceiver {
	return &MsgSetReceiver{
		Creator: creator.String(),
		Receiver: receiver.String(),
	}
}

func (msg *MsgSetReceiver) Route() string {
	return RouterKey
}

func (msg *MsgSetReceiver) Type() string {
	return "SetReceiver"
}

func (msg *MsgSetReceiver) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetReceiver) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetReceiver) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Receiver == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "empty receiver address (%s)", err)
	}
	return nil
}


