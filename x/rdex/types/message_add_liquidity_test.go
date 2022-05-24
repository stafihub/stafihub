package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stafihub/stafihub/testutil/sample"
	"github.com/stafihub/stafihub/x/rdex/types"
	"github.com/stretchr/testify/require"
)

func TestMsgAddLiquidity_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgAddLiquidity
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgAddLiquidity{
				Creator: "invalid_address",
				Token0:  sdk.NewCoin(sample.TestDenom, sdk.NewInt(21)),
				Token1:  sdk.NewCoin(sample.TestDenom1, sdk.NewInt(1)),
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgAddLiquidity{
				Creator: sample.AccAddress(),
				Token0:  sdk.NewCoin(sample.TestDenom, sdk.NewInt(10)),
				Token1:  sdk.NewCoin(sample.TestDenom1, sdk.NewInt(0)),
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
