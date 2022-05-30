package keeper_test

import (
	"context"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/stafihub/stafihub/testutil/keeper"
	"github.com/stafihub/stafihub/testutil/sample"
	"github.com/stafihub/stafihub/utils"
	"github.com/stafihub/stafihub/x/mining/keeper"
	"github.com/stafihub/stafihub/x/mining/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context, sdk.Context) {
	k, ctx := keepertest.MiningKeeper(t)
	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx), ctx
}

func TestAddAndRmRewarderSuccess(t *testing.T) {
	srv, miningKeeper, ctx, sdkCtx := setupMsgServer(t)
	admin := sample.TestAdminAcc
	rewarder := sample.OriginAccAddress()

	// add rewarder fail when user is not admin
	msgAddRewarder := types.MsgAddMiningProvider{
		Creator:     sample.AccAddress(),
		UserAddress: rewarder.String(),
	}
	_, err := srv.AddMiningProvider(ctx, &msgAddRewarder)
	require.Error(t, err)
	// add rewarder
	msgAddRewarder2 := types.MsgAddMiningProvider{
		Creator:     admin.String(),
		UserAddress: rewarder.String(),
	}
	_, err = srv.AddMiningProvider(ctx, &msgAddRewarder2)
	require.NoError(t, err)
	require.True(t, miningKeeper.HasMiningProvider(sdkCtx, rewarder))
	// rm rewarder
	msgRmRewarder := types.MsgRmMiningProvider{
		Creator:     admin.String(),
		UserAddress: rewarder.String(),
	}
	_, err = srv.RmMiningProvider(ctx, &msgRmRewarder)
	require.NoError(t, err)
	require.False(t, miningKeeper.HasMiningProvider(sdkCtx, rewarder))
}

func TestAddAndUpdateStakeItemSuccess(t *testing.T) {
	srv, miningKeeper, ctx, sdkCtx := setupMsgServer(t)
	admin := sample.TestAdminAcc
	stakeTokenDenom := sample.TestDenom
	rewardTokenDenom := sample.TestDenom1

	// add rewarder
	miningProvider := sample.OriginAccAddress()
	willMintCoins := sdk.NewCoins(sdk.NewCoin(rewardTokenDenom, sdk.NewInt(1e4)))
	keepertest.BankKeeper.MintCoins(sdkCtx, types.ModuleName, willMintCoins)
	keepertest.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, miningProvider, willMintCoins)

	msgAddRewarder := types.MsgAddMiningProvider{
		Creator:     admin.String(),
		UserAddress: miningProvider.String(),
	}
	_, err := srv.AddMiningProvider(ctx, &msgAddRewarder)
	require.NoError(t, err)

	// add reward token
	msgAddRewardToken := types.MsgAddRewardToken{
		Creator:              admin.String(),
		Denom:                rewardTokenDenom,
		MinTotalRewardAmount: sdk.NewInt(1e4),
	}
	_, err = srv.AddRewardToken(ctx, &msgAddRewardToken)
	require.NoError(t, err)

	// add stake token
	msgAddStakeToken := types.MsgAddStakeToken{
		Creator: admin.String(),
		Denom:   stakeTokenDenom,
	}
	_, err = srv.AddStakeToken(ctx, &msgAddStakeToken)
	require.NoError(t, err)

	//add stake pool
	lockSecond := uint64(100)
	powerRewardRate := utils.MustNewDecFromStr("1.5")
	msgAddStakePool := types.MsgAddStakePool{
		Creator:         miningProvider.String(),
		StakeTokenDenom: stakeTokenDenom,
		RewardPoolInfoList: []*types.CreateRewardPoolInfo{
			{
				RewardTokenDenom:  rewardTokenDenom,
				TotalRewardAmount: sdk.NewInt(1e4),
				RewardPerSecond:   sdk.NewInt(2),
				StartTimestamp:    4567,
			},
		},
		StakeItemInfoList: []*types.CreateStakeItemInfo{
			{
				LockSecond:      lockSecond,
				PowerRewardRate: powerRewardRate,
			},
		},
	}

	require.True(t, msgAddStakePool.ValidateBasic() == nil)

	_, found := miningKeeper.GetStakePool(sdkCtx, 0)
	require.False(t, found)

	_, err = srv.AddStakePool(ctx, &msgAddStakePool)
	require.NoError(t, err)

	stakePool, found := miningKeeper.GetStakePool(sdkCtx, 0)
	require.True(t, found)
	require.EqualValues(t, stakePool.StakeTokenDenom, msgAddStakePool.StakeTokenDenom)
	require.EqualValues(t, stakePool.TotalStakedAmount, sdk.ZeroInt())
	require.EqualValues(t, stakePool.TotalStakedPower, sdk.ZeroInt())
	require.EqualValues(t, stakePool.RewardPools[0].Index, 0)

	stakeItem, found := miningKeeper.GetStakeItem(sdkCtx, 0, 0)
	require.True(t, found)
	require.EqualValues(t, stakeItem.Index, 0)
	require.EqualValues(t, stakeItem.StakePoolIndex, 0)
	require.EqualValues(t, stakeItem.LockSecond, lockSecond)
	require.EqualValues(t, stakeItem.PowerRewardRate, powerRewardRate)

	// add stakeItem fail when user is not admin and creator
	msgAddStakeItem := types.MsgAddStakeItem{
		Creator:         sample.AccAddress(),
		StakePoolIndex:  0,
		LockSecond:      0,
		PowerRewardRate: utils.MustNewDecFromStr("1.5"),
		Enable:          true,
	}
	_, err = srv.AddStakeItem(ctx, &msgAddStakeItem)
	require.Error(t, err)

	// add stakeItem use admin
	powerRewardRate = utils.MustNewDecFromStr("2.5")
	msgAddStakeItem2 := types.MsgAddStakeItem{
		Creator:         admin.String(),
		StakePoolIndex:  0,
		LockSecond:      0,
		PowerRewardRate: powerRewardRate,
		Enable:          true,
	}
	_, err = srv.AddStakeItem(ctx, &msgAddStakeItem2)
	require.NoError(t, err)
	stakeItem, found = miningKeeper.GetStakeItem(sdkCtx, 0, 1)
	require.True(t, found)
	require.EqualValues(t, stakeItem.Index, 1)
	require.EqualValues(t, stakeItem.Enable, true)
	require.EqualValues(t, stakeItem.LockSecond, 0)
	require.EqualValues(t, stakeItem.PowerRewardRate, powerRewardRate)

	// update Stake item fail whe user is not admin/not provider
	newPowerRewardRate := utils.MustNewDecFromStr("1.5")
	msgUpdateStakeItem := types.MsgUpdateStakeItem{
		Creator:         sample.AccAddress(),
		Index:           0,
		StakePoolIndex:  0,
		LockSecond:      5,
		PowerRewardRate: newPowerRewardRate,
		Enable:          false,
	}
	_, err = srv.UpdateStakeItem(ctx, &msgUpdateStakeItem)
	require.Error(t, err)

	// update Stake item user creator
	msgUpdateStakeItem.Creator = miningProvider.String()
	_, err = srv.UpdateStakeItem(ctx, &msgUpdateStakeItem)
	require.NoError(t, err)
	stakeItem, found = miningKeeper.GetStakeItem(sdkCtx, 0, 0)
	require.True(t, found)
	require.EqualValues(t, stakeItem.Index, 0)
	require.EqualValues(t, stakeItem.Enable, false)
	require.EqualValues(t, stakeItem.LockSecond, 5)
	require.EqualValues(t, stakeItem.PowerRewardRate, newPowerRewardRate)

	nextIndex := miningKeeper.GetStakeItemNextIndex(sdkCtx, 0)
	require.EqualValues(t, nextIndex, 2)

}

func TestAddStakePoolSuccess(t *testing.T) {
	srv, miningKeeper, ctx, sdkCtx := setupMsgServer(t)
	admin := sample.TestAdminAcc
	stakeTokenDenom := sample.TestDenom
	rewardTokenDenom := sample.TestDenom1

	// add rewarder
	miningProvider := sample.OriginAccAddress()
	willMintCoins := sdk.NewCoins(sdk.NewCoin(rewardTokenDenom, sdk.NewInt(1e4)))
	keepertest.BankKeeper.MintCoins(sdkCtx, types.ModuleName, willMintCoins)
	keepertest.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, miningProvider, willMintCoins)

	msgAddRewarder := types.MsgAddMiningProvider{
		Creator:     admin.String(),
		UserAddress: miningProvider.String(),
	}
	_, err := srv.AddMiningProvider(ctx, &msgAddRewarder)
	require.NoError(t, err)

	// add reward token
	msgAddRewardToken := types.MsgAddRewardToken{
		Creator:              admin.String(),
		Denom:                rewardTokenDenom,
		MinTotalRewardAmount: sdk.NewInt(1e4),
	}
	_, err = srv.AddRewardToken(ctx, &msgAddRewardToken)
	require.NoError(t, err)

	// add stake token
	msgAddStakeToken := types.MsgAddStakeToken{
		Creator: admin.String(),
		Denom:   stakeTokenDenom,
	}
	_, err = srv.AddStakeToken(ctx, &msgAddStakeToken)
	require.NoError(t, err)

	//add stake pool
	now := time.Now()
	sdkCtx = sdkCtx.WithBlockTime(now)
	ctx = sdk.WrapSDKContext(sdkCtx)

	lockSecond := uint64(100)
	powerRewardRate := utils.MustNewDecFromStr("1.5")
	msgAddStakePool := types.MsgAddStakePool{
		Creator:         miningProvider.String(),
		StakeTokenDenom: stakeTokenDenom,
		RewardPoolInfoList: []*types.CreateRewardPoolInfo{
			{
				RewardTokenDenom:  rewardTokenDenom,
				TotalRewardAmount: sdk.NewInt(1e4),
				RewardPerSecond:   sdk.NewInt(2),
				StartTimestamp:    4567,
			},
		},
		StakeItemInfoList: []*types.CreateStakeItemInfo{
			{
				LockSecond:      lockSecond,
				PowerRewardRate: powerRewardRate,
			},
		},
	}

	require.True(t, msgAddStakePool.ValidateBasic() == nil)

	_, found := miningKeeper.GetStakePool(sdkCtx, 0)
	require.False(t, found)

	_, err = srv.AddStakePool(ctx, &msgAddStakePool)
	require.NoError(t, err)

	stakePool, found := miningKeeper.GetStakePool(sdkCtx, 0)
	require.True(t, found)
	require.EqualValues(t, stakePool.StakeTokenDenom, msgAddStakePool.StakeTokenDenom)
	require.EqualValues(t, stakePool.TotalStakedAmount, sdk.ZeroInt())
	require.EqualValues(t, stakePool.TotalStakedPower, sdk.ZeroInt())
	require.EqualValues(t, stakePool.RewardPools[0].Index, 0)
	require.EqualValues(t, stakePool.RewardPools[0].LastRewardTimestamp, now.Unix())
	require.EqualValues(t, stakePool.RewardPools[0].LeftRewardAmount, sdk.NewInt(1e4))
	require.EqualValues(t, stakePool.RewardPools[0].RewardPerPower, sdk.ZeroInt())
	require.EqualValues(t, stakePool.RewardPools[0].RewardPerSecond, sdk.NewInt(2))
	require.EqualValues(t, stakePool.RewardPools[0].RewardTokenDenom, rewardTokenDenom)
	require.EqualValues(t, stakePool.RewardPools[0].StartTimestamp, now.Unix())
	require.EqualValues(t, stakePool.RewardPools[0].TotalRewardAmount, sdk.NewInt(1e4))

	stakeItem, found := miningKeeper.GetStakeItem(sdkCtx, 0, 0)
	require.True(t, found)
	require.EqualValues(t, stakeItem.Index, 0)
	require.EqualValues(t, stakeItem.StakePoolIndex, 0)
	require.EqualValues(t, stakeItem.LockSecond, lockSecond)
	require.EqualValues(t, stakeItem.PowerRewardRate, powerRewardRate)
}

func TestAddStakePoolFail(t *testing.T) {
	srv, miningKeeper, ctx, sdkCtx := setupMsgServer(t)
	admin := sample.TestAdminAcc
	stakeTokenDenom := sample.TestDenom
	rewardTokenDenom := sample.TestDenom1

	// add mining provider
	miningProvider := sample.OriginAccAddress()
	willMintCoins := sdk.NewCoins(sdk.NewCoin(rewardTokenDenom, sdk.NewInt(1e4)))
	keepertest.BankKeeper.MintCoins(sdkCtx, types.ModuleName, willMintCoins)
	keepertest.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, miningProvider, willMintCoins)

	msgAddRewarder := types.MsgAddMiningProvider{
		Creator:     admin.String(),
		UserAddress: miningProvider.String(),
	}
	_, err := srv.AddMiningProvider(ctx, &msgAddRewarder)
	require.NoError(t, err)

	// add reward token
	msgAddRewardToken := types.MsgAddRewardToken{
		Creator:              admin.String(),
		Denom:                rewardTokenDenom,
		MinTotalRewardAmount: sdk.NewInt(1e4),
	}
	_, err = srv.AddRewardToken(ctx, &msgAddRewardToken)
	require.NoError(t, err)

	// add stake token
	msgAddStakeToken := types.MsgAddStakeToken{
		Creator: admin.String(),
		Denom:   stakeTokenDenom,
	}
	_, err = srv.AddStakeToken(ctx, &msgAddStakeToken)
	require.NoError(t, err)

	//add stake pool fail if use admin
	lockSecond := uint64(100)
	powerRewardRate := utils.MustNewDecFromStr("1.5")
	msgAddStakePool := types.MsgAddStakePool{
		Creator:         admin.String(),
		StakeTokenDenom: stakeTokenDenom,
		RewardPoolInfoList: []*types.CreateRewardPoolInfo{
			{
				RewardTokenDenom:  rewardTokenDenom,
				TotalRewardAmount: sdk.NewInt(1e4),
				RewardPerSecond:   sdk.NewInt(2),
				StartTimestamp:    4567,
			},
		},
		StakeItemInfoList: []*types.CreateStakeItemInfo{
			{
				LockSecond:      lockSecond,
				PowerRewardRate: powerRewardRate,
			},
		},
	}
	require.True(t, msgAddStakePool.ValidateBasic() == nil)

	_, found := miningKeeper.GetStakePool(sdkCtx, 0)
	require.False(t, found)
	_, err = srv.AddStakePool(ctx, &msgAddStakePool)
	require.Error(t, err)

	//add stake pool fail, totalrewardAmount less than limit
	msgAddStakePool1 := types.MsgAddStakePool{
		Creator:         miningProvider.String(),
		StakeTokenDenom: stakeTokenDenom,
		RewardPoolInfoList: []*types.CreateRewardPoolInfo{
			{
				RewardTokenDenom:  rewardTokenDenom,
				TotalRewardAmount: sdk.NewInt(1e4 - 1),
				RewardPerSecond:   sdk.NewInt(2),
				StartTimestamp:    4567,
			},
		},
		StakeItemInfoList: []*types.CreateStakeItemInfo{
			{
				LockSecond:      lockSecond,
				PowerRewardRate: powerRewardRate,
			},
		},
	}
	require.True(t, msgAddStakePool1.ValidateBasic() == nil)

	_, found = miningKeeper.GetStakePool(sdkCtx, 0)
	require.False(t, found)
	_, err = srv.AddStakePool(ctx, &msgAddStakePool1)
	require.Error(t, err)

}

func TestAddRewardPoolSuccess(t *testing.T) {
	srv, miningKeeper, ctx, sdkCtx := setupMsgServer(t)
	admin := sample.TestAdminAcc
	stakeTokenDenom := sample.TestDenom
	rewardTokenDenom := sample.TestDenom1
	rewardTokenDenom2 := sample.TestDenom2

	// add rewarder
	miningProvider := sample.OriginAccAddress()
	willMintCoins := sdk.NewCoins(sdk.NewCoin(rewardTokenDenom, sdk.NewInt(1e4)))
	keepertest.BankKeeper.MintCoins(sdkCtx, types.ModuleName, willMintCoins)
	keepertest.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, miningProvider, willMintCoins)

	msgAddRewarder := types.MsgAddMiningProvider{
		Creator:     admin.String(),
		UserAddress: miningProvider.String(),
	}
	_, err := srv.AddMiningProvider(ctx, &msgAddRewarder)
	require.NoError(t, err)

	// add reward token
	msgAddRewardToken := types.MsgAddRewardToken{
		Creator:              admin.String(),
		Denom:                rewardTokenDenom,
		MinTotalRewardAmount: sdk.NewInt(1e4),
	}
	_, err = srv.AddRewardToken(ctx, &msgAddRewardToken)
	require.NoError(t, err)

	msgAddRewardToken2 := types.MsgAddRewardToken{
		Creator:              admin.String(),
		Denom:                rewardTokenDenom2,
		MinTotalRewardAmount: sdk.NewInt(1e4),
	}
	_, err = srv.AddRewardToken(ctx, &msgAddRewardToken2)
	require.NoError(t, err)

	// add stake token
	msgAddStakeToken := types.MsgAddStakeToken{
		Creator: admin.String(),
		Denom:   stakeTokenDenom,
	}
	_, err = srv.AddStakeToken(ctx, &msgAddStakeToken)
	require.NoError(t, err)

	//add stake pool
	lockSecond := uint64(100)
	powerRewardRate := utils.MustNewDecFromStr("1.5")
	msgAddStakePool := types.MsgAddStakePool{
		Creator:         miningProvider.String(),
		StakeTokenDenom: stakeTokenDenom,
		RewardPoolInfoList: []*types.CreateRewardPoolInfo{
			{
				RewardTokenDenom:  rewardTokenDenom,
				TotalRewardAmount: sdk.NewInt(1e4),
				RewardPerSecond:   sdk.NewInt(2),
				StartTimestamp:    4567,
			},
		},
		StakeItemInfoList: []*types.CreateStakeItemInfo{
			{
				LockSecond:      lockSecond,
				PowerRewardRate: powerRewardRate,
			},
		},
	}
	require.True(t, msgAddStakePool.ValidateBasic() == nil)

	_, err = srv.AddStakePool(ctx, &msgAddStakePool)
	require.NoError(t, err)

	// add reward pool
	now := time.Now().Add(time.Second * 10000)
	sdkCtx = sdkCtx.WithBlockTime(now)
	ctx = sdk.WrapSDKContext(sdkCtx)

	userToAddRewardPool := sample.OriginAccAddress()

	totalRewardAmount := sdk.NewInt(1e4)
	rewardPerSecond := sdk.NewInt(10)
	startTimestamp := uint64(2)

	willMintCoins2 := sdk.NewCoins(sdk.NewCoin(rewardTokenDenom2, totalRewardAmount))
	keepertest.BankKeeper.MintCoins(sdkCtx, types.ModuleName, willMintCoins2)
	keepertest.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, userToAddRewardPool, willMintCoins2)

	msgAddRewardPool := types.MsgAddRewardPool{
		Creator:           userToAddRewardPool.String(),
		StakePoolIndex:    0,
		RewardTokenDenom:  rewardTokenDenom2,
		TotalRewardAmount: totalRewardAmount,
		RewardPerSecond:   rewardPerSecond,
		StartTimestamp:    startTimestamp,
	}
	require.True(t, msgAddRewardPool.ValidateBasic() == nil)

	_, err = srv.AddRewardPool(ctx, &msgAddRewardPool)
	require.NoError(t, err)

	stakePool, found := miningKeeper.GetStakePool(sdkCtx, 0)
	require.True(t, found)
	require.Equal(t, len(stakePool.RewardPools), 2)

	require.EqualValues(t, stakePool.RewardPools[1].Index, uint32(1))
	require.EqualValues(t, stakePool.RewardPools[1].LastRewardTimestamp, uint64(now.Unix()))
	require.EqualValues(t, stakePool.RewardPools[1].LeftRewardAmount, totalRewardAmount)
	require.EqualValues(t, stakePool.RewardPools[1].RewardPerPower, sdk.ZeroInt())
	require.EqualValues(t, stakePool.RewardPools[1].RewardPerSecond, rewardPerSecond)
	require.EqualValues(t, stakePool.RewardPools[1].RewardTokenDenom, rewardTokenDenom2)
	require.EqualValues(t, stakePool.RewardPools[1].StartTimestamp, now.Unix())
	require.EqualValues(t, stakePool.RewardPools[1].TotalRewardAmount, totalRewardAmount)
}

func TestAddRewardPoolFail(t *testing.T) {
	srv, _, ctx, sdkCtx := setupMsgServer(t)
	admin := sample.TestAdminAcc
	user := sample.OriginAccAddress()

	stakeTokenDenom := sample.TestDenom
	rewardTokenDenom := sample.TestDenom1
	rewardTokenDenom2 := sample.TestDenom2

	// add rewarder
	miningProvider := sample.OriginAccAddress()
	msgAddRewarder := types.MsgAddMiningProvider{
		Creator:     admin.String(),
		UserAddress: miningProvider.String(),
	}
	_, err := srv.AddMiningProvider(ctx, &msgAddRewarder)
	require.NoError(t, err)

	// add reward token
	msgAddRewardToken := types.MsgAddRewardToken{
		Creator:              admin.String(),
		Denom:                rewardTokenDenom,
		MinTotalRewardAmount: sdk.NewInt(1e4),
	}
	_, err = srv.AddRewardToken(ctx, &msgAddRewardToken)
	require.NoError(t, err)

	willMintCoins := sdk.NewCoins(sdk.NewCoin(rewardTokenDenom, sdk.NewInt(1e4)))
	keepertest.BankKeeper.MintCoins(sdkCtx, types.ModuleName, willMintCoins)
	keepertest.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, miningProvider, willMintCoins)
	keepertest.BankKeeper.MintCoins(sdkCtx, types.ModuleName, willMintCoins)
	keepertest.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, user, willMintCoins)

	msgAddRewardToken2 := types.MsgAddRewardToken{
		Creator:              admin.String(),
		Denom:                rewardTokenDenom2,
		MinTotalRewardAmount: sdk.NewInt(1e4),
	}
	_, err = srv.AddRewardToken(ctx, &msgAddRewardToken2)
	require.NoError(t, err)
	willMintCoins2 := sdk.NewCoins(sdk.NewCoin(rewardTokenDenom2, sdk.NewInt(1e4)))
	keepertest.BankKeeper.MintCoins(sdkCtx, types.ModuleName, willMintCoins2)
	keepertest.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, user, willMintCoins2)

	// add stake token
	msgAddStakeToken := types.MsgAddStakeToken{
		Creator: admin.String(),
		Denom:   stakeTokenDenom,
	}
	_, err = srv.AddStakeToken(ctx, &msgAddStakeToken)
	require.NoError(t, err)

	//add stake pool
	now := time.Now()
	sdkCtx = sdkCtx.WithBlockTime(now)
	ctx = sdk.WrapSDKContext(sdkCtx)

	totalRewardAmount := sdk.NewInt(1e4)
	rewardPerSecond := sdk.NewInt(10)
	startTimestamp := uint64(2)
	lockSecond := uint64(100)
	powerRewardRate := utils.MustNewDecFromStr("1.5")

	msgAddStakePool := types.MsgAddStakePool{
		Creator:         miningProvider.String(),
		StakeTokenDenom: stakeTokenDenom,
		RewardPoolInfoList: []*types.CreateRewardPoolInfo{
			{
				RewardTokenDenom:  rewardTokenDenom,
				TotalRewardAmount: totalRewardAmount,
				RewardPerSecond:   rewardPerSecond,
				StartTimestamp:    startTimestamp,
			},
		},
		StakeItemInfoList: []*types.CreateStakeItemInfo{
			{
				LockSecond:      lockSecond,
				PowerRewardRate: powerRewardRate,
			},
		},
	}
	require.True(t, msgAddStakePool.ValidateBasic() == nil)

	_, err = srv.AddStakePool(ctx, &msgAddStakePool)
	require.NoError(t, err)

	// add reward pool fail
	now = time.Now()
	sdkCtx = sdkCtx.WithBlockTime(now)
	ctx = sdk.WrapSDKContext(sdkCtx)

	testcases := []struct {
		name              string
		creator           string
		stakePoolIndex    uint32
		rewardTokenDenom  string
		totalRewardAmount sdk.Int
		rewardPerSecond   sdk.Int
		startTimestamp    uint64
	}{
		{
			name:              "stake pool not exist",
			creator:           user.String(),
			stakePoolIndex:    1,
			rewardTokenDenom:  rewardTokenDenom2,
			totalRewardAmount: sdk.NewInt(1e4),
			rewardPerSecond:   sdk.NewInt(10),
			startTimestamp:    8,
		},
		{
			name:              "less than min total reward amount",
			creator:           user.String(),
			stakePoolIndex:    0,
			rewardTokenDenom:  rewardTokenDenom2,
			totalRewardAmount: sdk.NewInt(1e4 - 1),
			rewardPerSecond:   sdk.NewInt(10),
			startTimestamp:    8,
		},
		{
			name:              "total reward amount is zero",
			creator:           user.String(),
			stakePoolIndex:    0,
			rewardTokenDenom:  rewardTokenDenom2,
			totalRewardAmount: sdk.NewInt(0),
			rewardPerSecond:   sdk.NewInt(10),
			startTimestamp:    8,
		},
		{
			name:              "reward per second is zero",
			creator:           user.String(),
			stakePoolIndex:    0,
			rewardTokenDenom:  rewardTokenDenom2,
			totalRewardAmount: sdk.NewInt(1e4),
			rewardPerSecond:   sdk.NewInt(0),
			startTimestamp:    8,
		},
		{
			name:              "duplicate denom",
			creator:           user.String(),
			stakePoolIndex:    0,
			rewardTokenDenom:  rewardTokenDenom,
			totalRewardAmount: sdk.NewInt(1e4),
			rewardPerSecond:   sdk.NewInt(2),
			startTimestamp:    8,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			msgAddRewardPool := types.MsgAddRewardPool{
				Creator:           tc.creator,
				StakePoolIndex:    tc.stakePoolIndex,
				RewardTokenDenom:  tc.rewardTokenDenom,
				TotalRewardAmount: tc.totalRewardAmount,
				RewardPerSecond:   tc.rewardPerSecond,
				StartTimestamp:    tc.startTimestamp,
			}

			err := msgAddRewardPool.ValidateBasic()
			if err != nil {
				t.Log(err)
				return
			}

			_, err = srv.AddRewardPool(ctx, &msgAddRewardPool)
			t.Log(err)
			require.Error(t, err)
		})
	}
}

func TestAddRewardSuccess(t *testing.T) {
	srv, miningKeeper, ctx, sdkCtx := setupMsgServer(t)
	admin := sample.TestAdminAcc
	stakeTokenDenom := sample.TestDenom
	rewardTokenDenom := sample.TestDenom1

	// add mining provider
	miningProvider := sample.OriginAccAddress()

	msgAddRewarder := types.MsgAddMiningProvider{
		Creator:     admin.String(),
		UserAddress: miningProvider.String(),
	}
	_, err := srv.AddMiningProvider(ctx, &msgAddRewarder)
	require.NoError(t, err)

	// add reward token
	msgAddRewardToken := types.MsgAddRewardToken{
		Creator:              admin.String(),
		Denom:                rewardTokenDenom,
		MinTotalRewardAmount: sdk.NewInt(1e4),
	}
	_, err = srv.AddRewardToken(ctx, &msgAddRewardToken)
	require.NoError(t, err)

	willMintCoins := sdk.NewCoins(sdk.NewCoin(rewardTokenDenom, sdk.NewInt(1e4)))
	keepertest.BankKeeper.MintCoins(sdkCtx, types.ModuleName, willMintCoins)
	keepertest.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, miningProvider, willMintCoins)

	// add stake token
	msgAddStakeToken := types.MsgAddStakeToken{
		Creator: admin.String(),
		Denom:   stakeTokenDenom,
	}
	_, err = srv.AddStakeToken(ctx, &msgAddStakeToken)
	require.NoError(t, err)

	//add stake pool
	now := time.Now()
	sdkCtx = sdkCtx.WithBlockTime(now)
	ctx = sdk.WrapSDKContext(sdkCtx)

	totalRewardAmount := sdk.NewInt(1e4)
	rewardPerSecond := sdk.NewInt(10)
	startTimestamp := uint64(2)
	lockSecond := uint64(100)

	powerRewardRate := utils.MustNewDecFromStr("1.5")
	msgAddStakePool := types.MsgAddStakePool{
		Creator:         miningProvider.String(),
		StakeTokenDenom: stakeTokenDenom,
		RewardPoolInfoList: []*types.CreateRewardPoolInfo{
			{
				RewardTokenDenom:  rewardTokenDenom,
				TotalRewardAmount: totalRewardAmount,
				RewardPerSecond:   rewardPerSecond,
				StartTimestamp:    startTimestamp,
			},
		},
		StakeItemInfoList: []*types.CreateStakeItemInfo{
			{
				LockSecond:      lockSecond,
				PowerRewardRate: powerRewardRate,
			},
		},
	}
	require.True(t, msgAddStakePool.ValidateBasic() == nil)

	_, err = srv.AddStakePool(ctx, &msgAddStakePool)
	require.NoError(t, err)

	// add reward
	now = time.Now()
	sdkCtx = sdkCtx.WithBlockTime(now)
	ctx = sdk.WrapSDKContext(sdkCtx)

	userToAddReward := sample.OriginAccAddress()

	addRewardAmount := sdk.NewInt(1e4 - 1)
	newRewardPerSecond := sdk.NewInt(0)
	newStartTimestamp := uint64(0)

	willMintCoins = sdk.NewCoins(sdk.NewCoin(rewardTokenDenom, addRewardAmount))
	keepertest.BankKeeper.MintCoins(sdkCtx, types.ModuleName, willMintCoins)
	keepertest.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, userToAddReward, willMintCoins)

	msgAddReward := types.MsgAddReward{
		Creator:         userToAddReward.String(),
		StakePoolIndex:  0,
		RewardPoolIndex: 0,
		AddAmount:       addRewardAmount,
		StartTimestamp:  newStartTimestamp,
		RewardPerSecond: newRewardPerSecond,
	}
	require.True(t, msgAddReward.ValidateBasic() == nil)

	_, err = srv.AddReward(ctx, &msgAddReward)
	require.NoError(t, err)

	stakePool, found := miningKeeper.GetStakePool(sdkCtx, 0)
	require.True(t, found)
	require.Equal(t, len(stakePool.RewardPools), 1)

	require.EqualValues(t, stakePool.RewardPools[0].Index, uint32(0))
	require.EqualValues(t, stakePool.RewardPools[0].LastRewardTimestamp, uint64(now.Unix()))
	require.EqualValues(t, stakePool.RewardPools[0].LeftRewardAmount, totalRewardAmount.Add(addRewardAmount))
	require.EqualValues(t, stakePool.RewardPools[0].RewardPerPower, sdk.ZeroInt())
	require.EqualValues(t, stakePool.RewardPools[0].RewardPerSecond, rewardPerSecond)
	require.EqualValues(t, stakePool.RewardPools[0].RewardTokenDenom, rewardTokenDenom)
	require.EqualValues(t, stakePool.RewardPools[0].StartTimestamp, now.Unix())
	require.EqualValues(t, stakePool.RewardPools[0].TotalRewardAmount, totalRewardAmount.Add(addRewardAmount))
}

func TestAddRewardFail(t *testing.T) {
	srv, _, ctx, sdkCtx := setupMsgServer(t)
	admin := sample.TestAdminAcc
	stakeTokenDenom := sample.TestDenom
	rewardTokenDenom := sample.TestDenom1
	rewardTokenDenom2 := sample.TestDenom2

	// add mining provider
	miningProvider := sample.OriginAccAddress()

	msgAddRewarder := types.MsgAddMiningProvider{
		Creator:     admin.String(),
		UserAddress: miningProvider.String(),
	}
	_, err := srv.AddMiningProvider(ctx, &msgAddRewarder)
	require.NoError(t, err)

	// add reward token
	msgAddRewardToken := types.MsgAddRewardToken{
		Creator:              admin.String(),
		Denom:                rewardTokenDenom,
		MinTotalRewardAmount: sdk.NewInt(1e4),
	}
	_, err = srv.AddRewardToken(ctx, &msgAddRewardToken)
	require.NoError(t, err)

	willMintCoins := sdk.NewCoins(sdk.NewCoin(rewardTokenDenom, sdk.NewInt(1e4)))
	keepertest.BankKeeper.MintCoins(sdkCtx, types.ModuleName, willMintCoins)
	keepertest.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, miningProvider, willMintCoins)

	// add stake token
	msgAddStakeToken := types.MsgAddStakeToken{
		Creator: admin.String(),
		Denom:   stakeTokenDenom,
	}
	_, err = srv.AddStakeToken(ctx, &msgAddStakeToken)
	require.NoError(t, err)

	//add stake pool
	now := time.Now()
	sdkCtx = sdkCtx.WithBlockTime(now)
	ctx = sdk.WrapSDKContext(sdkCtx)

	totalRewardAmount := sdk.NewInt(1e4)
	rewardPerSecond := sdk.NewInt(10)
	startTimestamp := uint64(2)
	lockSecond := uint64(100)

	powerRewardRate := utils.MustNewDecFromStr("1.5")
	msgAddStakePool := types.MsgAddStakePool{
		Creator:         miningProvider.String(),
		StakeTokenDenom: stakeTokenDenom,
		RewardPoolInfoList: []*types.CreateRewardPoolInfo{
			{
				RewardTokenDenom:  rewardTokenDenom,
				TotalRewardAmount: totalRewardAmount,
				RewardPerSecond:   rewardPerSecond,
				StartTimestamp:    startTimestamp,
			},
		},
		StakeItemInfoList: []*types.CreateStakeItemInfo{
			{
				LockSecond:      lockSecond,
				PowerRewardRate: powerRewardRate,
			},
		},
	}
	require.True(t, msgAddStakePool.ValidateBasic() == nil)

	_, err = srv.AddStakePool(ctx, &msgAddStakePool)
	require.NoError(t, err)

	// add reward
	now = time.Now()
	sdkCtx = sdkCtx.WithBlockTime(now)
	ctx = sdk.WrapSDKContext(sdkCtx)

	userToAddReward := sample.OriginAccAddress()

	addRewardAmount := sdk.NewInt(1e4 - 1)

	willMintCoins1 := sdk.NewCoins(sdk.NewCoin(rewardTokenDenom, addRewardAmount))
	keepertest.BankKeeper.MintCoins(sdkCtx, types.ModuleName, willMintCoins1)
	keepertest.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, userToAddReward, willMintCoins1)

	willMintCoins2 := sdk.NewCoins(sdk.NewCoin(rewardTokenDenom2, addRewardAmount))
	keepertest.BankKeeper.MintCoins(sdkCtx, types.ModuleName, willMintCoins2)
	keepertest.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, userToAddReward, willMintCoins2)

	testcases := []struct {
		name            string
		creator         string
		stakePoolIndex  uint32
		rewardPoolIndex uint32
		addAmount       sdk.Int
		rewardPerSecond sdk.Int
		startTimestamp  uint64
	}{
		{
			name:            "stake pool not exist",
			creator:         userToAddReward.String(),
			stakePoolIndex:  1,
			rewardPoolIndex: 0,
			addAmount:       addRewardAmount,
			rewardPerSecond: sdk.NewInt(10),
			startTimestamp:  8,
		},
		{
			name:            "add reward amount is zero",
			creator:         userToAddReward.String(),
			stakePoolIndex:  0,
			rewardPoolIndex: 0,
			addAmount:       sdk.NewInt(0),
			rewardPerSecond: sdk.NewInt(0),
			startTimestamp:  0,
		},
		{
			name:            "reward per second not zero",
			creator:         userToAddReward.String(),
			stakePoolIndex:  0,
			rewardPoolIndex: 0,
			addAmount:       addRewardAmount,
			rewardPerSecond: sdk.NewInt(2),
			startTimestamp:  0,
		},
		{
			name:            "start timestamp not zero",
			creator:         userToAddReward.String(),
			stakePoolIndex:  0,
			rewardPoolIndex: 0,
			addAmount:       addRewardAmount,
			rewardPerSecond: sdk.NewInt(0),
			startTimestamp:  6,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {

			msgAddReward := types.MsgAddReward{
				Creator:         tc.creator,
				StakePoolIndex:  tc.stakePoolIndex,
				RewardPoolIndex: tc.rewardPoolIndex,
				AddAmount:       tc.addAmount,
				StartTimestamp:  tc.startTimestamp,
				RewardPerSecond: tc.rewardPerSecond,
			}
			err := msgAddReward.ValidateBasic()
			if err != nil {
				t.Log(err)
				return
			}

			_, err = srv.AddReward(ctx, &msgAddReward)
			t.Log(err)
			require.Error(t, err)
		})
	}
}

func TestStakeSuccess(t *testing.T) {
	srv, miningKeeper, ctx, sdkCtx := setupMsgServer(t)
	admin := sample.TestAdminAcc
	stakeTokenDenom := sample.TestDenom
	rewardTokenDenom := sample.TestDenom1

	// add mining provider
	miningProvider := sample.OriginAccAddress()

	willMintCoins := sdk.NewCoins(sdk.NewCoin(rewardTokenDenom, sdk.NewInt(1e4)))
	keepertest.BankKeeper.MintCoins(sdkCtx, types.ModuleName, willMintCoins)
	keepertest.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, miningProvider, willMintCoins)

	msgAddMiningProvider := types.MsgAddMiningProvider{
		Creator:     admin.String(),
		UserAddress: miningProvider.String(),
	}
	_, err := srv.AddMiningProvider(ctx, &msgAddMiningProvider)
	require.NoError(t, err)

	// add reward token
	msgAddRewardToken := types.MsgAddRewardToken{
		Creator:              admin.String(),
		Denom:                rewardTokenDenom,
		MinTotalRewardAmount: sdk.NewInt(1e4),
	}
	_, err = srv.AddRewardToken(ctx, &msgAddRewardToken)
	require.NoError(t, err)

	// add stake token
	msgAddStakeToken := types.MsgAddStakeToken{
		Creator: admin.String(),
		Denom:   stakeTokenDenom,
	}
	_, err = srv.AddStakeToken(ctx, &msgAddStakeToken)
	require.NoError(t, err)

	//add stake pool
	now := time.Now()
	sdkCtx = sdkCtx.WithBlockTime(now)
	ctx = sdk.WrapSDKContext(sdkCtx)

	lockSecond := uint64(100)
	powerRewardRate := utils.MustNewDecFromStr("1.5")
	msgAddStakePool := types.MsgAddStakePool{
		Creator:         miningProvider.String(),
		StakeTokenDenom: stakeTokenDenom,
		RewardPoolInfoList: []*types.CreateRewardPoolInfo{
			{
				RewardTokenDenom:  rewardTokenDenom,
				TotalRewardAmount: sdk.NewInt(1e4),
				RewardPerSecond:   sdk.NewInt(2),
				StartTimestamp:    4567,
			},
		},
		StakeItemInfoList: []*types.CreateStakeItemInfo{
			{
				LockSecond:      lockSecond,
				PowerRewardRate: powerRewardRate,
			},
		},
	}

	require.True(t, msgAddStakePool.ValidateBasic() == nil)

	_, err = srv.AddStakePool(ctx, &msgAddStakePool)
	require.NoError(t, err)

	// stake
	now = time.Now()
	sdkCtx = sdkCtx.WithBlockTime(now)
	ctx = sdk.WrapSDKContext(sdkCtx)

	stakeUser := sample.OriginAccAddress()

	willMintCoins = sdk.NewCoins(sdk.NewCoin(stakeTokenDenom, sdk.NewInt(10)))
	keepertest.BankKeeper.MintCoins(sdkCtx, types.ModuleName, willMintCoins)
	keepertest.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, stakeUser, willMintCoins)

	msgStake := types.MsgStake{
		Creator:        stakeUser.String(),
		StakePoolIndex: 0,
		StakeAmount:    sdk.NewInt(10),
		StakeItemIndex: 0,
	}

	_, err = srv.Stake(ctx, &msgStake)
	require.NoError(t, err)

	stakeRecord, found := miningKeeper.GetUserStakeRecord(sdkCtx, stakeUser.String(), 0, 0)
	require.True(t, found)

	require.EqualValues(t, stakeRecord.Index, 0)
	require.EqualValues(t, stakeRecord.LockEndTimestamp, now.Unix()+int64(lockSecond))
	require.EqualValues(t, stakeRecord.StakeItemIndex, 0)
	require.EqualValues(t, stakeRecord.StakePoolIndex, 0)
	require.EqualValues(t, stakeRecord.StakedAmount, sdk.NewInt(10))
	require.EqualValues(t, stakeRecord.StakedPower, sdk.NewInt(15))
	require.EqualValues(t, stakeRecord.StartTimestamp, now.Unix())
	require.EqualValues(t, stakeRecord.UserAddress, stakeUser.String())
	require.EqualValues(t, len(stakeRecord.UserRewardInfos), 1)
	require.EqualValues(t, stakeRecord.UserRewardInfos[0].ClaimedAmount, sdk.ZeroInt())
	require.EqualValues(t, stakeRecord.UserRewardInfos[0].RewardDebt, sdk.ZeroInt())
	require.EqualValues(t, stakeRecord.UserRewardInfos[0].RewardPoolIndex, 0)
	require.EqualValues(t, stakeRecord.UserRewardInfos[0].RewardTokenDenom, rewardTokenDenom)
}

func TestClaimRewardSuccess(t *testing.T) {
	srv, miningKeeper, ctx, sdkCtx := setupMsgServer(t)
	admin := sample.TestAdminAcc
	stakeTokenDenom := sample.TestDenom
	rewardTokenDenom := sample.TestDenom1

	// add mining provider
	miningProvider := sample.OriginAccAddress()

	willMintCoins := sdk.NewCoins(sdk.NewCoin(rewardTokenDenom, sdk.NewInt(1e4)))
	keepertest.BankKeeper.MintCoins(sdkCtx, types.ModuleName, willMintCoins)
	keepertest.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, miningProvider, willMintCoins)

	msgAddMiningProvider := types.MsgAddMiningProvider{
		Creator:     admin.String(),
		UserAddress: miningProvider.String(),
	}
	_, err := srv.AddMiningProvider(ctx, &msgAddMiningProvider)
	require.NoError(t, err)

	// add reward token
	msgAddRewardToken := types.MsgAddRewardToken{
		Creator:              admin.String(),
		Denom:                rewardTokenDenom,
		MinTotalRewardAmount: sdk.NewInt(1e4),
	}
	_, err = srv.AddRewardToken(ctx, &msgAddRewardToken)
	require.NoError(t, err)

	// add stake token
	msgAddStakeToken := types.MsgAddStakeToken{
		Creator: admin.String(),
		Denom:   stakeTokenDenom,
	}
	_, err = srv.AddStakeToken(ctx, &msgAddStakeToken)
	require.NoError(t, err)

	//add stake pool
	now := time.Now()
	sdkCtx = sdkCtx.WithBlockTime(now)
	ctx = sdk.WrapSDKContext(sdkCtx)

	lockSecond := uint64(100)
	powerRewardRate := utils.MustNewDecFromStr("1.5")
	msgAddStakePool := types.MsgAddStakePool{
		Creator:         miningProvider.String(),
		StakeTokenDenom: stakeTokenDenom,
		RewardPoolInfoList: []*types.CreateRewardPoolInfo{
			{
				RewardTokenDenom:  rewardTokenDenom,
				TotalRewardAmount: sdk.NewInt(1e4),
				RewardPerSecond:   sdk.NewInt(2),
				StartTimestamp:    4567,
			},
		},
		StakeItemInfoList: []*types.CreateStakeItemInfo{
			{
				LockSecond:      lockSecond,
				PowerRewardRate: powerRewardRate,
			},
		},
	}

	require.True(t, msgAddStakePool.ValidateBasic() == nil)

	_, err = srv.AddStakePool(ctx, &msgAddStakePool)
	require.NoError(t, err)

	// stake
	now = time.Now()
	sdkCtx = sdkCtx.WithBlockTime(now)
	ctx = sdk.WrapSDKContext(sdkCtx)

	stakeUser := sample.OriginAccAddress()

	willMintCoins = sdk.NewCoins(sdk.NewCoin(stakeTokenDenom, sdk.NewInt(10)))
	keepertest.BankKeeper.MintCoins(sdkCtx, types.ModuleName, willMintCoins)
	keepertest.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, stakeUser, willMintCoins)

	msgStake := types.MsgStake{
		Creator:        stakeUser.String(),
		StakePoolIndex: 0,
		StakeAmount:    sdk.NewInt(10),
		StakeItemIndex: 0,
	}

	_, err = srv.Stake(ctx, &msgStake)
	require.NoError(t, err)

	// claim reward
	duration := 100
	now = now.Add(time.Second * time.Duration(duration))
	sdkCtx = sdkCtx.WithBlockTime(now)
	ctx = sdk.WrapSDKContext(sdkCtx)

	stakeReward, err := miningKeeper.StakeReward(ctx, &types.QueryStakeRewardRequest{
		StakeUserAddress: stakeUser.String(),
		StakePoolIndex:   0,
		StakeRecordIndex: 0,
	})
	require.NoError(t, err)

	require.EqualValues(t, len(stakeReward.RewardTokens), 1)
	require.EqualValues(t, stakeReward.RewardTokens[0].Amount, sdk.NewInt(199))
	require.EqualValues(t, stakeReward.RewardTokens[0].Denom, rewardTokenDenom)

	msgClaimReward := types.MsgClaimReward{
		Creator:          stakeUser.String(),
		StakePoolIndex:   0,
		StakeRecordIndex: 0,
	}

	_, err = srv.ClaimReward(ctx, &msgClaimReward)
	require.NoError(t, err)

	balance := keepertest.BankKeeper.GetBalance(sdkCtx, stakeUser, rewardTokenDenom)
	require.EqualValues(t, balance.Amount, sdk.NewInt(199))

	rspStakeRecord, err := miningKeeper.StakeRecord(ctx, &types.QueryStakeRecordRequest{
		UserAddress:      stakeUser.String(),
		StakePoolIndex:   0,
		StakeRecordIndex: 0,
	})
	require.NoError(t, err)

	require.EqualValues(t, len(rspStakeRecord.StakeRecord.UserRewardInfos), 1)
	require.EqualValues(t, rspStakeRecord.StakeRecord.UserRewardInfos[0].ClaimedAmount, sdk.NewInt(199))

}

func TestWithdrawSuccess(t *testing.T) {
	srv, miningKeeper, ctx, sdkCtx := setupMsgServer(t)
	admin := sample.TestAdminAcc
	stakeTokenDenom := sample.TestDenom
	rewardTokenDenom := sample.TestDenom1

	// add mining provider
	miningProvider := sample.OriginAccAddress()

	willMintCoins := sdk.NewCoins(sdk.NewCoin(rewardTokenDenom, sdk.NewInt(1e4)))
	keepertest.BankKeeper.MintCoins(sdkCtx, types.ModuleName, willMintCoins)
	keepertest.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, miningProvider, willMintCoins)

	msgAddMiningProvider := types.MsgAddMiningProvider{
		Creator:     admin.String(),
		UserAddress: miningProvider.String(),
	}
	_, err := srv.AddMiningProvider(ctx, &msgAddMiningProvider)
	require.NoError(t, err)

	// add reward token
	msgAddRewardToken := types.MsgAddRewardToken{
		Creator:              admin.String(),
		Denom:                rewardTokenDenom,
		MinTotalRewardAmount: sdk.NewInt(1e4),
	}
	_, err = srv.AddRewardToken(ctx, &msgAddRewardToken)
	require.NoError(t, err)

	// add stake token
	msgAddStakeToken := types.MsgAddStakeToken{
		Creator: admin.String(),
		Denom:   stakeTokenDenom,
	}
	_, err = srv.AddStakeToken(ctx, &msgAddStakeToken)
	require.NoError(t, err)

	//add stake pool
	now := time.Now()
	sdkCtx = sdkCtx.WithBlockTime(now)
	ctx = sdk.WrapSDKContext(sdkCtx)

	lockSecond := uint64(100)
	powerRewardRate := utils.MustNewDecFromStr("1.5")
	msgAddStakePool := types.MsgAddStakePool{
		Creator:         miningProvider.String(),
		StakeTokenDenom: stakeTokenDenom,
		RewardPoolInfoList: []*types.CreateRewardPoolInfo{
			{
				RewardTokenDenom:  rewardTokenDenom,
				TotalRewardAmount: sdk.NewInt(1e4),
				RewardPerSecond:   sdk.NewInt(2),
				StartTimestamp:    4567,
			},
		},
		StakeItemInfoList: []*types.CreateStakeItemInfo{
			{
				LockSecond:      lockSecond,
				PowerRewardRate: powerRewardRate,
			},
		},
	}

	require.True(t, msgAddStakePool.ValidateBasic() == nil)

	_, err = srv.AddStakePool(ctx, &msgAddStakePool)
	require.NoError(t, err)

	// stake
	now = time.Now()
	sdkCtx = sdkCtx.WithBlockTime(now)
	ctx = sdk.WrapSDKContext(sdkCtx)

	stakeUser := sample.OriginAccAddress()

	willMintCoins = sdk.NewCoins(sdk.NewCoin(stakeTokenDenom, sdk.NewInt(10)))
	keepertest.BankKeeper.MintCoins(sdkCtx, types.ModuleName, willMintCoins)
	keepertest.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, stakeUser, willMintCoins)

	msgStake := types.MsgStake{
		Creator:        stakeUser.String(),
		StakePoolIndex: 0,
		StakeAmount:    sdk.NewInt(10),
		StakeItemIndex: 0,
	}

	_, err = srv.Stake(ctx, &msgStake)
	require.NoError(t, err)

	// claim reward
	duration := 100
	now = now.Add(time.Second * time.Duration(duration))
	sdkCtx = sdkCtx.WithBlockTime(now)
	ctx = sdk.WrapSDKContext(sdkCtx)

	stakeReward, err := miningKeeper.StakeReward(ctx, &types.QueryStakeRewardRequest{
		StakeUserAddress: stakeUser.String(),
		StakePoolIndex:   0,
		StakeRecordIndex: 0,
	})
	require.NoError(t, err)

	require.EqualValues(t, len(stakeReward.RewardTokens), 1)
	require.EqualValues(t, stakeReward.RewardTokens[0].Amount, sdk.NewInt(199))
	require.EqualValues(t, stakeReward.RewardTokens[0].Denom, rewardTokenDenom)

	msgClaimReward := types.MsgClaimReward{
		Creator:          stakeUser.String(),
		StakePoolIndex:   0,
		StakeRecordIndex: 0,
	}

	_, err = srv.ClaimReward(ctx, &msgClaimReward)
	require.NoError(t, err)

	balance := keepertest.BankKeeper.GetBalance(sdkCtx, stakeUser, rewardTokenDenom)
	require.EqualValues(t, balance.Amount, sdk.NewInt(199))

	//withdraw
	duration = 100
	now = now.Add(time.Second * time.Duration(duration))
	sdkCtx = sdkCtx.WithBlockTime(now)
	ctx = sdk.WrapSDKContext(sdkCtx)
	stakeReward, err = miningKeeper.StakeReward(ctx, &types.QueryStakeRewardRequest{
		StakeUserAddress: stakeUser.String(),
		StakePoolIndex:   0,
		StakeRecordIndex: 0,
	})
	require.NoError(t, err)

	require.EqualValues(t, len(stakeReward.RewardTokens), 1)
	require.EqualValues(t, stakeReward.RewardTokens[0].Amount, sdk.NewInt(200))
	require.EqualValues(t, stakeReward.RewardTokens[0].Denom, rewardTokenDenom)

	msgWithdraw := types.MsgWithdraw{
		Creator:          stakeUser.String(),
		StakePoolIndex:   0,
		StakeRecordIndex: 0,
		WithdrawAmount:   sdk.NewInt(10),
	}

	_, err = srv.Withdraw(ctx, &msgWithdraw)
	require.NoError(t, err)

	balance = keepertest.BankKeeper.GetBalance(sdkCtx, stakeUser, rewardTokenDenom)
	require.EqualValues(t, balance.Amount, sdk.NewInt(399))
	balance = keepertest.BankKeeper.GetBalance(sdkCtx, stakeUser, stakeTokenDenom)
	require.EqualValues(t, balance.Amount, sdk.NewInt(10))

	rspStakeRecord, err := miningKeeper.StakeRecord(ctx, &types.QueryStakeRecordRequest{
		UserAddress:      stakeUser.String(),
		StakePoolIndex:   0,
		StakeRecordIndex: 0,
	})
	require.NoError(t, err)

	require.EqualValues(t, len(rspStakeRecord.StakeRecord.UserRewardInfos), 1)
	require.EqualValues(t, rspStakeRecord.StakeRecord.UserRewardInfos[0].ClaimedAmount, sdk.NewInt(399))
}

func TestClaimMultiRewardSuccess(t *testing.T) {
	srv, miningKeeper, ctx, sdkCtx := setupMsgServer(t)
	admin := sample.TestAdminAcc
	stakeTokenDenom := sample.TestDenom
	rewardTokenDenom := sample.TestDenom1
	rewardTokenDenom2 := sample.TestDenom2

	// add mining provider
	miningProvider := sample.OriginAccAddress()

	willMintCoins := sdk.NewCoins(sdk.NewCoin(rewardTokenDenom, sdk.NewInt(1e4)))
	keepertest.BankKeeper.MintCoins(sdkCtx, types.ModuleName, willMintCoins)
	keepertest.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, miningProvider, willMintCoins)

	msgAddMiningProvider := types.MsgAddMiningProvider{
		Creator:     admin.String(),
		UserAddress: miningProvider.String(),
	}
	_, err := srv.AddMiningProvider(ctx, &msgAddMiningProvider)
	require.NoError(t, err)

	// add reward token
	msgAddRewardToken := types.MsgAddRewardToken{
		Creator:              admin.String(),
		Denom:                rewardTokenDenom,
		MinTotalRewardAmount: sdk.NewInt(1e4),
	}
	_, err = srv.AddRewardToken(ctx, &msgAddRewardToken)
	require.NoError(t, err)

	msgAddRewardToken = types.MsgAddRewardToken{
		Creator:              admin.String(),
		Denom:                rewardTokenDenom2,
		MinTotalRewardAmount: sdk.NewInt(1e4),
	}
	_, err = srv.AddRewardToken(ctx, &msgAddRewardToken)
	require.NoError(t, err)

	// add stake token
	msgAddStakeToken := types.MsgAddStakeToken{
		Creator: admin.String(),
		Denom:   stakeTokenDenom,
	}
	_, err = srv.AddStakeToken(ctx, &msgAddStakeToken)
	require.NoError(t, err)

	//add stake pool
	now := time.Now()
	sdkCtx = sdkCtx.WithBlockTime(now)
	ctx = sdk.WrapSDKContext(sdkCtx)

	lockSecond := uint64(100)
	powerRewardRate := utils.MustNewDecFromStr("1.5")
	msgAddStakePool := types.MsgAddStakePool{
		Creator:         miningProvider.String(),
		StakeTokenDenom: stakeTokenDenom,
		RewardPoolInfoList: []*types.CreateRewardPoolInfo{
			{
				RewardTokenDenom:  rewardTokenDenom,
				TotalRewardAmount: sdk.NewInt(1e4),
				RewardPerSecond:   sdk.NewInt(2),
				StartTimestamp:    4567,
			},
		},
		StakeItemInfoList: []*types.CreateStakeItemInfo{
			{
				LockSecond:      lockSecond,
				PowerRewardRate: powerRewardRate,
			},
		},
	}
	require.True(t, msgAddStakePool.ValidateBasic() == nil)

	_, err = srv.AddStakePool(ctx, &msgAddStakePool)
	require.NoError(t, err)

	// stake
	now = time.Now()
	sdkCtx = sdkCtx.WithBlockTime(now)
	ctx = sdk.WrapSDKContext(sdkCtx)

	stakeUser := sample.OriginAccAddress()

	willMintCoins = sdk.NewCoins(sdk.NewCoin(stakeTokenDenom, sdk.NewInt(10)))
	keepertest.BankKeeper.MintCoins(sdkCtx, types.ModuleName, willMintCoins)
	keepertest.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, stakeUser, willMintCoins)

	msgStake := types.MsgStake{
		Creator:        stakeUser.String(),
		StakePoolIndex: 0,
		StakeAmount:    sdk.NewInt(10),
		StakeItemIndex: 0,
	}

	_, err = srv.Stake(ctx, &msgStake)
	require.NoError(t, err)

	// add reward
	duration := 50
	now = now.Add(time.Second * time.Duration(duration))
	sdkCtx = sdkCtx.WithBlockTime(now)
	ctx = sdk.WrapSDKContext(sdkCtx)

	userToAddReward := sample.OriginAccAddress()

	addRewardAmount := sdk.NewInt(1e4)
	newRewardPerSecond := sdk.NewInt(1)
	newStartTimestamp := uint64(0)

	willMintCoins = sdk.NewCoins(sdk.NewCoin(rewardTokenDenom2, addRewardAmount))
	keepertest.BankKeeper.MintCoins(sdkCtx, types.ModuleName, willMintCoins)
	keepertest.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, userToAddReward, willMintCoins)

	msgAddReward := types.MsgAddRewardPool{
		Creator:           userToAddReward.String(),
		StakePoolIndex:    0,
		RewardTokenDenom:  rewardTokenDenom2,
		TotalRewardAmount: addRewardAmount,
		StartTimestamp:    newStartTimestamp,
		RewardPerSecond:   newRewardPerSecond,
	}
	require.True(t, msgAddReward.ValidateBasic() == nil)

	_, err = srv.AddRewardPool(ctx, &msgAddReward)
	require.NoError(t, err)

	// claim reward
	duration = 50
	now = now.Add(time.Second * time.Duration(duration))
	sdkCtx = sdkCtx.WithBlockTime(now)
	ctx = sdk.WrapSDKContext(sdkCtx)

	stakeReward, err := miningKeeper.StakeReward(ctx, &types.QueryStakeRewardRequest{
		StakeUserAddress: stakeUser.String(),
		StakePoolIndex:   0,
		StakeRecordIndex: 0,
	})
	require.NoError(t, err)

	require.EqualValues(t, len(stakeReward.RewardTokens), 2)
	require.EqualValues(t, stakeReward.RewardTokens[0].Amount, sdk.NewInt(199))
	require.EqualValues(t, stakeReward.RewardTokens[0].Denom, rewardTokenDenom)
	require.EqualValues(t, stakeReward.RewardTokens[1].Amount, sdk.NewInt(49))
	require.EqualValues(t, stakeReward.RewardTokens[1].Denom, rewardTokenDenom2)

	msgClaimReward := types.MsgClaimReward{
		Creator:          stakeUser.String(),
		StakePoolIndex:   0,
		StakeRecordIndex: 0,
	}

	_, err = srv.ClaimReward(ctx, &msgClaimReward)
	require.NoError(t, err)

	balance := keepertest.BankKeeper.GetBalance(sdkCtx, stakeUser, rewardTokenDenom)
	require.EqualValues(t, balance.Amount, sdk.NewInt(199))
	balance = keepertest.BankKeeper.GetBalance(sdkCtx, stakeUser, rewardTokenDenom2)
	require.EqualValues(t, balance.Amount, sdk.NewInt(49))

	rspStakeRecord, err := miningKeeper.StakeRecord(ctx, &types.QueryStakeRecordRequest{
		UserAddress:      stakeUser.String(),
		StakePoolIndex:   0,
		StakeRecordIndex: 0,
	})
	require.NoError(t, err)

	require.EqualValues(t, len(rspStakeRecord.StakeRecord.UserRewardInfos), 2)
	require.EqualValues(t, rspStakeRecord.StakeRecord.UserRewardInfos[0].ClaimedAmount, sdk.NewInt(199))
	require.EqualValues(t, rspStakeRecord.StakeRecord.UserRewardInfos[1].ClaimedAmount, sdk.NewInt(49))

}
