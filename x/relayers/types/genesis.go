package types

import (
"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
	    Relayers: []*Relayer{},
		Thresholds: []*Threshold{},
// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
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
