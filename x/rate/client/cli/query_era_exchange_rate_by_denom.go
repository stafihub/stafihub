package cli

import (
    "strconv"
	

	
	"github.com/spf13/cobra"

    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/stafiprotocol/stafihub/x/rate/types"
)

var _ = strconv.Itoa(0)

func CmdEraExchangeRateByDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "era-exchange-rate-by-denom [denom]",
		Short: "Query EraExchangeRateByDenom",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			 reqDenom := args[0]
			
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryEraExchangeRateByDenomRequest{
				
                Denom: reqDenom, 
            }

            

			res, err := queryClient.EraExchangeRateByDenom(cmd.Context(), params)
            if err != nil {
                return err
            }

            return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

    return cmd
}