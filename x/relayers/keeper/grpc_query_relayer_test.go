package keeper_test

import (
    "strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/stafiprotocol/stafihub/x/relayers/types"
	keepertest "github.com/stafiprotocol/stafihub/testutil/keeper"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestRelayerQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.RelayersKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRelayer(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetRelayerRequest
		response *types.QueryGetRelayerResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetRelayerRequest{
			    Index: msgs[0].Index,
                
			},
			response: &types.QueryGetRelayerResponse{Relayer: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetRelayerRequest{
			    Index: msgs[1].Index,
                
			},
			response: &types.QueryGetRelayerResponse{Relayer: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetRelayerRequest{
			    Index:strconv.Itoa(100000),
                
			},
			err:     status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Relayer(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}

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
