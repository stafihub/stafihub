package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/stafiprotocol/stafihub/testutil/sample"
)

func TestMsgSetChainEra_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSetChainEra
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetChainEra{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSetChainEra{
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

func TestMsgActiveReport_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgActiveReport
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgActiveReport{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgActiveReport{
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

