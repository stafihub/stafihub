package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stafiprotocol/stafihub/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgAddNewPool_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgAddNewPool
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgAddNewPool{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgAddNewPool{
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

func TestMsgRemovePool_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgRemovePool
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgRemovePool{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgRemovePool{
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
		msg  MsgSetInitBond
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetInitBond{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSetInitBond{
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
		msg  MsgSetChainBondingDuration
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetChainBondingDuration{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSetChainBondingDuration{
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

func TestMsgSetPoolDetail_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSetPoolDetail
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetPoolDetail{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSetPoolDetail{
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
		msg  MsgSetCommission
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetCommission{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSetCommission{
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

func TestMsgSetReceiver_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSetReceiver
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetReceiver{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSetReceiver{
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

func TestMsgSetUnbondCommission_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSetUnbondCommission
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetUnbondCommission{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSetUnbondCommission{
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
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgLiquidityUnbond{
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

func TestMsgSetUnbondFee_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSetUnbondFee
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetUnbondFee{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSetUnbondFee{
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
