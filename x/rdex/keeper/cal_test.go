package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	rdexKeeper "github.com/stafihub/stafihub/x/rdex/keeper"
	"github.com/stretchr/testify/require"
)

func TestCalPoolUnit(t *testing.T) {
	testcases := []struct {
		name                 string
		oldPoolUnits         sdk.Int
		nativeAssetBalance   sdk.Int
		externalAssetBalance sdk.Int
		nativeAssetAmount    sdk.Int
		externalAssetAmount  sdk.Int
		poolUnits            sdk.Int
		lpunits              sdk.Int
		panicErr             string
	}{
		{
			name:                 "first-add-zero-native",
			oldPoolUnits:         sdk.ZeroInt(),
			nativeAssetBalance:   sdk.ZeroInt(),
			externalAssetBalance: sdk.ZeroInt(),
			nativeAssetAmount:    sdk.ZeroInt(),
			externalAssetAmount:  sdk.NewInt((100e8)),
			poolUnits:            sdk.ZeroInt(),
			lpunits:              sdk.ZeroInt(),
		},
		{
			name:                 "first-add-zero-external",
			oldPoolUnits:         sdk.ZeroInt(),
			nativeAssetBalance:   sdk.ZeroInt(),
			externalAssetBalance: sdk.ZeroInt(),
			nativeAssetAmount:    sdk.NewInt(100e8),
			externalAssetAmount:  sdk.ZeroInt(),
			poolUnits:            sdk.ZeroInt(),
			lpunits:              sdk.ZeroInt(),
		},
		{
			name:                 "first-add",
			oldPoolUnits:         sdk.ZeroInt(),
			nativeAssetBalance:   sdk.ZeroInt(),
			externalAssetBalance: sdk.ZeroInt(),
			nativeAssetAmount:    sdk.NewInt(100e8),
			externalAssetAmount:  sdk.NewInt(100e8),
			poolUnits:            sdk.NewInt(100e8),
			lpunits:              sdk.NewInt(100e8),
		},
		{
			name:                 "second-add",
			oldPoolUnits:         sdk.NewInt(500e8),
			nativeAssetBalance:   sdk.NewInt(500e8),
			externalAssetBalance: sdk.NewInt(500e8),
			nativeAssetAmount:    sdk.NewInt(345e8),
			externalAssetAmount:  sdk.NewInt(234e8),
			poolUnits:            sdk.NewInt(76359469067),
			lpunits:              sdk.NewInt(26359469067),
		},
		{
			name:                 "tx amount too low",
			oldPoolUnits:         sdk.ZeroInt(),
			nativeAssetBalance:   sdk.ZeroInt(),
			externalAssetBalance: sdk.ZeroInt(),
			nativeAssetAmount:    sdk.ZeroInt(),
			externalAssetAmount:  sdk.ZeroInt(),
			poolUnits:            sdk.ZeroInt(),
			lpunits:              sdk.ZeroInt(),
		},
		{
			name:                 "no pool balance",
			oldPoolUnits:         sdk.ZeroInt(),
			nativeAssetBalance:   sdk.ZeroInt(),
			externalAssetBalance: sdk.ZeroInt(),
			nativeAssetAmount:    sdk.OneInt(),
			externalAssetAmount:  sdk.OneInt(),
			poolUnits:            sdk.OneInt(),
			lpunits:              sdk.OneInt(),
		},
		{
			name:                 "no external balance and assets",
			oldPoolUnits:         sdk.ZeroInt(),
			nativeAssetBalance:   sdk.NewInt(100),
			externalAssetBalance: sdk.ZeroInt(),
			nativeAssetAmount:    sdk.OneInt(),
			externalAssetAmount:  sdk.ZeroInt(),
			poolUnits:            sdk.ZeroInt(),
			lpunits:              sdk.ZeroInt(),
		},
		{
			name:                 "as native asset balance zero then returns native asset amount",
			oldPoolUnits:         sdk.ZeroInt(),
			nativeAssetBalance:   sdk.ZeroInt(),
			externalAssetBalance: sdk.NewInt(100),
			nativeAssetAmount:    sdk.OneInt(),
			externalAssetAmount:  sdk.OneInt(),
			poolUnits:            sdk.OneInt(),
			lpunits:              sdk.OneInt(),
		},
		{
			name:                 "successful1",
			oldPoolUnits:         sdk.ZeroInt(),
			nativeAssetBalance:   sdk.NewInt(100),
			externalAssetBalance: sdk.NewInt(100),
			nativeAssetAmount:    sdk.OneInt(),
			externalAssetAmount:  sdk.OneInt(),
			poolUnits:            sdk.ZeroInt(),
			lpunits:              sdk.ZeroInt(),
		},
		{
			name:                 "successful2",
			oldPoolUnits:         sdk.ZeroInt(),
			nativeAssetBalance:   sdk.NewInt(10000),
			externalAssetBalance: sdk.NewInt(100),
			nativeAssetAmount:    sdk.OneInt(),
			externalAssetAmount:  sdk.OneInt(),
			poolUnits:            sdk.ZeroInt(),
			lpunits:              sdk.ZeroInt(),
		},
		{
			name:                 "successful3",
			oldPoolUnits:         sdk.NewInt(0),
			nativeAssetBalance:   sdk.NewInt(0),
			externalAssetBalance: sdk.NewInt(0),
			nativeAssetAmount:    stringToInt("803080648314941877218"),
			externalAssetAmount:  stringToInt("442072129"),
			lpunits:              stringToInt("803080648314941877218"),
			poolUnits:            stringToInt("803080648314941877218"),
		},
		{
			name:                 "successful4",
			oldPoolUnits:         stringToInt("803080648314941877218"),
			nativeAssetBalance:   stringToInt("803080648314941877218"),
			externalAssetBalance: stringToInt("442072129"),
			nativeAssetAmount:    stringToInt("803080648314941877218"),
			externalAssetAmount:  stringToInt("442072129"),
			lpunits:              stringToInt("803080648314941877218"),
			poolUnits:            stringToInt("803080648314941877218").Mul(sdk.NewInt(2)),
		},
		{
			name:                 "successful5",
			oldPoolUnits:         stringToInt("803080648314941877218"),
			nativeAssetBalance:   stringToInt("803080648314941877218"),
			externalAssetBalance: sdk.NewInt(442072129),
			nativeAssetAmount:    stringToInt("803080648314941877218").Mul(sdk.NewInt(2)),
			externalAssetAmount:  sdk.NewInt(442072129 * 2),
			lpunits:              stringToInt("803080648314941877218").Mul(sdk.NewInt(2)),
			poolUnits:            stringToInt("803080648314941877218").Mul(sdk.NewInt(3)),
		},
		//  {
		// 	oldPoolUnits: sdk.NewInt(1),
		// 	nativeAssetBalance: sdk.NewInt(1),
		// 	externalAssetBalance: sdk.NewInt(1),
		// 	nativeAssetAmount: sdk.int,
		// 	externalAssetAmount: u128::max_value(),
		// 	lpunits: u128::max_value(),
		// 	poolUnits: u128::max_value(),
		// },
		{
			name:                 "successful6",
			oldPoolUnits:         sdk.NewInt(2),
			nativeAssetBalance:   sdk.NewInt(2),
			externalAssetBalance: sdk.NewInt(2),
			nativeAssetAmount:    sdk.NewInt(3),
			externalAssetAmount:  sdk.NewInt(3),
			lpunits:              sdk.NewInt(3),
			poolUnits:            sdk.NewInt(5),
		},
		{
			name:                 "successful7",
			oldPoolUnits:         sdk.NewInt(1),
			nativeAssetBalance:   sdk.NewInt(3),
			externalAssetBalance: sdk.NewInt(3),
			nativeAssetAmount:    sdk.NewInt(4),
			externalAssetAmount:  sdk.NewInt(4),
			lpunits:              sdk.NewInt(1),
			poolUnits:            sdk.NewInt(2),
		},
		{
			name:                 "successful8",
			oldPoolUnits:         sdk.NewInt(100),
			nativeAssetBalance:   sdk.NewInt(100),
			externalAssetBalance: sdk.NewInt(1),
			nativeAssetAmount:    sdk.NewInt(0),
			externalAssetAmount:  sdk.NewInt(2),
			lpunits:              sdk.NewInt(33),
			poolUnits:            sdk.NewInt(133),
		},
		{
			name:                 "successful9",
			oldPoolUnits:         sdk.NewInt(100000000000000),
			nativeAssetBalance:   sdk.NewInt(100000000000000),
			externalAssetBalance: sdk.NewInt(10000000000),
			nativeAssetAmount:    sdk.NewInt(0),
			externalAssetAmount:  sdk.NewInt(20000000000),
			lpunits:              sdk.NewInt(33333333333333),
			poolUnits:            sdk.NewInt(133333333333333),
		},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if tc.panicErr != "" {
				// nolint:errcheck
				require.PanicsWithError(t, tc.panicErr, func() {
					rdexKeeper.CalPoolUnit(
						tc.oldPoolUnits,
						tc.nativeAssetBalance,
						tc.externalAssetBalance,
						tc.nativeAssetAmount,
						tc.externalAssetAmount,
					)
				})
				return
			}

			poolUnits, lpunits := rdexKeeper.CalPoolUnit(
				tc.oldPoolUnits,
				tc.nativeAssetBalance,
				tc.externalAssetBalance,
				tc.nativeAssetAmount,
				tc.externalAssetAmount,
			)

			require.Equal(t, tc.poolUnits.String(), poolUnits.String())
			require.Equal(t, tc.lpunits.String(), lpunits.String())
		})
	}
}

func TestCalcSwapResult(t *testing.T) {
	testcases := []struct {
		name            string
		inputIsBase     bool
		Y, X, x, y, fee sdk.Int
	}{
		{
			Y:           sdk.NewInt(0),
			X:           stringToInt("803080648314941877218"),
			x:           stringToInt("803080648314941877218"),
			inputIsBase: true,
			y:           sdk.NewInt(0),
			fee:         sdk.NewInt(0),
		},
		{
			Y:           stringToInt("803080648314941877218"),
			X:           sdk.NewInt(0),
			x:           stringToInt("803080648314941877218"),
			inputIsBase: true,
			y:           sdk.NewInt(0),
			fee:         sdk.NewInt(0),
		},
		{
			Y:           stringToInt("803080648314941877218"),
			X:           stringToInt("803080648314941877218"),
			x:           sdk.NewInt(0),
			inputIsBase: true,
			y:           sdk.NewInt(0),
			fee:         sdk.NewInt(0),
		},
		// {
		// 	Y:           stringToInt("1000000000000000000000000000"),
		// 	X:           stringToInt("1000000000000000000000000000"),
		// 	x:           stringToInt("100000000000000000000000000"),
		// 	inputIsBase: true,
		// 	y:           stringToInt("82644628099173553719008264"),
		// 	fee:         stringToInt("8264462809917355371900826"),
		// },
		{
			Y:           sdk.NewInt(100),
			X:           sdk.NewInt(1000),
			x:           sdk.NewInt(10),
			inputIsBase: true,
			y:           sdk.NewInt(82),
			fee:         sdk.NewInt(8),
		},
		{
			Y:           sdk.NewInt(1000),
			X:           sdk.NewInt(100),
			x:           sdk.NewInt(10),
			inputIsBase: false,
			y:           sdk.NewInt(82),
			fee:         sdk.NewInt(8),
		},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			y, fee := rdexKeeper.CalSwapResult(tc.Y, tc.X, tc.x, tc.inputIsBase)
			require.Equal(t, tc.y, y)
			require.Equal(t, tc.fee, fee)
		})
	}
}

func stringToInt(s string) sdk.Int {
	value, _ := sdk.NewIntFromString(s)
	return value
}

func TestCalRemoveAmount(t *testing.T) {
	testcases := []struct {
		name                    string
		poolUnits               sdk.Int
		rmUnits                 sdk.Int
		swapUnits               sdk.Int
		nativeAssetBalance      sdk.Int
		externalAssetBalance    sdk.Int
		inputIsBase             bool
		expectRmBaseTokenAmount sdk.Int
		expectRmTokenAmount     sdk.Int
		expectSwapAmount        sdk.Int
		panicErr                string
	}{
		{
			poolUnits:               sdk.NewInt(0),
			rmUnits:                 sdk.NewInt(1),
			swapUnits:               sdk.NewInt(0),
			nativeAssetBalance:      sdk.NewInt(0),
			externalAssetBalance:    sdk.NewInt(0),
			inputIsBase:             true,
			expectRmBaseTokenAmount: sdk.NewInt(0),
			expectRmTokenAmount:     sdk.NewInt(0),
			expectSwapAmount:        sdk.NewInt(0),
		},
		{
			poolUnits:               sdk.NewInt(1),
			rmUnits:                 sdk.NewInt(0),
			swapUnits:               sdk.NewInt(0),
			nativeAssetBalance:      sdk.NewInt(1),
			externalAssetBalance:    sdk.NewInt(1),
			inputIsBase:             true,
			expectRmBaseTokenAmount: sdk.NewInt(0),
			expectRmTokenAmount:     sdk.NewInt(0),
			expectSwapAmount:        sdk.NewInt(0),
		},
		{
			poolUnits:               sdk.NewInt(10),
			rmUnits:                 sdk.NewInt(1),
			swapUnits:               sdk.NewInt(0),
			nativeAssetBalance:      sdk.NewInt(20),
			externalAssetBalance:    sdk.NewInt(20),
			inputIsBase:             true,
			expectRmBaseTokenAmount: sdk.NewInt(2),
			expectRmTokenAmount:     sdk.NewInt(2),
			expectSwapAmount:        sdk.NewInt(0),
		},
		{
			poolUnits:               sdk.NewInt(10),
			rmUnits:                 sdk.NewInt(1),
			swapUnits:               sdk.NewInt(1),
			nativeAssetBalance:      sdk.NewInt(20),
			externalAssetBalance:    sdk.NewInt(20),
			inputIsBase:             true,
			expectRmBaseTokenAmount: sdk.NewInt(2),
			expectRmTokenAmount:     sdk.NewInt(2),
			expectSwapAmount:        sdk.NewInt(2),
		},
		// {
		// 	poolUnits: u128::max_value(),
		// 	rmUnits: u128::max_value(),
		// 	swapUnits: 1,
		// 	nativeAssetBalance: u128::max_value(),
		// 	externalAssetBalance: u128::max_value(),
		// 	inputIsBase: true,
		// 	expectRmBaseTokenAmount: u128::max_value(),
		// 	expectRmTokenAmount: u128::max_value(),
		// 	expectSwapAmount: 1,
		// },
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if tc.panicErr != "" {
				require.PanicsWithError(t, tc.panicErr, func() {
					rdexKeeper.CalRemoveAmount(tc.poolUnits, tc.rmUnits, tc.swapUnits, tc.nativeAssetBalance, tc.externalAssetBalance, tc.inputIsBase)
				})
				return
			}

			baseTokenAmount, tokenAMount, swapAmount := rdexKeeper.CalRemoveAmount(tc.poolUnits, tc.rmUnits, tc.swapUnits, tc.nativeAssetBalance, tc.externalAssetBalance, tc.inputIsBase)

			require.Equal(t, tc.expectRmBaseTokenAmount, baseTokenAmount)
			require.Equal(t, tc.expectRmTokenAmount, tokenAMount)
			require.Equal(t, tc.expectSwapAmount, swapAmount)
		})
	}
}
