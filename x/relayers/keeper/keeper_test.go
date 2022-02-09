package keeper_test

import (
	"testing"

	//sdk "github.com/cosmos/cosmos-sdk/types"
	testkeeper "github.com/stafiprotocol/stafihub/testutil/keeper"
	"github.com/stafiprotocol/stafihub/testutil/sample"
	"github.com/stafiprotocol/stafihub/x/relayers/types"
	"github.com/stretchr/testify/require"
)

func Test_LastVoter(t *testing.T) {
	k, ctx := testkeeper.RelayersKeeper(t)

	_, ok := k.LastVoter(ctx, sample.TestDenom)
	require.False(t, ok)

	addr := sample.AccAddress()
	k.SetLastVoter(ctx, sample.TestDenom, addr)

	lv, ok := k.LastVoter(ctx, sample.TestDenom)
	require.True(t, ok)
	require.Equal(t, addr, lv.Voter)
}

func Test_Relayer(t *testing.T) {
	k, ctx := testkeeper.RelayersKeeper(t)

	_, ok := k.GetRelayerByDenom(ctx, sample.TestDenom)
	require.False(t, ok)

	addr := sample.AccAddress()
	require.False(t, k.IsRelayer(ctx, sample.TestDenom, addr))
	k.AddRelayer(ctx, sample.TestDenom, addr)
	require.True(t, k.IsRelayer(ctx, sample.TestDenom, addr))

	addr1 := sample.AccAddress()
	require.False(t, k.IsRelayer(ctx, sample.TestDenom, addr1))
	k.AddRelayer(ctx, sample.TestDenom, addr1)
	require.True(t, k.IsRelayer(ctx, sample.TestDenom, addr1))

	rel, ok := k.GetRelayerByDenom(ctx, sample.TestDenom)
	require.True(t, ok)
	t.Log(rel)

	k.RemoveRelayer(ctx, sample.TestDenom, addr)
	require.False(t, k.IsRelayer(ctx, sample.TestDenom, addr))
	k.RemoveRelayer(ctx, sample.TestDenom, addr1)
	require.False(t, k.IsRelayer(ctx, sample.TestDenom, addr1))

	rel, ok = k.GetRelayerByDenom(ctx, sample.TestDenom)
	require.True(t, ok)
	t.Log(rel)
}

func Test_AllRelayer(t *testing.T) {
	k, ctx := testkeeper.RelayersKeeper(t)
	rels := k.GetAllRelayer(ctx)
	require.True(t, len(rels) == 0)

	addr := sample.AccAddress()
	k.AddRelayer(ctx, sample.TestDenom, addr)
	rels = k.GetAllRelayer(ctx)
	require.True(t, len(rels) == 1)
}

func Test_Threshold(t *testing.T) {
	k, ctx := testkeeper.RelayersKeeper(t)

	_, ok := k.GetThreshold(ctx, sample.TestDenom)
	require.False(t, ok)

	th1 := types.Threshold{Denom: sample.TestDenom, Value: 3}
	k.SetThreshold(ctx, th1)
	th, ok := k.GetThreshold(ctx, sample.TestDenom)
	require.True(t, ok)
	require.Equal(t, th, th1)

	th2 := types.Threshold{Denom: sample.TestDenom, Value: 5}
	k.SetThreshold(ctx, th2)
	th, _ = k.GetThreshold(ctx, sample.TestDenom)
	require.True(t, ok)
	require.Equal(t, th, th2)
}

func Test_AllThreshold(t *testing.T) {
	k, ctx := testkeeper.RelayersKeeper(t)
	ths := k.GetAllThreshold(ctx)
	require.True(t, len(ths) == 0)

	th1 := types.Threshold{Denom: sample.TestDenom, Value: 3}
	k.SetThreshold(ctx, th1)

	ths = k.GetAllThreshold(ctx)
	require.True(t, len(ths) == 1)
	t.Log(ths)
}
