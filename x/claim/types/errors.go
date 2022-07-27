package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/claim module sentinel errors
var (
	ErrMerkleRootFormat = sdkerrors.Register(ModuleName, 1100, "merkle root format err")
)
