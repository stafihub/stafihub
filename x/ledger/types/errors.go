package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/ledger module sentinel errors
var (
	ErrPoolAlreadyAdded = sdkerrors.Register(ModuleName, 1, "pool already added error")
	ErrPoolNotFound = sdkerrors.Register(ModuleName, 2, "pool not found error")
	ErrPoolNotBonded = sdkerrors.Register(ModuleName, 3, "pool not bonded error")
	ErrRepeatInitBond = sdkerrors.Register(ModuleName, 4, "repeat init bond error")
	ErrActiveAlreadySet = sdkerrors.Register(ModuleName, 5, "active already set error")
	ErrBondPipelineNotFound = sdkerrors.Register(ModuleName, 6, "bond pipeline not found error")
)
