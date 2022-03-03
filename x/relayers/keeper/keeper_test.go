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

	_, ok := k.GetRelayer(ctx, types.ModuleName, sample.TestDenom)
	require.False(t, ok)

	addr := sample.AccAddress()
	require.False(t, k.HasRelayer(ctx, types.ModuleName, sample.TestDenom, addr))
	k.AddRelayer(ctx, types.ModuleName, sample.TestDenom, addr)
	require.True(t, k.HasRelayer(ctx, types.ModuleName, sample.TestDenom, addr))

	addr1 := sample.AccAddress()
	require.False(t, k.HasRelayer(ctx, types.ModuleName, sample.TestDenom, addr1))
	k.AddRelayer(ctx, types.ModuleName, sample.TestDenom, addr1)
	require.True(t, k.HasRelayer(ctx, types.ModuleName, sample.TestDenom, addr1))

	rel, ok := k.GetRelayer(ctx, types.ModuleName, sample.TestDenom)
	require.True(t, ok)
	t.Log(rel)

	k.RemoveRelayer(ctx, types.ModuleName, sample.TestDenom, addr)
	require.False(t, k.HasRelayer(ctx, types.ModuleName, sample.TestDenom, addr))
	k.RemoveRelayer(ctx, types.ModuleName, sample.TestDenom, addr1)
	require.False(t, k.HasRelayer(ctx, types.ModuleName, sample.TestDenom, addr1))

	rel, ok = k.GetRelayer(ctx, types.ModuleName, sample.TestDenom)
	require.True(t, ok)
	t.Log(rel)
}

func Test_AllRelayer(t *testing.T) {
	k, ctx := testkeeper.RelayersKeeper(t)
	rels := k.GetRelayersByTaipeAndDenom(ctx, types.ModuleName, sample.TestDenom)
	require.True(t, len(rels) == 0)

	addr := sample.AccAddress()
	k.AddRelayer(ctx, types.ModuleName, sample.TestDenom, addr)
	rels = k.GetRelayersByTaipeAndDenom(ctx, types.ModuleName, sample.TestDenom)
	require.True(t, len(rels) == 1)
}

func Test_Threshold(t *testing.T) {
	k, ctx := testkeeper.RelayersKeeper(t)

	_, ok := k.GetThreshold(ctx, types.ModuleName, sample.TestDenom)
	require.False(t, ok)

	k.SetThreshold(ctx, types.ModuleName, sample.TestDenom, 3)
	th, ok := k.GetThreshold(ctx, types.ModuleName, sample.TestDenom)
	require.True(t, ok)
	require.Equal(t, 3, th)

	k.SetThreshold(ctx, types.ModuleName, sample.TestDenom, 5)
	th, _ = k.GetThreshold(ctx, types.ModuleName, sample.TestDenom)
	require.True(t, ok)
	require.Equal(t, 5, th)
}
