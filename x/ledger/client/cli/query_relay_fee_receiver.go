package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/ledger/types"
)

var _ = strconv.Itoa(0)

func CmdRelayFeeReceiver() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "relay-fee-receiver [denom]",
		Short: "Query relay fee receiver",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryRelayFeeReceiverRequest{
				Denom: args[0],
			}

			res, err := queryClient.RelayFeeReceiver(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
