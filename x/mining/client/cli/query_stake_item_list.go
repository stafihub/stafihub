package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/mining/types"
)

var _ = strconv.Itoa(0)

func CmdStakeItemList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stake-item-list [stake-pool-index]",
		Short: "Query stake item list",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argStakePoolIndex, err := sdk.ParseUint(args[0])
			if err != nil {
				return err
			}
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryStakeItemListRequest{
				StakePoolIndex: uint32(argStakePoolIndex.Uint64()),
			}

			res, err := queryClient.StakeItemList(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
