package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/relayers module sentinel errors
var (
	ErrEmptyRelayerAddr = sdkerrors.Register(ModuleName, 1, "relayer address is empty")
	ErrCreatorNotAdmin = sdkerrors.Register(ModuleName, 2, "creator is not admin")
	ErrRelayerAlreadySet = sdkerrors.Register(ModuleName, 3, "relayer already set")
)
