package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/mining module sentinel errors
var (
	ErrStakePoolAlreadyExist      = sdkerrors.Register(ModuleName, 1101, "stake pool already exist")
	ErrStakePoolNotExist          = sdkerrors.Register(ModuleName, 1102, "stake pool not exist")
	ErrRewardPoolNumberReachLimit = sdkerrors.Register(ModuleName, 1103, "reward pool number reach limit")
	ErrStakeItemNotExist          = sdkerrors.Register(ModuleName, 1104, "stake item not exist")
	ErrUserStakeRecordNotExist    = sdkerrors.Register(ModuleName, 1105, "user stake record not exist")
)
