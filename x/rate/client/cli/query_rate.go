package cli

import (
    "strconv"
	

	
	"github.com/spf13/cobra"

    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/stafiprotocol/stafihub/x/rate/types"
)

var _ = strconv.Itoa(0)

func CmdRate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rate [denom]",
		Short: "Query rate",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			 reqDenom := args[0]
			
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryRateRequest{
				
                Denom: reqDenom, 
            }

            

			res, err := queryClient.Rate(cmd.Context(), params)
            if err != nil {
                return err
            }

            return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

    return cmd
}