package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/claim module sentinel errors
var (
	ErrMerkleRootFormatNotMatch = sdkerrors.Register(ModuleName, 1100, "merkle root format not match")
	ErrAlreadyClaimed           = sdkerrors.Register(ModuleName, 1101, "already claimed err")
	ErrNodeHashFormatNotMatch   = sdkerrors.Register(ModuleName, 1102, "node hash format not match")
	ErrAccountFormatNotMatch    = sdkerrors.Register(ModuleName, 1103, "account format not match")
	ErrAmountFormatNotMatch     = sdkerrors.Register(ModuleName, 1104, "amount format not match")
	ErrMerkleRootNotExist       = sdkerrors.Register(ModuleName, 1105, "merkle root not exist")
	ErrMerkleProofNotMatch      = sdkerrors.Register(ModuleName, 1106, "merkle proof not match")
	ErrClaimSwitchClosed        = sdkerrors.Register(ModuleName, 1107, "claim switch closed")
)
