package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/bridge/types"
)

var _ = strconv.Itoa(0)

func CmdGetDepositCount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-deposit-count [chain-id]",
		Short: "Query depositCount",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argChainId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetDepositCountRequest{
				ChainId: uint32(argChainId),
			}

			res, err := queryClient.GetDepositCount(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
