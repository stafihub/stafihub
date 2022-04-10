package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/rstaking module sentinel errors
var (
	ErrValAlreadyInWhitelist = sdkerrors.Register(ModuleName, 1100, "validator already in whitelist error")
	ErrInsufficientFunds     = sdkerrors.Register(ModuleName, 1101, "insufficient funds error")
)
