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
	_ sdk.Msg = &MsgSetUnbondCommission{}
	_ sdk.Msg = &MsgLiquidityUnbond{}
	_ sdk.Msg = &MsgSetUnbondFee{}
	_ sdk.Msg = &MsgSubmitSignature{}
)

func NewMsgAddNewPool(creator sdk.AccAddress, denom string, addr string) *MsgAddNewPool {
	return &MsgAddNewPool{
		Creator: creator.String(),
		Denom:   denom,
		Addr:    addr,
	}
}

func (msg *MsgAddNewPool) Route() string {
	return RouterKey
}

func (msg *MsgAddNewPool) Type() string {
	return "AddNewPool"
}

func (msg *MsgAddNewPool) GetSigners() []sdk.AccAddress {
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddNewPool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddNewPool) ValidateBasic() error {
	if msg.Creator == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}

	return nil
}

func NewMsgRemovePool(creator sdk.AccAddress, denom string, addr string) *MsgRemovePool {
	return &MsgRemovePool{
		Creator: creator.String(),
		Denom:   denom,
		Addr:    addr,
	}
}

func (msg *MsgRemovePool) Route() string {
	return RouterKey
}

func (msg *MsgRemovePool) Type() string {
	return "RemovePool"
}

func (msg *MsgRemovePool) GetSigners() []sdk.AccAddress {
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemovePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemovePool) ValidateBasic() error {
	if msg.Creator == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}

	return nil
}

func NewMsgSetEraUnbondLimit(creator sdk.AccAddress, denom string, limit uint32) *MsgSetEraUnbondLimit {
	return &MsgSetEraUnbondLimit{
		Creator: creator.String(),
		Denom:   denom,
		Limit:   limit,
	}
}

func (msg *MsgSetEraUnbondLimit) Route() string {
	return RouterKey
}

func (msg *MsgSetEraUnbondLimit) Type() string {
	return "SetEraUnbondLimit"
}

func (msg *MsgSetEraUnbondLimit) GetSigners() []sdk.AccAddress {
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetEraUnbondLimit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetEraUnbondLimit) ValidateBasic() error {
	if msg.Creator == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}

	return nil
}

func NewMsgSetInitBond(creator sdk.AccAddress, pool string, coin sdk.Coin, receiver sdk.AccAddress) *MsgSetInitBond {
	return &MsgSetInitBond{
		Creator:  creator.String(),
		Pool:     pool,
		Coin:     coin,
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
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetInitBond) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetInitBond) ValidateBasic() error {
	if msg.Creator == "" || msg.Receiver == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator (%s) or receiver (%s)", msg.Creator, msg.Receiver)
	}
	return nil
}

func NewMsgSetChainBondingDuration(creator sdk.AccAddress, denom string, era uint32) *MsgSetChainBondingDuration {
	return &MsgSetChainBondingDuration{
		Creator: creator.String(),
		Denom:   denom,
		Era:     era,
	}
}

func (msg *MsgSetChainBondingDuration) Route() string {
	return RouterKey
}

func (msg *MsgSetChainBondingDuration) Type() string {
	return "SetChainBondingDuration"
}

func (msg *MsgSetChainBondingDuration) GetSigners() []sdk.AccAddress {
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetChainBondingDuration) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetChainBondingDuration) ValidateBasic() error {
	if msg.Creator == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}

	return nil
}

func NewMsgSetPoolDetail(creator sdk.AccAddress, denom string, pool string, subAccounts []string, threshold uint32) *MsgSetPoolDetail {
	return &MsgSetPoolDetail{
		Creator:     creator.String(),
		Denom:       denom,
		Pool:        pool,
		SubAccounts: subAccounts,
		Threshold:   threshold,
	}
}

func (msg *MsgSetPoolDetail) Route() string {
	return RouterKey
}

func (msg *MsgSetPoolDetail) Type() string {
	return "SetPoolDetail"
}

func (msg *MsgSetPoolDetail) GetSigners() []sdk.AccAddress {
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetPoolDetail) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetPoolDetail) ValidateBasic() error {
	if msg.Creator == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address")
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
		Denom:   denom,
		Amount:  amount,
	}
}

func (msg *MsgSetLeastBond) Route() string {
	return RouterKey
}

func (msg *MsgSetLeastBond) Type() string {
	return "SetLeastBond"
}

func (msg *MsgSetLeastBond) GetSigners() []sdk.AccAddress {
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetLeastBond) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetLeastBond) ValidateBasic() error {
	if msg.Creator == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}

	return nil
}

func NewMsgClearCurrentEraSnapShots(creator sdk.AccAddress, denom string) *MsgClearCurrentEraSnapShots {
	return &MsgClearCurrentEraSnapShots{
		Creator: creator.String(),
		Denom:   denom,
	}
}

func (msg *MsgClearCurrentEraSnapShots) Route() string {
	return RouterKey
}

func (msg *MsgClearCurrentEraSnapShots) Type() string {
	return "ClearCurrentEraSnapShots"
}

func (msg *MsgClearCurrentEraSnapShots) GetSigners() []sdk.AccAddress {
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{creator}
}

func (msg *MsgClearCurrentEraSnapShots) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgClearCurrentEraSnapShots) ValidateBasic() error {
	if msg.Creator == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}

	return nil
}

func NewMsgSetCommission(creator sdk.AccAddress, commission sdk.Dec) *MsgSetCommission {
	return &MsgSetCommission{
		Creator:    creator.String(),
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
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetCommission) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetCommission) ValidateBasic() error {
	if msg.Creator == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}

	if msg.Commission.LT(sdk.ZeroDec()) {
		return fmt.Errorf("commission %s less than zeroDec", msg.Commission.String())
	}

	if msg.Commission.GTE(sdk.OneDec()) {
		return fmt.Errorf("commission %s bigger than oneDec", msg.Commission.String())
	}
	return nil
}

func NewMsgSetReceiver(creator sdk.AccAddress, receiver sdk.AccAddress) *MsgSetReceiver {
	return &MsgSetReceiver{
		Creator:  creator.String(),
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
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetReceiver) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetReceiver) ValidateBasic() error {
	if msg.Creator == "" || msg.Receiver == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator (%s) or receiver (%s)", msg.Creator, msg.Receiver)
	}
	return nil
}

func NewMsgSetUnbondCommission(creator sdk.AccAddress, commission sdk.Dec) *MsgSetUnbondCommission {
	return &MsgSetUnbondCommission{
		Creator:    creator.String(),
		Commission: commission,
	}
}

func (msg *MsgSetUnbondCommission) Route() string {
	return RouterKey
}

func (msg *MsgSetUnbondCommission) Type() string {
	return "SetUnbondCommission"
}

func (msg *MsgSetUnbondCommission) GetSigners() []sdk.AccAddress {
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetUnbondCommission) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetUnbondCommission) ValidateBasic() error {
	if msg.Creator == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}

	if msg.Commission.LT(sdk.ZeroDec()) {
		return fmt.Errorf("commission %s less than zeroDec", msg.Commission.String())
	}

	if msg.Commission.GTE(sdk.OneDec()) {
		return fmt.Errorf("commission %s bigger than oneDec", msg.Commission.String())
	}

	return nil
}

func NewMsgLiquidityUnbond(creator sdk.AccAddress, pool string, value sdk.Coin, recipient string) *MsgLiquidityUnbond {
	return &MsgLiquidityUnbond{
		Creator:   creator.String(),
		Pool:      pool,
		Value:     value,
		Recipient: recipient,
	}
}

func (msg *MsgLiquidityUnbond) Route() string {
	return RouterKey
}

func (msg *MsgLiquidityUnbond) Type() string {
	return "LiquidityUnbond"
}

func (msg *MsgLiquidityUnbond) GetSigners() []sdk.AccAddress {
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{creator}
}

func (msg *MsgLiquidityUnbond) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgLiquidityUnbond) ValidateBasic() error {
	if msg.Creator == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}

	if msg.Value.Amount.LTE(sdk.ZeroInt()) {
		return fmt.Errorf("unbond value %s less than zeroInt", msg.Value.String())
	}
	return nil
}

func NewMsgSetUnbondFee(creator sdk.AccAddress, denom string, value sdk.Coin) *MsgSetUnbondFee {
	return &MsgSetUnbondFee{
		Creator: creator.String(),
		Denom:   denom,
		Value:   value,
	}
}

func (msg *MsgSetUnbondFee) Route() string {
	return RouterKey
}

func (msg *MsgSetUnbondFee) Type() string {
	return "SetUnbondFee"
}

func (msg *MsgSetUnbondFee) GetSigners() []sdk.AccAddress {
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetUnbondFee) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetUnbondFee) ValidateBasic() error {
	if msg.Creator == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}

	return nil
}

func NewMsgSubmitSignature(creator string, denom string, era uint32, pool string, txType OriginalTxType, propId string, signature string) *MsgSubmitSignature {
	return &MsgSubmitSignature{
		Creator:   creator,
		Denom:     denom,
		Era:       era,
		Pool:      pool,
		TxType:    txType,
		PropId:    propId,
		Signature: signature,
	}
}

func (msg *MsgSubmitSignature) Route() string {
	return RouterKey
}

func (msg *MsgSubmitSignature) Type() string {
	return "SubmitSignature"
}

func (msg *MsgSubmitSignature) GetSigners() []sdk.AccAddress {
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{creator}
}

func (msg *MsgSubmitSignature) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSubmitSignature) ValidateBasic() error {
	if msg.Creator == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}
	return nil
}
