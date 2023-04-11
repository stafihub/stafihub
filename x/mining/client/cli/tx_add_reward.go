package cli

import (
	"fmt"
	"strconv"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/mining/types"
)

var _ = strconv.Itoa(0)

func CmdAddReward() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-reward [stake-pool-index] [reward-pool-index] [add-amount] [start-timestamp] [reward-per-second]",
		Short: "Add new reward to reward pool, if pool is not end, startTimestamp/rewardPersecond should be zero",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argStakePoolIndex, err := math.ParseUint(args[0])
			if err != nil {
				return err
			}
			argRewardPoolIndex, err := math.ParseUint(args[1])
			if err != nil {
				return err
			}
			argAddAmount, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return fmt.Errorf("argAddAmount err")
			}
			argStartTimestamp, err := math.ParseUint(args[3])
			if err != nil {
				return err
			}
			argRewardPerSecond, ok := sdk.NewIntFromString(args[4])
			if !ok {
				return fmt.Errorf("argRewardPersecond err")
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddReward(
				clientCtx.GetFromAddress().String(),
				uint32(argStakePoolIndex.Uint64()),
				uint32(argRewardPoolIndex.Uint64()),
				argAddAmount,
				argStartTimestamp.Uint64(),
				argRewardPerSecond,
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
