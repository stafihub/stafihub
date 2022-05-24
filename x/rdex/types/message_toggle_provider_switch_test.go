package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stafihub/stafihub/testutil/sample"
	"github.com/stafihub/stafihub/x/rdex/types"
	"github.com/stretchr/testify/require"
)

func TestMsgToggleProviderSwitch_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgToggleProviderSwitch
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgToggleProviderSwitch{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgToggleProviderSwitch{
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
