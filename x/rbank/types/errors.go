package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/rbank module sentinel errors
var (
	ErrDenomAlreadyExist    = sdkerrors.Register(ModuleName, 1, "denom already exist error")
	ErrAddrPrefixNotExist   = sdkerrors.Register(ModuleName, 2, "addr prefix not exist error")
	ErrDenomNotMatched      = sdkerrors.Register(ModuleName, 3, "denom not matched error")
	ErrAddrPrefixNotMatched = sdkerrors.Register(ModuleName, 4, "addr prefix not matched error")
)
