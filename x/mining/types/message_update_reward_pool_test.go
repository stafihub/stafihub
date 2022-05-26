package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stafihub/stafihub/testutil/sample"
	"github.com/stafihub/stafihub/x/mining/types"
	"github.com/stretchr/testify/require"
)

func TestMsgUpdateRewardPool_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgUpdateRewardPool
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgUpdateRewardPool{
				Creator:         "invalid_address",
				RewardPerSecond: sdk.NewInt(2),
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgUpdateRewardPool{
				Creator:         sample.AccAddress(),
				RewardPerSecond: sdk.NewInt(2),
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
