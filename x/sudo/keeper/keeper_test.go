package keeper_test

import (
	"testing"

	testkeeper "github.com/stafihub/stafihub/testutil/keeper"
	"github.com/stafihub/stafihub/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestAdmin(t *testing.T) {
	k, ctx := testkeeper.SudoKeeper(t)

	admin := k.GetAdmin(ctx)
	require.Equal(t, sample.TestAdminAcc, admin)

	addr := sample.OriginAccAddress()
	k.SetAdmin(ctx, addr)
	require.True(t, k.IsAdmin(ctx, addr.String()))

	admin = k.GetAdmin(ctx)
	require.Equal(t, addr.String(), admin.String())

	addr1 := sample.AccAddress()
	require.NotEqual(t, addr.String(), addr1)
	require.False(t, k.IsAdmin(ctx, addr1))
}

func TestKeeper_SetAddressPrefix(t *testing.T) {
	k, ctx := testkeeper.SudoKeeper(t)

	_, found := k.GetAddressPrefix(ctx, sample.TestDenom)
	require.False(t, found)

	k.SetAddressPrefix(ctx, sample.TestDenom, sample.TestAddrPrefix)
	val, found := k.GetAddressPrefix(ctx, sample.TestDenom)
	require.True(t, found)
	require.Equal(t, sample.TestAddrPrefix, val)
}


