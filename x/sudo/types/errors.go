package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/sudo module sentinel errors
var (
	ErrCreatorNotAdmin        = sdkerrors.Register(ModuleName, 1, "creator is not admin error")
	ErrLastAdminEqualNewAdmin = sdkerrors.Register(ModuleName, 2, "last admin euqal new admin error")
)
