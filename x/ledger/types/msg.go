package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgSetEraUnbondLimit{}
	_ sdk.Msg = &MsgSetPoolDetail{}
	_ sdk.Msg = &MsgSetLeastBond{}
	_ sdk.Msg = &MsgClearCurrentEraSnapShots{}
	_ sdk.Msg = &MsgSetStakingRewardCommission{}
	_ sdk.Msg = &MsgSetProtocolFeeReceiver{}
	_ sdk.Msg = &MsgSetUnbondCommission{}
	_ sdk.Msg = &MsgLiquidityUnbond{}
	_ sdk.Msg = &MsgSetUnbondRelayFee{}
	_ sdk.Msg = &MsgSubmitSignature{}
	_ sdk.Msg = &MsgSetRParams{}
)

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

func NewMsgSetLeastBond(creator sdk.AccAddress, denom, leastBond string) *MsgSetLeastBond {
	return &MsgSetLeastBond{
		Creator:   creator.String(),
		Denom:     denom,
		LeastBond: leastBond,
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

func NewMsgSetStakingRewardCommission(creator sdk.AccAddress, denom string, commission sdk.Dec) *MsgSetStakingRewardCommission {
	return &MsgSetStakingRewardCommission{
		Creator:    creator.String(),
		Denom:      denom,
		Commission: commission,
	}
}

func (msg *MsgSetStakingRewardCommission) Route() string {
	return RouterKey
}

func (msg *MsgSetStakingRewardCommission) Type() string {
	return "SetStakingRewardCommission"
}

func (msg *MsgSetStakingRewardCommission) GetSigners() []sdk.AccAddress {
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetStakingRewardCommission) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetStakingRewardCommission) ValidateBasic() error {
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

func NewMsgSetProtocolFeeReceiver(creator sdk.AccAddress, receiver sdk.AccAddress) *MsgSetProtocolFeeReceiver {
	return &MsgSetProtocolFeeReceiver{
		Creator:  creator.String(),
		Receiver: receiver.String(),
	}
}

func (msg *MsgSetProtocolFeeReceiver) Route() string {
	return RouterKey
}

func (msg *MsgSetProtocolFeeReceiver) Type() string {
	return "SetProtocolFeeReceiver"
}

func (msg *MsgSetProtocolFeeReceiver) GetSigners() []sdk.AccAddress {
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetProtocolFeeReceiver) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetProtocolFeeReceiver) ValidateBasic() error {
	if msg.Creator == "" || msg.Receiver == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator (%s) or receiver (%s)", msg.Creator, msg.Receiver)
	}
	return nil
}

func NewMsgSetUnbondCommission(creator sdk.AccAddress, denom string, commission sdk.Dec) *MsgSetUnbondCommission {
	return &MsgSetUnbondCommission{
		Creator:    creator.String(),
		Denom:      denom,
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

func NewMsgSetUnbondRelayFee(creator sdk.AccAddress, denom string, value sdk.Coin) *MsgSetUnbondRelayFee {
	return &MsgSetUnbondRelayFee{
		Creator: creator.String(),
		Denom:   denom,
		Value:   value,
	}
}

func (msg *MsgSetUnbondRelayFee) Route() string {
	return RouterKey
}

func (msg *MsgSetUnbondRelayFee) Type() string {
	return "SetUnbondRelayFee"
}

func (msg *MsgSetUnbondRelayFee) GetSigners() []sdk.AccAddress {
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetUnbondRelayFee) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetUnbondRelayFee) ValidateBasic() error {
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

const TypeMsgSetRParams = "set_r_params"

func NewMsgSetRParams(creator string, denom string, gasPrice string, eraSeconds string, offset string, bondingDuration uint32, leastBond string, validators []string) *MsgSetRParams {
	return &MsgSetRParams{
		Creator:         creator,
		Denom:           denom,
		GasPrice:        gasPrice,
		EraSeconds:      eraSeconds,
		Offset:          offset,
		BondingDuration: bondingDuration,
		LeastBond:       leastBond,
		Validators:      validators,
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
