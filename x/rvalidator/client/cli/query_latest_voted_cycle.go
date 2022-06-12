package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/rvalidator/types"
)

var _ = strconv.Itoa(0)

func CmdLatestVotedCycle() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "latest-voted-cycle [denom]",
		Short: "Query latest voted cycle",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDenom := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryLatestVotedCycleRequest{

				Denom: reqDenom,
			}

			res, err := queryClient.LatestVotedCycle(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
