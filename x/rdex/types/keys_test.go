package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rdex/types"
	"github.com/stretchr/testify/require"
)

func TestGetLpTokenDenom(t *testing.T) {
	coinsA := sdk.NewCoins(sdk.NewCoin("ufis", sdk.NewInt(5)), sdk.NewCoin("ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2", sdk.NewInt(7)))
	coinsB := sdk.NewCoins(sdk.NewCoin("ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2", sdk.NewInt(7)), sdk.NewCoin("ufis", sdk.NewInt(5)))

	require.Equal(t, types.GetLpTokenDenom(coinsA), types.GetLpTokenDenom(coinsB))
	require.Equal(t, types.GetLpTokenDenom(coinsA), "rdexlp/30bc849e631c1b8230cb9ee822fa6848235fdef5653205c298883864592b9079")
	require.Equal(t, coinsA.String(), "7ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2,5ufis")

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
