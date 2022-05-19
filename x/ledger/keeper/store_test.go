package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	testkeeper "github.com/stafihub/stafihub/testutil/keeper"
	"github.com/stafihub/stafihub/testutil/sample"
	"github.com/stafihub/stafihub/utils"
	"github.com/stafihub/stafihub/x/ledger/types"
	"github.com/stretchr/testify/require"
)

func TestKeeper_AddBondedPool(t *testing.T) {
	k, ctx := testkeeper.LedgerKeeper(t)

	pool := sample.AccAddress()
	require.False(t, k.IsBondedPoolExist(ctx, sample.TestDenom, pool))

	_, found := k.GetBondedPool(ctx, sample.TestDenom)
	require.False(t, found)

	k.AddBondedPool(ctx, sample.TestDenom, pool)
	require.True(t, k.IsBondedPoolExist(ctx, sample.TestDenom, pool))

	bpl, found := k.GetBondedPool(ctx, sample.TestDenom)
	require.True(t, found)
	require.Equal(t, types.Pool{Denom: sample.TestDenom, Addrs: []string{pool}}, bpl)

	k.RemoveBondedPool(ctx, sample.TestDenom, pool)
	require.False(t, k.IsBondedPoolExist(ctx, sample.TestDenom, pool))

	bpl, found = k.GetBondedPool(ctx, sample.TestDenom)
	require.True(t, found)
	require.Equal(t, types.Pool{Denom: sample.TestDenom, Addrs: []string{}}, bpl)
}

func TestKeeper_BondPipeline(t *testing.T) {
	k, ctx := testkeeper.LedgerKeeper(t)

	pool := sample.AccAddress()
	_, found := k.GetBondPipeline(ctx, sample.TestDenom, pool)
	require.False(t, found)

	bp := types.NewBondPipeline(sample.TestDenom, pool)
	k.SetBondPipeline(ctx, bp)

	bpl, found := k.GetBondPipeline(ctx, sample.TestDenom, pool)
	require.True(t, found)
	require.Equal(t, bp, bpl)
}

func TestKeeper_SetEraUnbondLimit(t *testing.T) {
	k, ctx := testkeeper.LedgerKeeper(t)

	_ = k.GetEraUnbondLimit(ctx, sample.TestDenom)

	limit := uint32(30)
	k.SetEraUnbondLimit(ctx, sample.TestDenom, limit)

	eul := k.GetEraUnbondLimit(ctx, sample.TestDenom)
	eptEul := types.EraUnbondLimit{Denom: sample.TestDenom, Limit: limit}
	require.Equal(t, eptEul, eul)
}

func TestKeeper_SetPoolDetail(t *testing.T) {
	k, ctx := testkeeper.LedgerKeeper(t)

	pool := sample.AccAddress()
	_, found := k.GetPoolDetail(ctx, sample.TestDenom, pool)
	require.False(t, found)

	sa1 := sample.AccAddress()
	sa2 := sample.AccAddress()
	subAccounts := []string{sa1, sa2}
	threshold := uint32(2)
	k.SetPoolDetail(ctx, sample.TestDenom, pool, subAccounts, threshold)

	pd, found := k.GetPoolDetail(ctx, sample.TestDenom, pool)
	require.True(t, found)
	require.Equal(t, types.NewPoolDetail(sample.TestDenom, pool, subAccounts, threshold), pd)
}

func TestKeeper_CurrentEraSnapshots(t *testing.T) {
	k, ctx := testkeeper.LedgerKeeper(t)

	shot := k.CurrentEraSnapshots(ctx, sample.TestDenom)
	require.Equal(t, types.NewEraSnapshot(sample.TestDenom), shot)

	shot1 := types.EraSnapshot{
		Denom:   sample.TestDenom,
		ShotIds: []string{"shotId1", "shotId2"},
	}
	k.SetCurrentEraSnapshot(ctx, shot1)

	shot = k.CurrentEraSnapshots(ctx, sample.TestDenom)
	require.Equal(t, shot1, shot)

	k.ClearCurrentEraSnapshots(ctx, sample.TestDenom)

	shot = k.CurrentEraSnapshots(ctx, sample.TestDenom)
	require.Equal(t, types.NewEraSnapshot(sample.TestDenom), shot)
}

func TestKeeper_SetSnapshot(t *testing.T) {
	k, ctx := testkeeper.LedgerKeeper(t)

	shotId := "testShotId"
	_, found := k.Snapshot(ctx, shotId)
	require.False(t, found)

	bs := types.NewBondSnapshot(sample.TestDenom, sample.AccAddress(), uint32(100),
		types.LinkChunk{Bond: sdk.NewInt(0), Unbond: sdk.NewInt(0), Active: sdk.NewInt(0)})
	k.SetSnapshot(ctx, shotId, bs)

	shot, found := k.Snapshot(ctx, shotId)
	require.True(t, found)
	require.Equal(t, bs, shot)
}

func TestKeeper_SetEraSnapshot(t *testing.T) {
	k, ctx := testkeeper.LedgerKeeper(t)

	era := uint32(100)
	shot := k.EraSnapshot(ctx, sample.TestDenom, era)
	require.Equal(t, types.NewEraSnapshot(sample.TestDenom), shot)

	shot1 := types.EraSnapshot{
		Denom:   sample.TestDenom,
		ShotIds: []string{"shotId1", "shotId2"},
	}
	k.SetEraSnapshot(ctx, era, shot1)

	shot = k.EraSnapshot(ctx, sample.TestDenom, era)
	require.Equal(t, shot1, shot)
}

func TestKeeper_SetChainEra(t *testing.T) {
	k, ctx := testkeeper.LedgerKeeper(t)

	_, found := k.GetChainEra(ctx, sample.TestDenom)
	require.False(t, found)

	era := uint32(100)
	k.SetChainEra(ctx, sample.TestDenom, era)

	ce, found := k.GetChainEra(ctx, sample.TestDenom)
	require.True(t, found)
	require.Equal(t, types.ChainEra{Denom: sample.TestDenom, Era: era}, ce)
}

func TestKeeper_Commission(t *testing.T) {
	k, ctx := testkeeper.LedgerKeeper(t)

	cms := k.GetStakingRewardCommission(ctx, "uratom")
	require.Equal(t, types.DefaultStakingRewardCommission, cms)

	dec, err := utils.NewDecFromStr("0.5")
	require.Nil(t, err)
	t.Log(dec)
	k.SetStakingRewardCommission(ctx, "uratom", dec)

	cms = k.GetStakingRewardCommission(ctx, "uratom")
	require.Equal(t, dec, cms)
}

func TestKeeper_SetTotalExpectedActive(t *testing.T) {
	k, ctx := testkeeper.LedgerKeeper(t)

	era := uint32(100)
	active := k.TotalExpectedActive(ctx, sample.TestDenom, era)
	require.Equal(t, sdk.NewInt(0), active)

	active1 := sdk.NewInt(10000)
	k.SetTotalExpectedActive(ctx, sample.TestDenom, era, active1)

	active = k.TotalExpectedActive(ctx, sample.TestDenom, era)
	require.Equal(t, active1, active)
}

func TestKeeper_GetPoolUnbond(t *testing.T) {
	k, ctx := testkeeper.LedgerKeeper(t)

	pool := sample.AccAddress()
	era := uint32(100)
	_, found := k.GetPoolUnbonding(ctx, sample.TestDenom, pool, era, 1)
	require.False(t, found)

	pu1 := types.NewUnbonding(sample.TestDenom, pool, sdk.NewInt(10))
	k.SetPoolUnbonding(ctx, sample.TestDenom, pool, era, 1, &pu1)

	pu, found := k.GetPoolUnbonding(ctx, sample.TestDenom, pool, era, 1)
	require.True(t, found)
	require.Equal(t, pu1, *pu)
}

func TestKeeper_SetUnbondFee(t *testing.T) {
	k, ctx := testkeeper.LedgerKeeper(t)

	retFee := k.GetUnbondRelayFee(ctx, sample.TestDenom)
	require.Equal(t, types.UnbondRelayFee{
		Denom: sample.TestDenom,
		Value: types.DefaultUnbondRelayFee,
	}, retFee)

	fee := sdk.NewCoin(sample.TestDenom, sdk.NewInt(1))

	k.SetUnbondRelayFee(ctx, sample.TestDenom, fee)

	uf := k.GetUnbondRelayFee(ctx, sample.TestDenom)
	require.Equal(t, fee.String(), uf.Value.String())
}

func TestKeeper_SetUnbondCommission(t *testing.T) {
	k, ctx := testkeeper.LedgerKeeper(t)

	cms := k.GetUnbondCommission(ctx, sample.TestDenom)
	require.Equal(t, types.DefaultUnbondCommission, cms)

	dec, err := utils.NewDecFromStr("0.5")
	require.Nil(t, err)
	k.SetUnbondCommission(ctx, sample.TestDenom, dec)

	cms = k.GetUnbondCommission(ctx, sample.TestDenom)
	require.Equal(t, dec.String(), cms.String())
}

func TestKeeper_SetSignature(t *testing.T) {
	k, ctx := testkeeper.LedgerKeeper(t)

	sig1 := types.NewSignature(sample.TestDenom, uint32(100), sample.AccAddress(), types.TxTypeBond, "testPropId")
	_, found := k.GetSignature(ctx, sig1.Denom, sig1.Era, sig1.Pool, sig1.TxType, sig1.PropId)
	require.False(t, found)

	k.SetSignature(ctx, sig1)

	sig, found := k.GetSignature(ctx, sig1.Denom, sig1.Era, sig1.Pool, sig1.TxType, sig1.PropId)
	require.True(t, found)
	require.Equal(t, sig1.String(), sig.String())
}
