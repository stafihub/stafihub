package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/ledger/types"
)

var _ = strconv.Itoa(0)

func CmdPoolUnbondNextSequence() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pool-unbond-next-sequence [denom] [pool] [unlock-era]",
		Short: "Query pool unbond next sequence",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDenom := args[0]
			reqPool := args[1]
			reqUnlockEra, err := sdk.ParseUint(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryPoolUnbondNextSequenceRequest{

				Denom:     reqDenom,
				Pool:      reqPool,
				UnlockEra: uint32(reqUnlockEra.Uint64()),
			}

			res, err := queryClient.PoolUnbondNextSequence(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
