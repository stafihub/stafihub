package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/rbank module sentinel errors
var (
	ErrDenomAlreadyExist     = sdkerrors.Register(ModuleName, 1, "denom already exist error")
	ErrAccAddrPrefixNotExist = sdkerrors.Register(ModuleName, 2, "acc addr prefix not exist error")
	ErrDenomNotMatched       = sdkerrors.Register(ModuleName, 3, "denom not matched error")
	ErrAddrPrefixNotMatched  = sdkerrors.Register(ModuleName, 4, "addr prefix not matched error")
	ErrValAddrPrefixNotExist = sdkerrors.Register(ModuleName, 5, "val addr prefix not exist error")
)
