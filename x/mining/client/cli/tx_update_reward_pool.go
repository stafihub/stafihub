package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/mining/types"
)

var _ = strconv.Itoa(0)

func CmdUpdateRewardPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-reward-pool [stake-pool-index] [reward-pool-index] [new-reward-amount]",
		Short: "Add new reward to reward pool",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argStakePoolIndex, err := sdk.ParseUint(args[0])
			if err != nil {
				return err
			}
			argRewardPoolIndex, err := sdk.ParseUint(args[1])
			if err != nil {
				return err
			}
			argNewRewardAmount, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return fmt.Errorf("argNewRewardAmount err")
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateRewardPool(
				clientCtx.GetFromAddress().String(),
				uint32(argStakePoolIndex.Uint64()),
				uint32(argRewardPoolIndex.Uint64()),
				argNewRewardAmount,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
