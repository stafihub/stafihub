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

func CmdSetEraSeconds() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-era-seconds [denom] [era-seconds] [bonding-duration] [offset]",
		Short: "Set eraSeconds, bonding duration and offset",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argEraSeconds := args[1]
			argBondingDuration, err := strconv.ParseUint(args[2], 10, 32)
			if err != nil {
				return err
			}
			argOffset := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetEraSeconds(
				clientCtx.GetFromAddress().String(),
				argDenom,
				argEraSeconds,
				uint32(argBondingDuration),
				argOffset,
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
