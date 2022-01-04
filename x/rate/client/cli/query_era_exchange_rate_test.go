package cli_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/stafiprotocol/stafihub/testutil/network"
	"github.com/stafiprotocol/stafihub/x/rate/client/cli"
    "github.com/stafiprotocol/stafihub/x/rate/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func networkWithEraExchangeRateObjects(t *testing.T, n int) (*network.Network, []types.EraExchangeRate) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
    require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := 0; i < n; i++ {
		state.EraExchangeRateList = append(state.EraExchangeRateList, types.EraExchangeRate{
		    Index: strconv.Itoa(i),

		})
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.EraExchangeRateList
}

func TestShowEraExchangeRate(t *testing.T) {
	net, objs := networkWithEraExchangeRateObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc string
		idIndex string

		args []string
		err  error
		obj  types.EraExchangeRate
	}{
		{
			desc: "found",
			idIndex: objs[0].Index,

			args: common,
			obj:  objs[0],
		},
		{
			desc: "not found",
			idIndex: strconv.Itoa(100000),

			args: common,
			err:  status.Error(codes.InvalidArgument, "not found"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
			    tc.idIndex,

			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowEraExchangeRate(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetEraExchangeRateResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.EraExchangeRate)
				require.Equal(t, tc.obj, resp.EraExchangeRate)
			}
		})
	}
}
