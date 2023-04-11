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

func CmdStake() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stake [stake-pool-index] [stake-amount] [stake-item-index]",
		Short: "Stake token ",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argStakePoolIndex, err := math.ParseUint(args[0])
			if err != nil {
				return err
			}

			argStakeAmount, ok := sdk.NewIntFromString(args[1])
			if !ok {
				return fmt.Errorf("argStakeAmount err")
			}

			argStakeItemIndex, err := math.ParseUint(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgStake(
				clientCtx.GetFromAddress().String(),
				uint32(argStakePoolIndex.Uint64()),
				argStakeAmount,
				uint32(argStakeItemIndex.Uint64()),
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
