package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stafihub/stafihub/testutil/sample"
	"github.com/stafihub/stafihub/utils"
	"github.com/stretchr/testify/require"
)

func TestMsgRemovePool_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgRmBondedPool
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgRmBondedPool{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgRmBondedPool{
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

func TestMsgSetEraUnbondLimit_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSetEraUnbondLimit
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetEraUnbondLimit{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSetEraUnbondLimit{
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

func TestMsgSetInitBond_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSetLeastBond
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetLeastBond{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSetLeastBond{
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

func TestMsgSetChainBondingDuration_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSetRParams
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetRParams{
				Creator:    "invalid_address",
				EraSeconds: 1000,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSetRParams{
				Creator:    sample.AccAddress(),
				EraSeconds: 20000,
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

func TestMsgSetPoolDetail_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSetPoolDetail
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetPoolDetail{
				Creator:     "invalid_address",
				Threshold:   1,
				SubAccounts: []string{sample.AccAddress()},
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSetPoolDetail{
				Creator:     sample.AccAddress(),
				Threshold:   1,
				SubAccounts: []string{sample.AccAddress()},
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

func TestMsgSetLeastBond_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSetLeastBond
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetLeastBond{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSetLeastBond{
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

func TestMsgClearCurrentEraSnapShots_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgClearCurrentEraSnapShots
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgClearCurrentEraSnapShots{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgClearCurrentEraSnapShots{
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

func TestMsgSetCommission_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSetStakingRewardCommission
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetStakingRewardCommission{
				Creator:    "invalid_address",
				Commission: utils.MustNewDecFromStr("0.1"),
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSetStakingRewardCommission{
				Creator:    sample.AccAddress(),
				Commission: utils.MustNewDecFromStr("0.02"),
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

func TestMsgSetReceiver_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSetProtocolFeeReceiver
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetProtocolFeeReceiver{
				Creator:  "invalid_address",
				Receiver: "ff",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSetProtocolFeeReceiver{
				Creator:  sample.AccAddress(),
				Receiver: sample.AccAddress(),
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

func TestMsgSetUnbondCommission_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSetUnbondCommission
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetUnbondCommission{
				Creator:    "invalid_address",
				Commission: utils.MustNewDecFromStr("0.1"),
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSetUnbondCommission{
				Creator:    sample.AccAddress(),
				Commission: utils.MustNewDecFromStr("0"),
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

func TestMsgLiquidityUnbond_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgLiquidityUnbond
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgLiquidityUnbond{
				Creator: "invalid_address",
				Value:   sdk.NewCoin(sample.TestDenom, sdk.NewInt(1)),
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgLiquidityUnbond{
				Creator: sample.AccAddress(),
				Value:   sdk.NewCoin(sample.TestDenom, sdk.NewInt(2)),
			},
		}, {
			name: "valid address",
			msg: MsgLiquidityUnbond{
				Creator: sample.AccAddress(),
				Value:   sdk.NewCoin(sample.TestDenom, sdk.NewInt(0)),
			},
			err: sdkerrors.ErrInvalidCoins,
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

func TestMsgSetUnbondFee_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSetUnbondRelayFee
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetUnbondRelayFee{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSetUnbondRelayFee{
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

func TestMsgSubmitSignature_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSubmitSignature
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSubmitSignature{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSubmitSignature{
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

func TestMsgLiquidityUnbond_GetSignBytes(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32("stafi1wz9ax9xlxjtw9akxyf29aflau4f63p5duvkg9z")
	msg := NewMsgLiquidityUnbond(addr, "cosmos1gsth46z50w256p4kq36xquh4q90mfjq0t4lm9scln6zucg64epyqudzqzm", sdk.Coin{
		Denom:  "uratom",
		Amount: sdk.NewInt(200000),
	}, "cosmos1wz9ax9xlxjtw9akxyf29aflau4f63p5d88xz36")
	res := msg.GetSignBytes()
	t.Log(string(res))
	expected := `{"type":"ledger/LiquidityUnbond","value":{"creator":"stafi1wz9ax9xlxjtw9akxyf29aflau4f63p5duvkg9z","pool":"cosmos1gsth46z50w256p4kq36xquh4q90mfjq0t4lm9scln6zucg64epyqudzqzm","recipient":"cosmos1wz9ax9xlxjtw9akxyf29aflau4f63p5d88xz36","value":{"amount":"200000","denom":"uratom"}}}`
	require.Equal(t, expected, string(res))
}
