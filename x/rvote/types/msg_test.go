package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stafiprotocol/stafihub/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgSetProposalLife_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSetProposalLife
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetProposalLife{
				Creator: "invalid_address",

			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSetProposalLife{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgSubmitProposal_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSubmitProposal
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSubmitProposal{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSubmitProposal{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
