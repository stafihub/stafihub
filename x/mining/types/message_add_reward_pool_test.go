package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stafihub/stafihub/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgAddRewardPool_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgAddRewardPool
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgAddRewardPool{
				Creator:           "invalid_address",
				RewardTokenDenom:  sample.TestDenom,
				TotalRewardAmount: sdk.NewInt(10),
				RewardPerSecond:   sdk.NewInt(1),
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgAddRewardPool{
				Creator:           sample.AccAddress(),
				RewardTokenDenom:  sample.TestDenom,
				TotalRewardAmount: sdk.NewInt(5),
				RewardPerSecond:   sdk.NewInt(1),
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
