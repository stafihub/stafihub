package cli

import (
    "context"



    "github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
    "github.com/stafiprotocol/stafihub/x/relayers/types"
)

func CmdListRelayer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-relayer",
		Short: "list all relayer",
		RunE: func(cmd *cobra.Command, args []string) error {
            clientCtx := client.GetClientContextFromCmd(cmd)

            pageReq, err := client.ReadPageRequest(cmd.Flags())
            if err != nil {
                return err
            }

            queryClient := types.NewQueryClient(clientCtx)

            params := &types.QueryAllRelayerRequest{
                Pagination: pageReq,
            }

            res, err := queryClient.RelayerAll(context.Background(), params)
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
