package bank

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stafihub/stafihub/utils"
)

var (
	_ module.AppModuleBasic = AppModuleBasic{}
)

// AppModuleBasic defines the basic application module used by the staking module.
type AppModuleBasic struct {
	bank.AppModuleBasic
}

// DefaultGenesis returns default genesis state as raw bytes for the gov
// module.
func (am AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	defaultGenesisState := types.DefaultGenesisState()
	// {
	//     "description": "The native staking token of the StaFi Hub",
	//     "denom_units": [
	//         {
	//             "denom": "ufis",
	//             "exponent": 0,
	//             "aliases": [
	//                 "microfis"
	//             ]
	//         },
	//         {
	//             "denom": "mfis",
	//             "exponent": 3,
	//             "aliases": [
	//               "millifis"
	//             ]
	//         },
	//         {
	//             "denom": "fis",
	//             "exponent": 6
	//         }
	//     ],
	//     "base": "ufis",
	//     "display": "fis",
	//     "name": "FIS",
	//     "symbol": "FIS"
	// }
	defaultGenesisState.DenomMetadata = append(defaultGenesisState.DenomMetadata,
		types.Metadata{
			Description: "The native staking token of the StaFi Hub",
			DenomUnits: []*types.DenomUnit{
				{
					Denom:    utils.FisDenom,
					Exponent: 0,
					Aliases:  []string{"microfis"},
				},
				{
					Denom:    "mfis",
					Exponent: 3,
					Aliases:  []string{"millifis"},
				},
				{
					Denom:    "fis",
					Exponent: 6,
				},
			},
			Base:    utils.FisDenom,
			Display: "fis",
			Name:    "FIS",
			Symbol:  "FIS",
		},
	)
	return cdc.MustMarshalJSON(defaultGenesisState)
}
