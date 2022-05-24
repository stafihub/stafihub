package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stafihub/stafihub/testutil/sample"
	"github.com/stafihub/stafihub/x/rdex/types"
	"github.com/stretchr/testify/require"
)

func TestMsgSwap_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgSwap
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgSwap{
				Creator:     "invalid_address",
				InputToken:  sdk.NewCoin(sample.TestDenom, sdk.NewInt(1)),
				MinOutToken: sdk.NewCoin(sample.TestDenom, sdk.NewInt(1)),
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgSwap{
				Creator:     sample.AccAddress(),
				InputToken:  sdk.NewCoin(sample.TestDenom, sdk.NewInt(1)),
				MinOutToken: sdk.NewCoin(sample.TestDenom, sdk.NewInt(0)),
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
