package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stafiprotocol/stafihub/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgUpdateAdmin_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateAdmin
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateAdmin{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateAdmin{
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

func TestMsgAddDenom_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgAddDenom
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgAddDenom{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgAddDenom{
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
