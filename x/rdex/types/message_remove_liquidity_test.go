package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stafihub/stafihub/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgRemoveLiquidity_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgRemoveLiquidity
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgRemoveLiquidity{
				Creator:      "invalid_address",
				RmUnit:       sdk.NewInt(20),
				SwapUnit:     sdk.NewInt(0),
				MinOutToken0: sdk.NewCoin(sample.TestDenom, sdk.NewInt(1)),
				MinOutToken1: sdk.NewCoin(sample.TestDenom1, sdk.NewInt(2)),
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgRemoveLiquidity{
				Creator:      sample.AccAddress(),
				RmUnit:       sdk.NewInt(20),
				SwapUnit:     sdk.NewInt(10),
				MinOutToken0: sdk.NewCoin(sample.TestDenom, sdk.NewInt(1)),
				MinOutToken1: sdk.NewCoin(sample.TestDenom1, sdk.NewInt(2)),
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
