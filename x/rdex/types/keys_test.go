package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rdex/types"
	"github.com/stretchr/testify/require"
)

func TestGetLpTokenDenom(t *testing.T) {
	require.NotEqual(t, types.GetLpTokenDenom(0), types.GetLpTokenDenom(1))
	require.Equal(t, types.GetLpTokenDenom(0), "rdexlp/0")

}

func TestParseCoins(t *testing.T) {
	coind, err := sdk.ParseCoinsNormalized("30ufis,0uratom")
	require.NoError(t, err)
	require.Equal(t, coind.String(), "30ufis")

	coins := sdk.Coins{sdk.NewCoin("ufis", sdk.NewInt(30)), sdk.NewCoin("uratom", sdk.NewInt(0))}
	require.False(t, coins.IsValid())
}

func TestAddCoins(t *testing.T) {
	coinA := sdk.NewCoin("ufis", sdk.NewInt(30))
	coinB := sdk.NewCoin("uratom", sdk.NewInt(1))

	coins := sdk.NewCoins()
	coins = coins.Add(coinA)
	coins = coins.Add(coinB)
	t.Log(coins.String())
}
