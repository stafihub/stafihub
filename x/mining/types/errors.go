package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/mining module sentinel errors
var (
	ErrStakePoolAlreadyExist             = sdkerrors.Register(ModuleName, 1101, "stake pool already exist")
	ErrStakePoolNotExist                 = sdkerrors.Register(ModuleName, 1102, "stake pool not exist")
	ErrRewardPoolNumberReachLimit        = sdkerrors.Register(ModuleName, 1103, "reward pool number reach limit")
	ErrStakeItemNotExist                 = sdkerrors.Register(ModuleName, 1104, "stake item not exist")
	ErrUserStakeRecordNotExist           = sdkerrors.Register(ModuleName, 1105, "user stake record not exist")
	ErrStakeTokenStillLocked             = sdkerrors.Register(ModuleName, 1106, "stake token still locked")
	ErrUserNotAdminOrRewarder            = sdkerrors.Register(ModuleName, 1107, "user is not admin or rewarder")
	ErrUserNotRewarder                   = sdkerrors.Register(ModuleName, 1108, "user is not rewarder")
	ErrTotalRewardAmountLessThanLimit    = sdkerrors.Register(ModuleName, 1109, "total reward amount is less than limit")
	ErrWithdrawAmountMoreThanStakeRecord = sdkerrors.Register(ModuleName, 1110, "withdraw amount is more than stake record")
)
