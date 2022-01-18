package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stafiprotocol/stafihub/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateRelayer_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateRelayer
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateRelayer{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateRelayer{
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

func TestMsgDeleteRelayer_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteRelayer
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteRelayer{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteRelayer{
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

func TestMsgUpdateThreshold_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateThreshold
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateThreshold{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateThreshold{
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
