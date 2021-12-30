package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/rvote module sentinel errors
var (
	ErrInvalidProposalContent  = sdkerrors.Register(ModuleName, 1, "invalid proposal content error")
	ErrInvalidProposalType     = sdkerrors.Register(ModuleName, 2, "invalid proposal type error")
	ErrNoProposalHandlerExists = sdkerrors.Register(ModuleName, 3, "no handler exists for proposal type error")
)
