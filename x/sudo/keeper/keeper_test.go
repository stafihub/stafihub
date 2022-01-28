package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	testkeeper "github.com/stafiprotocol/stafihub/testutil/keeper"
	"github.com/stafiprotocol/stafihub/testutil/sample"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAdmin(t *testing.T) {
	k, ctx := testkeeper.SudoKeeper(t)

	addr := sample.AccAddress()
	t.Log("addr", addr)
	ac, _ := sdk.AccAddressFromBech32(addr)
	k.SetAdmin(ctx, ac)

	require.True(t, k.IsAdmin(ctx, addr))

	admin := k.GetAdmin(ctx)
	require.Equal(t, addr, admin.String())

	addr1 := sample.AccAddress()
	t.Log("addr1", addr1)
	require.NotEqual(t, addr, addr1)
	require.False(t, k.IsAdmin(ctx, addr1))
}
