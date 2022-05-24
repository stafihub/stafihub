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

func CmdAddRewardPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-reward-pool [stake-pool-index] [reward-token-denom] [total-reward-amount] [reward-per-second] [start-timestamp]",
		Short: "Add reward pool",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argStakePoolIndex, err := sdk.ParseUint(args[0])
			if err != nil {
				return err
			}

			argRewardTokenDenom := args[1]

			argTotalRewardAmount, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return fmt.Errorf("totalRewardAmount params err")
			}

			argRewardPerSecond, ok := sdk.NewIntFromString(args[3])
			if !ok {
				return fmt.Errorf("rewardPerSecond params err")
			}

			argStartTimestamp, err := sdk.ParseUint(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddRewardPool(
				clientCtx.GetFromAddress().String(),
				uint32(argStakePoolIndex.Uint64()),
				argRewardTokenDenom,
				argTotalRewardAmount,
				argRewardPerSecond,
				argStartTimestamp.Uint64(),
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
