package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/relayers module sentinel errors
var (
	ErrEmptyRelayerAddr        = sdkerrors.Register(ModuleName, 1, "relayer address is empty")
	ErrRelayerAlreadySet       = sdkerrors.Register(ModuleName, 2, "relayer already set")
	ErrThresholdNotFound       = sdkerrors.Register(ModuleName, 3, "threshold not found")
	ErrProposerNotRelayer      = sdkerrors.Register(ModuleName, 4, "proposer is not a relayer")
	ErrInvalidProposalContent  = sdkerrors.Register(ModuleName, 5, "invalid proposal content")
	ErrInvalidProposalType     = sdkerrors.Register(ModuleName, 6, "invalid proposal type")
	ErrAlreadyVoted = sdkerrors.Register(ModuleName, 7, "already voted")
	ErrNoProposalHandlerExists = sdkerrors.Register(ModuleName, 8, "no handler exists for proposal type")


)
