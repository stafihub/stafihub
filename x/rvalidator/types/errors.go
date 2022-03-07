package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/rvalidator module sentinel errors
var (
	ErrRValidatorAlreadyOnboard  = sdkerrors.Register(ModuleName, 1, "rvalidator already onboard error")
	ErrRValidatorIndicatorNotSet = sdkerrors.Register(ModuleName, 2, "rvalidator indicator not set error")
	ErrLockedNotEnough           = sdkerrors.Register(ModuleName, 3, "locked not enough error")
)
