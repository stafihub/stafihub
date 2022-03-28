package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/rmintreward/types"
)

var _ = strconv.Itoa(0)

func CmdUserMintCount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user-mint-count [address] [denom] [cycle]",
		Short: "Query user mint count by denom and cycle",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqAddress := args[0]
			reqDenom := args[1]

			reqCycle, err := sdk.ParseUint(args[2])
			if err != nil {
				return err
			}
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryUserMintCountRequest{

				Address: reqAddress,
				Denom:   reqDenom,
				Cycle:   reqCycle.Uint64(),
			}

			res, err := queryClient.UserMintCount(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
