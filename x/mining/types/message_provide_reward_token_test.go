package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stafihub/stafihub/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgProvideRewardToken_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgProvideRewardToken
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgProvideRewardToken{
				Creator: "invalid_address",
				Token:   sdk.NewCoin(sample.TestDenom, sdk.NewInt(1)),
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgProvideRewardToken{
				Creator: sample.AccAddress(),
				Token:   sdk.NewCoin(sample.TestDenom, sdk.NewInt(0)),
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
