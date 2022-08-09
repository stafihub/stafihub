package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/claim/types"
)

var _ = strconv.Itoa(0)

func CmdIsClaimed() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "is-claimed [round] [index]",
		Short: "Query drop is claimed",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqRound, err := sdk.ParseUint(args[0])
			if err != nil {
				return err
			}
			reqIndex, err := sdk.ParseUint(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryIsClaimedRequest{
				Round: reqRound.Uint64(),
				Index: reqIndex.Uint64(),
			}

			res, err := queryClient.IsClaimed(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
