package cli

import (
    "context"



    "github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
    "github.com/stafiprotocol/stafihub/x/relayers/types"
)

func CmdListThreshold() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-threshold",
		Short: "list all threshold",
		RunE: func(cmd *cobra.Command, args []string) error {
            clientCtx := client.GetClientContextFromCmd(cmd)

            pageReq, err := client.ReadPageRequest(cmd.Flags())
            if err != nil {
                return err
            }

            queryClient := types.NewQueryClient(clientCtx)

            params := &types.QueryAllThresholdRequest{
                Pagination: pageReq,
            }

            res, err := queryClient.ThresholdAll(context.Background(), params)
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

func CmdShowThreshold() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-threshold [denom]",
		Short: "shows a threshold",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
            clientCtx := client.GetClientContextFromCmd(cmd)

            queryClient := types.NewQueryClient(clientCtx)

            params := &types.QueryGetThresholdRequest{
                Denom: args[0],

            }

            res, err := queryClient.Threshold(context.Background(), params)
            if err != nil {
                return err
            }

            return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

    return cmd
}
