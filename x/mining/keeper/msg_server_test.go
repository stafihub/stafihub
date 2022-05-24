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

	// add stakeItem fail when user is not admin
	msgAddStakeItem := types.MsgAddStakeItem{
		Creator:         sample.AccAddress(),
		LockSecond:      0,
		PowerRewardRate: utils.MustNewDecFromStr("1.5"),
		Enable:          true,
	}
	_, err := srv.AddStakeItem(ctx, &msgAddStakeItem)
	require.Error(t, err)

	// add stakeItem
	powerRewardRate := utils.MustNewDecFromStr("2.5")
	msgAddStakeItem2 := types.MsgAddStakeItem{
		Creator:         admin.String(),
		LockSecond:      0,
		PowerRewardRate: powerRewardRate,
		Enable:          true,
	}
	_, err = srv.AddStakeItem(ctx, &msgAddStakeItem2)
	require.NoError(t, err)
	stakeItem, found := miningKeeper.GetStakeItem(sdkCtx, 0)
	require.True(t, found)
	require.EqualValues(t, stakeItem.Index, 0)
	require.EqualValues(t, stakeItem.Enable, true)
	require.EqualValues(t, stakeItem.LockSecond, 0)
	require.EqualValues(t, stakeItem.PowerRewardRate, powerRewardRate)

	// update Stake item fail whe user is not admin
	newPowerRewardRate := utils.MustNewDecFromStr("1.5")
	msgUpdateStakeItem := types.MsgUpdateStakeItem{
		Creator:         sample.AccAddress(),
		Index:           0,
		LockSecond:      5,
		PowerRewardRate: newPowerRewardRate,
		Enable:          false,
	}
	_, err = srv.UpdateStakeItem(ctx, &msgUpdateStakeItem)
	require.Error(t, err)

	// update Stake item
	msgUpdateStakeItem.Creator = admin.String()
	_, err = srv.UpdateStakeItem(ctx, &msgUpdateStakeItem)
	require.NoError(t, err)
	stakeItem, found = miningKeeper.GetStakeItem(sdkCtx, 0)
	require.True(t, found)
	require.EqualValues(t, stakeItem.Index, 0)
	require.EqualValues(t, stakeItem.Enable, false)
	require.EqualValues(t, stakeItem.LockSecond, 5)
	require.EqualValues(t, stakeItem.PowerRewardRate, newPowerRewardRate)

	nextIndex := miningKeeper.GetStakeItemNextIndex(sdkCtx)
	require.EqualValues(t, nextIndex, 1)

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

	//add stake pool
	msgAddStakePool := types.MsgAddStakePool{
		Creator:           miningProvider.String(),
		StakeTokenDenom:   stakeTokenDenom,
		RewardTokenDenom:  rewardTokenDenom,
		TotalRewardAmount: sdk.NewInt(1e4),
		RewardPerSecond:   sdk.NewInt(2),
		StartTimestamp:    4567,
	}

	_, found := miningKeeper.GetStakePool(sdkCtx, 0)
	require.False(t, found)
	_, err = srv.AddStakePool(ctx, &msgAddStakePool)
	require.NoError(t, err)

	stakePool, found := miningKeeper.GetStakePool(sdkCtx, 0)
	require.True(t, found)
	require.EqualValues(t, stakePool.StakeTokenDenom, msgAddStakePool.StakeTokenDenom)
	require.EqualValues(t, stakePool.TotalStakedAmount, sdk.ZeroInt())
	require.EqualValues(t, stakePool.TotalStakedPower, sdk.ZeroInt())
}

func TestAddStakePoolFail(t *testing.T) {
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

	//add stake pool fail
	msgAddStakePool := types.MsgAddStakePool{
		Creator:           admin.String(),
		StakeTokenDenom:   stakeTokenDenom,
		RewardTokenDenom:  rewardTokenDenom,
		TotalRewardAmount: sdk.NewInt(1e4),
		RewardPerSecond:   sdk.NewInt(2),
		StartTimestamp:    4567,
	}

	_, found := miningKeeper.GetStakePool(sdkCtx, 0)
	require.False(t, found)
	_, err = srv.AddStakePool(ctx, &msgAddStakePool)
	require.Error(t, err)

	//add stake pool fail
	msgAddStakePool1 := types.MsgAddStakePool{
		Creator:           miningProvider.String(),
		StakeTokenDenom:   stakeTokenDenom,
		RewardTokenDenom:  rewardTokenDenom,
		TotalRewardAmount: sdk.NewInt(1e4 - 1),
		RewardPerSecond:   sdk.NewInt(2),
		StartTimestamp:    4567,
	}

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

	//add stake pool
	msgAddStakePool := types.MsgAddStakePool{
		Creator:           miningProvider.String(),
		StakeTokenDenom:   stakeTokenDenom,
		RewardTokenDenom:  rewardTokenDenom,
		TotalRewardAmount: sdk.NewInt(1e4),
		RewardPerSecond:   sdk.NewInt(2),
		StartTimestamp:    4567,
	}

	_, found := miningKeeper.GetStakePool(sdkCtx, 0)
	require.False(t, found)
	_, err = srv.AddStakePool(ctx, &msgAddStakePool)
	require.NoError(t, err)

	stakePool, found := miningKeeper.GetStakePool(sdkCtx, 0)
	require.True(t, found)
	require.EqualValues(t, stakePool.StakeTokenDenom, msgAddStakePool.StakeTokenDenom)
	require.EqualValues(t, stakePool.TotalStakedAmount, sdk.ZeroInt())
	require.EqualValues(t, stakePool.TotalStakedPower, sdk.ZeroInt())

	// add reward pool
	now := time.Now()
	sdkCtx = sdkCtx.WithBlockTime(now)
	ctx = sdk.WrapSDKContext(sdkCtx)

	userToAddRewardPool := sample.OriginAccAddress()

	totalRewardAmount := sdk.NewInt(1e4)
	rewardPerSecond := sdk.NewInt(10)
	startTimestamp := uint64(2)

	willMintCoins = sdk.NewCoins(sdk.NewCoin(rewardTokenDenom, totalRewardAmount))
	keepertest.BankKeeper.MintCoins(sdkCtx, types.ModuleName, willMintCoins)
	keepertest.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, userToAddRewardPool, willMintCoins)

	msgAddRewardPool := types.MsgAddRewardPool{
		Creator:           userToAddRewardPool.String(),
		StakePoolIndex:    0,
		RewardTokenDenom:  rewardTokenDenom,
		TotalRewardAmount: totalRewardAmount,
		RewardPerSecond:   rewardPerSecond,
		StartTimestamp:    startTimestamp,
	}
	_, err = srv.AddRewardPool(ctx, &msgAddRewardPool)
	require.NoError(t, err)

	stakePool, found = miningKeeper.GetStakePool(sdkCtx, 0)
	require.True(t, found)
	require.Equal(t, len(stakePool.RewardPools), 2)

	require.EqualValues(t, stakePool.RewardPools[1].Index, uint32(1))
	require.EqualValues(t, stakePool.RewardPools[1].LastRewardTimestamp, uint64(now.Unix()))
	require.EqualValues(t, stakePool.RewardPools[1].LeftRewardAmount, totalRewardAmount)
	require.EqualValues(t, stakePool.RewardPools[1].RewardPerPower, sdk.ZeroInt())
	require.EqualValues(t, stakePool.RewardPools[1].RewardPerSecond, rewardPerSecond)
	require.EqualValues(t, stakePool.RewardPools[1].RewardTokenDenom, rewardTokenDenom)
	require.EqualValues(t, stakePool.RewardPools[1].StartTimestamp, startTimestamp)
	require.EqualValues(t, stakePool.RewardPools[1].TotalRewardAmount, totalRewardAmount)
}

func TestAddRewardPoolFail(t *testing.T) {
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
	//add stake pool
	msgAddStakePool := types.MsgAddStakePool{
		Creator:           miningProvider.String(),
		StakeTokenDenom:   stakeTokenDenom,
		RewardTokenDenom:  rewardTokenDenom,
		TotalRewardAmount: sdk.NewInt(1e4),
		RewardPerSecond:   sdk.NewInt(2),
		StartTimestamp:    4567,
	}

	_, found := miningKeeper.GetStakePool(sdkCtx, 0)
	require.False(t, found)
	_, err = srv.AddStakePool(ctx, &msgAddStakePool)
	require.NoError(t, err)

	stakePool, found := miningKeeper.GetStakePool(sdkCtx, 0)
	require.True(t, found)
	require.EqualValues(t, stakePool.StakeTokenDenom, msgAddStakePool.StakeTokenDenom)
	require.EqualValues(t, stakePool.TotalStakedAmount, sdk.ZeroInt())
	require.EqualValues(t, stakePool.TotalStakedPower, sdk.ZeroInt())

	// add reward pool fail
	now := time.Now()
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
			name:              "not admin not rewarder",
			creator:           sample.AccAddress(),
			stakePoolIndex:    0,
			rewardTokenDenom:  sample.TestDenom1,
			totalRewardAmount: sdk.NewInt(1e4),
			rewardPerSecond:   sdk.NewInt(10),
			startTimestamp:    8,
		},
		{
			name:              "stake pool not exist",
			creator:           admin.String(),
			stakePoolIndex:    1,
			rewardTokenDenom:  sample.TestDenom1,
			totalRewardAmount: sdk.NewInt(1e4),
			rewardPerSecond:   sdk.NewInt(10),
			startTimestamp:    8,
		},
		{
			name:              "less than min total reward amount",
			creator:           admin.String(),
			stakePoolIndex:    0,
			rewardTokenDenom:  sample.TestDenom1,
			totalRewardAmount: sdk.NewInt(1e4 - 1),
			rewardPerSecond:   sdk.NewInt(10),
			startTimestamp:    8,
		},
		{
			name:              "total reward amount is zero",
			creator:           admin.String(),
			stakePoolIndex:    0,
			rewardTokenDenom:  sample.TestDenom1,
			totalRewardAmount: sdk.NewInt(0),
			rewardPerSecond:   sdk.NewInt(10),
			startTimestamp:    8,
		},
		{
			name:              "reward per second is zero",
			creator:           admin.String(),
			stakePoolIndex:    0,
			rewardTokenDenom:  sample.TestDenom1,
			totalRewardAmount: sdk.NewInt(1e4),
			rewardPerSecond:   sdk.NewInt(0),
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
			require.Error(t, err)
		})
	}
}
