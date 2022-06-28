package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/rvalidator/types"
)

var _ = strconv.Itoa(0)

func CmdLatestDealedCycle() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "latest-dealed-cycle [denom] [pool-address]",
		Short: "Query latest dealed cycle",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDenom := args[0]
			reqPoolAddress := args[1]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryLatestDealedCycleRequest{

				Denom:       reqDenom,
				PoolAddress: reqPoolAddress,
			}

			res, err := queryClient.LatestDealedCycle(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
