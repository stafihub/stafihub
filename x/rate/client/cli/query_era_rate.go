package cli

import (
    "strconv"
	

	
	"github.com/spf13/cobra"

    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/stafiprotocol/stafihub/x/rate/types"
)

var _ = strconv.Itoa(0)

func CmdEraRate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "era-rate [denom] [era]",
		Short: "Query era_rate",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			 reqDenom := args[0]
			 reqEra := args[1]
			
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryEraRateRequest{
				
                Denom: reqDenom, 
                Era: reqEra, 
            }

            

			res, err := queryClient.EraRate(cmd.Context(), params)
            if err != nil {
                return err
            }

            return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

    return cmd
}