package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	testkeeper "github.com/stafihub/stafihub/testutil/keeper"
	"github.com/stafihub/stafihub/testutil/sample"
	"github.com/stafihub/stafihub/utils"
	"github.com/stretchr/testify/require"
)

func Test_ExchangeRate(t *testing.T) {
	k, ctx := testkeeper.LedgerKeeper(t)

	rates := k.GetAllExchangeRate(ctx)
	require.True(t, len(rates) == 0)

	_, found := k.GetExchangeRate(ctx, sample.TestDenom)
	require.False(t, found)
	rtotal := sdk.NewInt(1000)
	total := sdk.NewInt(1200)
	k.SetExchangeRate(ctx, sample.TestDenom, total, rtotal)

	rate, found := k.GetExchangeRate(ctx, sample.TestDenom)
	require.True(t, found)
	require.Equal(t, utils.OneDec().MulInt(total).QuoInt(rtotal), rate.Value)

	rates = k.GetAllExchangeRate(ctx)
	require.True(t, len(rates) == 1)

	k.SetExchangeRate(ctx, sample.TestDenom1, total, rtotal)
	rates = k.GetAllExchangeRate(ctx)
	require.True(t, len(rates) == 2)
	t.Log(rates)
}

func Test_EraExchangeRate(t *testing.T) {
	k, ctx := testkeeper.LedgerKeeper(t)

	rates := k.GetEraExchangeRateByDenom(ctx, sample.TestDenom)
	require.True(t, len(rates) == 0)

	_, found := k.GetEraExchangeRate(ctx, sample.TestDenom, 1)
	require.False(t, found)

	k.SetEraExchangeRate(ctx, sample.TestDenom, 1, utils.OneDec())
	rate, found := k.GetEraExchangeRate(ctx, sample.TestDenom, 1)
	require.True(t, found)
	require.Equal(t, sample.TestDenom, rate.Denom)
	require.Equal(t, utils.OneDec(), rate.Value)

	k.SetEraExchangeRate(ctx, sample.TestDenom, 2, utils.OneDec())

	rates = k.GetEraExchangeRateByDenom(ctx, sample.TestDenom)
	require.True(t, len(rates) == 2)
	t.Log(rates)
}

func TestKeeper_TokenToRtoken(t *testing.T) {
	k, ctx := testkeeper.LedgerKeeper(t)

	token := sdk.NewInt(100)
	rtoken := k.TokenToRtoken(ctx, sample.TestDenom, token)
	require.Equal(t, token, rtoken)

	total := sdk.NewInt(1000)
	rtotal := sdk.NewInt(1200)
	k.SetExchangeRate(ctx, sample.TestDenom, total, rtotal)

	rtoken = k.TokenToRtoken(ctx, sample.TestDenom, token)
	require.Equal(t, sdk.NewInt(120), rtoken)
}

func TestKeeper_RtokenToToken(t *testing.T) {
	k, ctx := testkeeper.LedgerKeeper(t)

	rtoken := sdk.NewInt(150)
	token := k.RtokenToToken(ctx, sample.TestDenom, rtoken)
	require.Equal(t, rtoken, token)

	rtotal := sdk.NewInt(1000)
	total := sdk.NewInt(1200)
	k.SetExchangeRate(ctx, sample.TestDenom, total, rtotal)

	token = k.TokenToRtoken(ctx, sample.TestDenom, token)
	require.Equal(t, sdk.NewInt(125).String(), token.String())

	rtoken = sdk.NewInt(100)
	token = k.RtokenToToken(ctx, sample.TestDenom, rtoken)
	require.Equal(t, sdk.NewInt(120), token)
}
