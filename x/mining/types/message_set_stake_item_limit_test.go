package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stafihub/stafihub/testutil/sample"
	"github.com/stafihub/stafihub/utils"
	"github.com/stretchr/testify/require"
)

func TestMsgSetStakeItemLimit_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSetStakeItemLimit
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetStakeItemLimit{
				Creator:            "invalid_address",
				MaxLockSecond:      1,
				MaxPowerRewardRate: utils.NewDec(1),
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSetStakeItemLimit{
				Creator:            sample.AccAddress(),
				MaxLockSecond:      1,
				MaxPowerRewardRate: utils.NewDec(1),
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
