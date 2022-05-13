package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/mining module sentinel errors
var (
	ErrStakePoolAlreadyExist = sdkerrors.Register(ModuleName, 1101, "stake pool already exist")
	ErrStakePoolNotExist     = sdkerrors.Register(ModuleName, 1102, "stake pool not exist")
)
