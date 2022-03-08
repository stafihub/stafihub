package crisis

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/cosmos/cosmos-sdk/x/crisis/types"
	"github.com/stafihub/stafihub/utils"
)

var (
	_ module.AppModuleBasic = AppModuleBasic{}
)

// AppModuleBasic defines the basic application module used by the distribution module.
type AppModuleBasic struct {
	crisis.AppModuleBasic
}

// DefaultGenesis returns default genesis state as raw bytes for the gov
// module.
func (am AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	defaultGenesisState := types.DefaultGenesisState()
	defaultGenesisState.ConstantFee.Denom = utils.FisDenom

	return cdc.MustMarshalJSON(defaultGenesisState)
}
