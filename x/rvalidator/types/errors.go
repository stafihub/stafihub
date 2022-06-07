package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/rvalidator module sentinel errors
var (
	ErrRValidatorAlreadyExist = sdkerrors.Register(ModuleName, 1101, "rValidator already exist")
	ErrRValidatorNotExist     = sdkerrors.Register(ModuleName, 1102, "rValidator not exist")
	ErrCycleBehindLatestCycle = sdkerrors.Register(ModuleName, 1103, "cycle behind latest voted cycle")
)
