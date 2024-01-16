package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// staking message types
//
// #nosec G101
const (
	TypeMsgRedeemTokensForShares       = "redeem_tokens_for_shares"
	TypeMsgTransferTokenizeShareRecord = "transfer_tokenize_share_record"
	TypeMsgTokenizeShares              = "tokenize_shares"
)

var (
	_ sdk.Msg = &MsgRedeemTokensForShares{}
	_ sdk.Msg = &MsgTransferTokenizeShareRecord{}
	_ sdk.Msg = &MsgTokenizeShares{}
)

// NewMsgRedeemTokensForShares creates a new MsgRedeemTokensForShares instance.
//
//nolint:interfacer
func NewMsgRedeemTokensForShares(delAddr sdk.AccAddress, amount sdk.Coin) *MsgRedeemTokensForShares {
	return &MsgRedeemTokensForShares{
		DelegatorAddress: delAddr.String(),
		Amount:           amount,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgRedeemTokensForShares) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgRedeemTokensForShares) Type() string { return TypeMsgRedeemTokensForShares }

// GetSigners implements the sdk.Msg interface.
func (msg MsgRedeemTokensForShares) GetSigners() []sdk.AccAddress {
	delegator, err := sdk.AccAddressFromBech32(msg.DelegatorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{delegator}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgRedeemTokensForShares) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgRedeemTokensForShares) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.DelegatorAddress); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid delegator address: %s", err)
	}

	if !msg.Amount.IsValid() || !msg.Amount.Amount.IsPositive() {
		return errorsmod.Wrap(
			sdkerrors.ErrInvalidRequest,
			"invalid shares amount",
		)
	}

	return nil
}

// NewMsgTransferTokenizeShareRecord creates a new MsgTransferTokenizeShareRecord instance.
//
//nolint:interfacer
func NewMsgTransferTokenizeShareRecord(recordId uint64, sender, newOwner sdk.AccAddress) *MsgTransferTokenizeShareRecord {
	return &MsgTransferTokenizeShareRecord{
		TokenizeShareRecordId: recordId,
		Sender:                sender.String(),
		NewOwner:              newOwner.String(),
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgTransferTokenizeShareRecord) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgTransferTokenizeShareRecord) Type() string { return TypeMsgTransferTokenizeShareRecord }

// GetSigners implements the sdk.Msg interface.
func (msg MsgTransferTokenizeShareRecord) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgTransferTokenizeShareRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgTransferTokenizeShareRecord) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid sender address: %s", err)
	}
	if _, err := sdk.AccAddressFromBech32(msg.NewOwner); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid new owner address: %s", err)
	}

	return nil
}

// Route implements the sdk.Msg interface.
func (msg MsgTokenizeShares) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgTokenizeShares) Type() string { return TypeMsgTokenizeShares }

// GetSigners implements the sdk.Msg interface.
func (msg MsgTokenizeShares) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.DelegatorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgTokenizeShares) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgTokenizeShares) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.DelegatorAddress); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid sender address: %s", err)
	}
	if _, err := sdk.AccAddressFromBech32(msg.TokenizedShareOwner); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid new owner address: %s", err)
	}

	return nil
}
