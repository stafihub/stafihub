package cli

import (
    "context"
	

	
    "github.com/spf13/cobra"
    
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
    "github.com/stafiprotocol/stafihub/x/rate/types"
)

func CmdListEraExchangeRate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-era-exchange-rate",
		Short: "list all EraExchangeRate",
		RunE: func(cmd *cobra.Command, args []string) error {
            clientCtx := client.GetClientContextFromCmd(cmd)

            pageReq, err := client.ReadPageRequest(cmd.Flags())
            if err != nil {
                return err
            }

            queryClient := types.NewQueryClient(clientCtx)

            params := &types.QueryAllEraExchangeRateRequest{
                Pagination: pageReq,
            }

            res, err := queryClient.EraExchangeRateAll(context.Background(), params)
            if err != nil {
                return err
            }

            return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

    return cmd
}

func CmdShowEraExchangeRate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-era-exchange-rate [index]",
		Short: "shows a EraExchangeRate",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
            clientCtx := client.GetClientContextFromCmd(cmd)

            queryClient := types.NewQueryClient(clientCtx)

             argIndex := args[0]
            
            params := &types.QueryGetEraExchangeRateRequest{
                Index: argIndex,
                
            }

            res, err := queryClient.EraExchangeRate(context.Background(), params)
            if err != nil {
                return err
            }

            return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

    return cmd
}
