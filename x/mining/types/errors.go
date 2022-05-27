package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/mining module sentinel errors
var (
	ErrStakePoolAlreadyExist                   = sdkerrors.Register(ModuleName, 1101, "stake pool already exist")
	ErrStakePoolNotExist                       = sdkerrors.Register(ModuleName, 1102, "stake pool not exist")
	ErrRewardPoolNumberReachLimit              = sdkerrors.Register(ModuleName, 1103, "reward pool number reach limit")
	ErrStakeItemNotExist                       = sdkerrors.Register(ModuleName, 1104, "stake item not exist")
	ErrUserStakeRecordNotExist                 = sdkerrors.Register(ModuleName, 1105, "user stake record not exist")
	ErrStakeTokenStillLocked                   = sdkerrors.Register(ModuleName, 1106, "stake token still locked")
	ErrUserNotAdminOrMiningProvider            = sdkerrors.Register(ModuleName, 1107, "user is not admin or mining provider")
	ErrUserNotMiningProvider                   = sdkerrors.Register(ModuleName, 1108, "user is not mining provider")
	ErrTotalRewardAmountLessThanLimit          = sdkerrors.Register(ModuleName, 1109, "total reward amount is less than limit")
	ErrWithdrawAmountMoreThanStakeRecord       = sdkerrors.Register(ModuleName, 1110, "withdraw amount is more than stake record")
	ErrRewardTokenNotSupport                   = sdkerrors.Register(ModuleName, 1111, "reward token not support")
	ErrRewardPoolNotExist                      = sdkerrors.Register(ModuleName, 1112, "reward pool not exist")
	ErrRewardPoolLeftRewardAmountNotZero       = sdkerrors.Register(ModuleName, 1113, "reward pool left reward amount not zero")
	ErrMiningProviderNotExist                  = sdkerrors.Register(ModuleName, 1114, "mining provider not exist")
	ErrStakeItemNumberReachLimit               = sdkerrors.Register(ModuleName, 1115, "stake item number reach limit")
	ErrStartTimestampAndRewardPerSecondNotZero = sdkerrors.Register(ModuleName, 1116, "start timestamp and reward per second not zero")
	ErrRewardTokenDenomDuplicate               = sdkerrors.Register(ModuleName, 1117, "reward token denom duplicate")
	ErrEmergencySwitchOpen                     = sdkerrors.Register(ModuleName, 1119, "emergency switch is open")
	ErrEmergencySwitchClose                    = sdkerrors.Register(ModuleName, 1120, "emergency switch is close")
	ErrLockTimeOverRewardTime                  = sdkerrors.Register(ModuleName, 1121, "lock time over reward time")
	ErrUpdateStakeItemPermissionDeny           = sdkerrors.Register(ModuleName, 1122, "update stake item permisson deny")
)
