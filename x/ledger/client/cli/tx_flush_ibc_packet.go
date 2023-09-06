package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/ledger/types"
)

var _ = strconv.Itoa(0)

func CmdFlushIbcPacket() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "flush-ibc-packet [port-id] [channel-id] [sequence]",
		Short: "Flush ibc packet",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPortID := args[0]
			argChannelID := args[1]
			argSequence, err := strconv.ParseUint(args[2], 10, 32)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgFlushIbcPacket(
				clientCtx.GetFromAddress().String(),
				argPortID,
				argChannelID,
				uint32(argSequence),
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
