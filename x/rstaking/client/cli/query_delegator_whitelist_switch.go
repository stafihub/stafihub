package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/rstaking/types"
)

var _ = strconv.Itoa(0)

func CmdDelegatorWhitelistSwitch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delegator-whitelist-switch",
		Short: "query delegator whitelist switch",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryDelegatorWhitelistSwitchRequest{}

			res, err := queryClient.DelegatorWhitelistSwitch(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
