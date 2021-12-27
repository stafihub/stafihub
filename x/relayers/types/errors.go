package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/relayers module sentinel errors
var (
	ErrRelayerAlreadySet       = sdkerrors.Register(ModuleName, 1, "relayer already set error")
	ErrRelayerNotFound         = sdkerrors.Register(ModuleName, 2, "relayer not found error")
	ErrThresholdNotFound       = sdkerrors.Register(ModuleName, 3, "threshold not found error")
	ErrProposerNotRelayer      = sdkerrors.Register(ModuleName, 4, "proposer is not a relayer error")
	ErrInvalidProposalContent  = sdkerrors.Register(ModuleName, 5, "invalid proposal content error")
	ErrInvalidProposalType     = sdkerrors.Register(ModuleName, 6, "invalid proposal type error")
	ErrAlreadyVoted            = sdkerrors.Register(ModuleName, 7, "already voted error")
	ErrNoProposalHandlerExists = sdkerrors.Register(ModuleName, 8, "no handler exists for proposal type error")
)
