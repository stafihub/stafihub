package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/stafiprotocol/stafihub/testutil/keeper"
	"github.com/stafiprotocol/stafihub/x/relayers/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestRelayerQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.RelayersKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRelayer(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllRelayerRequest {
		return &types.QueryAllRelayerRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.RelayerAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Relayer), step)
			require.Subset(t, msgs, resp.Relayer)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.RelayerAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Relayer), step)
			require.Subset(t, msgs, resp.Relayer)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.RelayerAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.RelayerAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

func TestThresholdQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.RelayersKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNThreshold(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetThresholdRequest
		response *types.QueryGetThresholdResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetThresholdRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetThresholdResponse{Threshold: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetThresholdRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetThresholdResponse{Threshold: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetThresholdRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Threshold(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}

func TestThresholdQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.RelayersKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNThreshold(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllThresholdRequest {
		return &types.QueryAllThresholdRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ThresholdAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Threshold), step)
			require.Subset(t, msgs, resp.Threshold)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ThresholdAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Threshold), step)
			require.Subset(t, msgs, resp.Threshold)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ThresholdAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ThresholdAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
