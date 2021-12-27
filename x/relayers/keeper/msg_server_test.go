package keeper_test

import (
	"context"
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/stafiprotocol/stafihub/x/relayers/types"
    "github.com/stafiprotocol/stafihub/x/relayers/keeper"
    keepertest "github.com/stafiprotocol/stafihub/testutil/keeper"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.RelayersKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestRelayerMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.RelayersKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateRelayer{Creator: creator,
			Index: strconv.Itoa(i),

		}
		_, err := srv.CreateRelayer(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetRelayer(ctx,
			expected.Index,

		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestRelayerMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteRelayer
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteRelayer{Creator: creator,
				Index: strconv.Itoa(0),

			},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteRelayer{Creator: "B",
				Index: strconv.Itoa(0),

			},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteRelayer{Creator: creator,
				Index: strconv.Itoa(100000),

			},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.RelayersKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateRelayer(wctx, &types.MsgCreateRelayer{Creator: creator,
				Index: strconv.Itoa(0),

			})
			require.NoError(t, err)
			_, err = srv.DeleteRelayer(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetRelayer(ctx,
					tc.request.Index,

				)
				require.False(t, found)
			}
		})
	}
}

func TestThresholdMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateThreshold
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateThreshold{Creator: creator,
				Index: strconv.Itoa(0),

			},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateThreshold{Creator: "B",
				Index: strconv.Itoa(0),

			},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgUpdateThreshold{Creator: creator,
				Index: strconv.Itoa(100000),

			},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.RelayersKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateThreshold{Creator: creator,
				Index: strconv.Itoa(0),

			}
			_, err := srv.CreateThreshold(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateThreshold(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetThreshold(ctx,
					expected.Index,

				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}
