package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stafihub/stafihub/testutil/sample"
	"github.com/stafihub/stafihub/utils"
	"github.com/stretchr/testify/require"
)

func TestMsgUpdateStakeItem_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateStakeItem
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateStakeItem{
				Creator: "invalid_address",
				PowerRewardRate: utils.MustNewDecFromStr("0.1"),
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateStakeItem{
				Creator: sample.AccAddress(),
				PowerRewardRate: utils.MustNewDecFromStr("3.1"),
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
