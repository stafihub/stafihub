package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stafihub/stafihub/testutil/sample"
	"github.com/stafihub/stafihub/x/mining/types"
	"github.com/stretchr/testify/require"
)

func TestMsgAddRewardPool_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgAddRewardPool
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgAddRewardPool{
				Creator:           "invalid_address",
				RewardTokenDenom:  sample.TestDenom,
				TotalRewardAmount: sdk.NewInt(10),
				RewardPerSecond:   sdk.NewInt(1),
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgAddRewardPool{
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
