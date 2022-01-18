package types

import (
	// this line is used by starport scaffolding # genesis/types/import
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ProposalLife: 0,
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if gs.ProposalLife < 0 {
		return fmt.Errorf("proposalLife %d is negative", gs.ProposalLife)
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}
