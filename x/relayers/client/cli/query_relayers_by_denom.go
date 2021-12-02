package cli

import (
    "strconv"



	"github.com/spf13/cobra"

    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/stafiprotocol/stafihub/x/relayers/types"
)

var _ = strconv.Itoa(0)

func CmdRelayersByDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "relayers-by-denom [denom]",
		Short: "Query relayers_by_denom",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			 reqDenom := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryRelayersByDenomRequest{
                Denom: reqDenom,
            }

			res, err := queryClient.RelayersByDenom(cmd.Context(), params)
            if err != nil {
                return err
            }

            return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

    return cmd
}