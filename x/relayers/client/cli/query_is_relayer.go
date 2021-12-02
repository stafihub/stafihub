package cli

import (
    "strconv"



	"github.com/spf13/cobra"

    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/stafiprotocol/stafihub/x/relayers/types"
)

var _ = strconv.Itoa(0)

func CmdIsRelayer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "is-relayer [denom] [address]",
		Short: "Query is_relayer",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			 reqDenom := args[0]
			 reqAddress := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryIsRelayerRequest{

                Denom: reqDenom,
                Address: reqAddress,
            }

			res, err := queryClient.IsRelayer(cmd.Context(), params)
            if err != nil {
                return err
            }

            return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

    return cmd
}