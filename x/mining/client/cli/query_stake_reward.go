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

func CmdStakeReward() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stake-reward [stake-user-address] [stake-token-denom] [stake-record-index]",
		Short: "Query stake reward",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqAddress := args[0]
			stakePoolIndex, err := sdk.ParseUint(args[1])
			if err != nil {
				return err
			}
			reqIndex, err := sdk.ParseUint(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryStakeRewardRequest{
				StakeUserAddress: reqAddress,
				StakePoolIndex:   uint32(stakePoolIndex.Uint64()),
				StakeRecordIndex: uint32(reqIndex.Uint64()),
			}

			res, err := queryClient.StakeReward(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
