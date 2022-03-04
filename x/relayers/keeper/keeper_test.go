package keeper_test

import (
	"testing"

	//sdk "github.com/cosmos/cosmos-sdk/types"
	testkeeper "github.com/stafihub/stafihub/testutil/keeper"
	"github.com/stafihub/stafihub/testutil/sample"
	"github.com/stafihub/stafihub/x/relayers/types"
	"github.com/stretchr/testify/require"
)

func Test_LastVoter(t *testing.T) {
	k, ctx := testkeeper.RelayersKeeper(t)

	_, ok := k.LastVoter(ctx, sample.TestLedgerArena, sample.TestDenom)
	require.False(t, ok)

	addr := sample.AccAddress()
	k.SetLastVoter(ctx, sample.TestLedgerArena, sample.TestDenom, addr)

	lv, ok := k.LastVoter(ctx, sample.TestLedgerArena, sample.TestDenom)
	require.True(t, ok)
	require.Equal(t, addr, lv.Voter)
}

func Test_Relayer(t *testing.T) {
	k, ctx := testkeeper.RelayersKeeper(t)

	_, ok := k.GetRelayer(ctx, sample.TestLedgerArena, sample.TestDenom)
	require.False(t, ok)

	addr := sample.AccAddress()
	require.False(t, k.HasRelayer(ctx, sample.TestLedgerArena, sample.TestDenom, addr))
	k.AddRelayer(ctx, sample.TestLedgerArena, sample.TestDenom, addr)
	require.True(t, k.HasRelayer(ctx, sample.TestLedgerArena, sample.TestDenom, addr))

	addr1 := sample.AccAddress()
	require.False(t, k.HasRelayer(ctx, sample.TestLedgerArena, sample.TestDenom, addr1))
	k.AddRelayer(ctx, sample.TestLedgerArena, sample.TestDenom, addr1)
	require.True(t, k.HasRelayer(ctx, sample.TestLedgerArena, sample.TestDenom, addr1))

	rel, ok := k.GetRelayer(ctx, sample.TestLedgerArena, sample.TestDenom)
	require.True(t, ok)
	t.Log(rel)

	k.RemoveRelayer(ctx, sample.TestLedgerArena, sample.TestDenom, addr)
	require.False(t, k.HasRelayer(ctx, sample.TestLedgerArena, sample.TestDenom, addr))
	k.RemoveRelayer(ctx, sample.TestLedgerArena, sample.TestDenom, addr1)
	require.False(t, k.HasRelayer(ctx, sample.TestLedgerArena, sample.TestDenom, addr1))

	rel, ok = k.GetRelayer(ctx, sample.TestLedgerArena, sample.TestDenom)
	require.True(t, ok)
	t.Log(rel)
}

func Test_AllRelayer(t *testing.T) {
	k, ctx := testkeeper.RelayersKeeper(t)
	rels := k.GetAllRelayer(ctx)
	require.True(t, len(rels) == 0)

	addr := sample.AccAddress()
	k.AddRelayer(ctx, sample.TestLedgerArena, sample.TestDenom, addr)
	rels = k.GetAllRelayer(ctx)
	require.True(t, len(rels) == 1)
}

func Test_Threshold(t *testing.T) {
	k, ctx := testkeeper.RelayersKeeper(t)

	_, ok := k.GetThreshold(ctx, sample.TestLedgerArena, sample.TestDenom)
	require.False(t, ok)

	th1 := types.Threshold{Arena: sample.TestLedgerArena, Denom: sample.TestDenom, Value: 3}
	k.SetThreshold(ctx, th1)
	th, ok := k.GetThreshold(ctx, sample.TestLedgerArena, sample.TestDenom)
	require.True(t, ok)
	require.Equal(t, th1, th)

	th2 := types.Threshold{Arena: sample.TestLedgerArena, Denom: sample.TestDenom, Value: 5}
	k.SetThreshold(ctx, th2)
	th, _ = k.GetThreshold(ctx, sample.TestLedgerArena, sample.TestDenom)
	require.True(t, ok)
	require.Equal(t, th2, th)
}

func Test_AllThreshold(t *testing.T) {
	k, ctx := testkeeper.RelayersKeeper(t)
	ths := k.GetAllThreshold(ctx)
	require.True(t, len(ths) == 0)

	th1 := types.Threshold{Arena: sample.TestLedgerArena, Denom: sample.TestDenom, Value: 3}
	k.SetThreshold(ctx, th1)

	ths = k.GetAllThreshold(ctx)
	require.True(t, len(ths) == 1)
	t.Log(ths)
}
