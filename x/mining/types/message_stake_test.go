package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stafihub/stafihub/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgStake_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgStake
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgStake{
				Creator:    "invalid_address",
				StakeToken: sdk.NewCoin(sample.TestDenom, sdk.NewInt(1)),
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgStake{
				Creator:    sample.AccAddress(),
				StakeToken: sdk.NewCoin(sample.TestDenom, sdk.NewInt(1)),
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
