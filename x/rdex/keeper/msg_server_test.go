package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/stafihub/stafihub/testutil/keeper"
	"github.com/stafihub/stafihub/testutil/sample"
	"github.com/stafihub/stafihub/x/rdex/keeper"
	"github.com/stafihub/stafihub/x/rdex/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context, sdk.Context) {

	k, ctx := keepertest.RdexKeeper(t)
	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx), ctx
}

func TestMsgServerCreatePool(t *testing.T) {
	srv, rdexKeeper, ctx, sdkCtx := setupMsgServer(t)

	creator := sample.TestAdmin

	token0 := sdk.NewCoin(sample.TestDenom, sdk.NewInt(30))
	token1 := sdk.NewCoin(sample.TestDenom1, sdk.NewInt(10))
	lpDenom := types.GetLpTokenDenom(sdk.Coins{token0, token1})
	willMintLpToken := sdk.NewCoin(lpDenom, token0.Amount)

	msgCreatePool := types.MsgCreatePool{
		Creator: creator,
		Token0:  token0,
		Token1:  token1,
	}

	_, err := srv.CreatePool(ctx, &msgCreatePool)
	require.NoError(t, err)

	swapPool, found := rdexKeeper.GetSwapPool(sdkCtx, lpDenom)
	require.True(t, found)

	require.Equal(t, swapPool.LpToken, willMintLpToken)
	require.Equal(t, swapPool.BaseToken, token0)
	require.Equal(t, swapPool.Token, token1)

	lpBalance := keepertest.BankKeeper.GetBalance(sdkCtx, sample.TestAdminAcc, lpDenom)
	require.Equal(t, lpBalance, swapPool.LpToken)

}
