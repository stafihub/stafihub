package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stafihub/stafihub/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgRmStakeToken_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgRmStakeToken
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgRmStakeToken{
				Creator: "invalid_address",
				Denom:   "testDenom",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgRmStakeToken{
				Creator: sample.AccAddress(),
				Denom:   "testDenom",
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
