package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stafihub/stafihub/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgAddReward_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgAddReward
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgAddReward{
				Creator:         "invalid_address",
				StakePoolIndex:  0,
				RewardPoolIndex: 0,
				AddAmount:       sdk.NewInt(10),
				StartTimestamp:  0,
				RewardPerSecond: sdk.NewInt(0),
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgAddReward{
				Creator:         sample.AccAddress(),
				StakePoolIndex:  0,
				RewardPoolIndex: 0,
				AddAmount:       sdk.NewInt(10),
				StartTimestamp:  0,
				RewardPerSecond: sdk.NewInt(0),
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
