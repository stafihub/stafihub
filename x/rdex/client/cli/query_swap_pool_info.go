package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/rdex/types"
)

var _ = strconv.Itoa(0)

func CmdSwapPoolInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swap-pool-info [lp-denom]",
		Short: "Query swap pool",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqLpDenom := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QuerySwapPoolInfoRequest{

				LpDenom: reqLpDenom,
			}

			res, err := queryClient.SwapPoolInfo(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
