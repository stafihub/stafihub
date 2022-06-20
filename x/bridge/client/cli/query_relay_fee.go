package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/bridge/types"
)

var _ = strconv.Itoa(0)

func CmdRelayFee() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "relay-fee [chain-id]",
		Short: "Query relay fee",
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

			params := &types.QueryRelayFeeRequest{
				ChainId: uint32(argChainId),
			}

			res, err := queryClient.RelayFee(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
