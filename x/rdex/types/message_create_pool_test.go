package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stafihub/stafihub/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreatePool_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreatePool
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreatePool{
				Creator: "invalid_address",
				Token0:  sdk.NewCoin(sample.TestDenom, sdk.NewInt(21)),
				Token1:  sdk.NewCoin(sample.TestDenom1, sdk.NewInt(1)),
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreatePool{
				Creator: sample.AccAddress(),
				Token0:  sdk.NewCoin(sample.TestDenom, sdk.NewInt(21)),
				Token1:  sdk.NewCoin(sample.TestDenom1, sdk.NewInt(1)),
			},
		}, {
			name: "invalid coins",
			msg: MsgCreatePool{
				Creator: sample.AccAddress(),
				Token0:  sdk.NewCoin(sample.TestDenom, sdk.NewInt(0)),
				Token1:  sdk.NewCoin(sample.TestDenom1, sdk.NewInt(1)),
			},
			err: sdkerrors.ErrInvalidCoins,
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
