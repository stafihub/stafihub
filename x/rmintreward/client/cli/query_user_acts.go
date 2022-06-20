package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/rmintreward/types"
)

var _ = strconv.Itoa(0)

func CmdUserActs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user-acts [address] [denom]",
		Short: "Query user acts by denom",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqAddress := args[0]
			reqDenom := args[1]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryUserActsRequest{

				Address: reqAddress,
				Denom:   reqDenom,
			}

			res, err := queryClient.UserActs(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
