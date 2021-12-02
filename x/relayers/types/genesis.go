package types

import (
"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Admin: "fis1xrhhus43kjhqdccee7aqnuukh8ugv09affsg84",
	    Relayers: []Relayer{},
		Thresholds: []Threshold{},
// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	_, err := sdk.AccAddressFromBech32(gs.Admin)
	if err != nil {
		return fmt.Errorf("invalid admin address %s for relayer", gs.Admin)
	}

    // Check for duplicated index in relayer
relayerDenomMap := make(map[string]struct{})

for _, elem := range gs.Relayers {
	key := elem.Denom+elem.Address
	if _, ok := relayerDenomMap[key]; ok {
		return fmt.Errorf("duplicated denom %s and address %s for relayer", elem.Denom, elem.Address)
	}
	relayerDenomMap[key] = struct{}{}
}
// Check for duplicated index in threshold
thresholdDenomMap := make(map[string]struct{})

for _, elem := range gs.Thresholds {
	if _, ok := thresholdDenomMap[elem.Denom]; ok {
		return fmt.Errorf("duplicated denom %s for threshold", elem.Denom)
	}
	thresholdDenomMap[elem.Denom] = struct{}{}
}
// this line is used by starport scaffolding # genesis/types/validate

	return nil
}
