package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stafihub/stafihub/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgAddStakePool_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgAddStakePool
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgAddStakePool{
				Creator:           "invalid_address",
				StakeTokenDenom:   sample.TestDenom,
				RewardTokenDenom:  sample.TestDenom1,
				TotalRewardAmount: sdk.NewInt(2),
				RewardPerSecond:   sdk.NewInt(2),
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgAddStakePool{
				Creator:           sample.AccAddress(),
				StakeTokenDenom:   sample.TestDenom,
				RewardTokenDenom:  sample.TestDenom1,
				TotalRewardAmount: sdk.NewInt(0),
				RewardPerSecond:   sdk.NewInt(2),
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
