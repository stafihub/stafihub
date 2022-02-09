package keeper_test

import (
	"testing"

	testkeeper "github.com/stafiprotocol/stafihub/testutil/keeper"
	"github.com/stafiprotocol/stafihub/testutil/sample"
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
