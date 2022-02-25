package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/bridge/types"
)

var _ = strconv.Itoa(0)

func CmdSetResourceidToDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-resourceid-to-denom [resource-id] [denom]",
		Short: "Broadcast message set resourceid to denom",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argResourceId := args[0]
			argDenom := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetResourceidToDenom(
				clientCtx.GetFromAddress().String(),
				argResourceId,
				argDenom,
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
