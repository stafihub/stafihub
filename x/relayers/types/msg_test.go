package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stafihub/stafihub/testutil/sample"
	"github.com/stafihub/stafihub/x/relayers/types"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateRelayer_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgAddRelayer
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgAddRelayer{
				Creator:   "invalid_address",
				Arena:     types.ModuleName,
				Addresses: []string{sample.AccAddress()},
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgAddRelayer{
				Creator:   sample.AccAddress(),
				Arena:     types.ModuleName,
				Addresses: []string{sample.AccAddress()},
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
		msg  types.MsgDeleteRelayer
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgDeleteRelayer{
				Creator: "invalid_address",
				Arena:   types.ModuleName,
				Address: sample.AccAddress(),
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgDeleteRelayer{
				Creator: sample.AccAddress(),
				Arena:   types.ModuleName,
				Address: sample.AccAddress(),
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
		msg  types.MsgSetThreshold
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgSetThreshold{
				Creator: "invalid_address",
				Arena:   types.ModuleName,
				Denom:   sample.TestDenom,
				Value:   1,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgSetThreshold{
				Creator: sample.AccAddress(),
				Arena:   types.ModuleName,
				Denom:   sample.TestDenom,
				Value:   1,
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
