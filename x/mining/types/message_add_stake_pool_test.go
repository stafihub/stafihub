package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stafihub/stafihub/testutil/sample"
	"github.com/stafihub/stafihub/utils"
	"github.com/stafihub/stafihub/x/mining/types"
	"github.com/stretchr/testify/require"
)

func TestMsgAddStakePool_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgAddStakePool
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgAddStakePool{
				Creator:         "invalid_address",
				StakeTokenDenom: sample.TestDenom,
				RewardPoolInfoList: []*types.CreateRewardPoolInfo{
					{RewardTokenDenom: sample.TestDenom1,
						TotalRewardAmount: sdk.NewInt(2),
						RewardPerSecond:   sdk.NewInt(2),
					},
				},
				StakeItemInfoList: []*types.CreateStakeItemInfo{
					{
						LockSecond:      100,
						PowerRewardRate: utils.MustNewDecFromStr("1.5"),
					},
				},
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgAddStakePool{
				Creator:         sample.AccAddress(),
				StakeTokenDenom: sample.TestDenom,
				RewardPoolInfoList: []*types.CreateRewardPoolInfo{
					{
						RewardTokenDenom:  sample.TestDenom1,
						TotalRewardAmount: sdk.NewInt(1),
						RewardPerSecond:   sdk.NewInt(2),
					},
				},
				StakeItemInfoList: []*types.CreateStakeItemInfo{
					{
						LockSecond:      100,
						PowerRewardRate: utils.MustNewDecFromStr("1.5"),
					},
				},
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
