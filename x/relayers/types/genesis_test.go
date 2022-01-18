package types_test

import (
	"testing"

	"github.com/stafiprotocol/stafihub/x/relayers/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				Relayers: []types.Relayer{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				Thresholds: []types.Threshold{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				ProposalList: []types.Proposal{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated relayer",
			genState: &types.GenesisState{
				Relayers: []types.Relayer{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated threshold",
			genState: &types.GenesisState{
				Thresholds: []types.Threshold{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated proposal",
			genState: &types.GenesisState{
				ProposalList: []types.Proposal{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
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
