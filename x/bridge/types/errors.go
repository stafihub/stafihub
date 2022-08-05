package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/bridge module sentinel errors
var (
	ErrSample                   = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrResourceIdNotFound       = sdkerrors.Register(ModuleName, 1101, "resourceId not found")
	ErrBalanceNotEnough         = sdkerrors.Register(ModuleName, 1102, "balance not enough")
	ErrChainIdNotSupport        = sdkerrors.Register(ModuleName, 1103, "chainId not support")
	ErrUnKnownResourceIdType    = sdkerrors.Register(ModuleName, 1104, "unknown resource id type")
	ErrAlreadyVoted             = sdkerrors.Register(ModuleName, 1105, "voter already voted")
	ErrAlreadyExecuted          = sdkerrors.Register(ModuleName, 1106, "voter already executed")
	ErrThresholdNotSet          = sdkerrors.Register(ModuleName, 1107, "threshold not set")
	ErrRelayFeeReceiverNotSet   = sdkerrors.Register(ModuleName, 1108, "relay fee receiver not set")
	ErrRelayerNotExist          = sdkerrors.Register(ModuleName, 1109, "relay not exist")
	ErrReceiverFormatNotRight   = sdkerrors.Register(ModuleName, 1110, "receiver format not right")
	ErrResourceIdFormatNotRight = sdkerrors.Register(ModuleName, 1111, "resource id format not right")
	ErrDepositAmountZero        = sdkerrors.Register(ModuleName, 1112, "deposit amount zero")
	ErrVoteProposalAmountZero   = sdkerrors.Register(ModuleName, 1113, "vote proposal amount zero")
	ErrBannedDenom              = sdkerrors.Register(ModuleName, 1114, "banned denom")
	ErrDenomTypeUnmatch         = sdkerrors.Register(ModuleName, 1115, "denom type unmatch")
)
