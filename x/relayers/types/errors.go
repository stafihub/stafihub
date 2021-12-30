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
	ErrAlreadyVoted            = sdkerrors.Register(ModuleName, 5, "already voted error")
)
