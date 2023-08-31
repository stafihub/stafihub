package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stafihub/stafihub/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgForceRmBondedPool_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgForceRmBondedPool
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgForceRmBondedPool{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgForceRmBondedPool{
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
