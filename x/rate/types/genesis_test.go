package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stafiprotocol/stafihub/x/rate/types"
)

func TestGenesisState_Validate(t *testing.T) {
    for _, tc := range []struct {
    		desc          string
    		genState      *types.GenesisState
    		valid bool
    } {
        {
            desc:     "default is valid",
            genState: types.DefaultGenesis(),
            valid:    true,
        },
        {
            desc:     "valid genesis state",
            genState: &types.GenesisState{
                ExchangeRateList: []types.ExchangeRate{
	{
		Index: "0",
},
	{
		Index: "1",
},
},
EraExchangeRateList: []types.EraExchangeRate{
	{
		Index: "0",
},
	{
		Index: "1",
},
},
// this line is used by starport scaffolding # types/genesis/validField
            },
            valid:    true,
        },
        {
	desc:     "duplicated exchangeRate",
	genState: &types.GenesisState{
		ExchangeRateList: []types.ExchangeRate{
			{
				Index: "0",
},
			{
				Index: "0",
},
		},
	},
	valid:    false,
},
{
	desc:     "duplicated eraExchangeRate",
	genState: &types.GenesisState{
		EraExchangeRateList: []types.EraExchangeRate{
			{
				Index: "0",
},
			{
				Index: "0",
},
		},
	},
	valid:    false,
},
// this line is used by starport scaffolding # types/genesis/testcase
    } {
        t.Run(tc.desc, func(t *testing.T) {
            err := tc.genState.Validate()
            if tc.valid {
                require.NoError(t, err)
            } else {
                require.Error(t, err)
            }
        })
    }
}