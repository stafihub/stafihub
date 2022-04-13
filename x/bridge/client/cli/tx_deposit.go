package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/bridge/types"
)

var _ = strconv.Itoa(0)

func CmdDeposit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit [dest-chain-id] [denom] [amount] [receiver]",
		Short: "Broadcast message initiates a transfer to other chain",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDestChainId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			argDenom := args[1]
			argAmount, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return fmt.Errorf("amount format err")
			}
			argReceiver := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeposit(
				clientCtx.GetFromAddress().String(),
				uint32(argDestChainId),
				argDenom,
				argAmount,
				argReceiver,
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
