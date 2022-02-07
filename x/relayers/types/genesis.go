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
		Relayers:   []Relayer{},
		Thresholds: []Threshold{},
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in relayer
	relayerDenomMap := make(map[string]struct{})

	for _, elem := range gs.Relayers {
		if sdk.ValidateDenom(elem.Denom) != nil {
			return fmt.Errorf("invalid denom %s", elem.Denom)
		}

		relayerDenomMap[elem.Denom] = struct{}{}
	}
	// Check for duplicated index in threshold
	thresholdDenomMap := make(map[string]struct{})

	for _, elem := range gs.Thresholds {
		if sdk.ValidateDenom(elem.Denom) != nil {
			return fmt.Errorf("invalid denom %s", elem.Denom)
		}

		if _, ok := thresholdDenomMap[elem.Denom]; ok {
			return fmt.Errorf("duplicated denom %s for threshold", elem.Denom)
		}
		thresholdDenomMap[elem.Denom] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}
