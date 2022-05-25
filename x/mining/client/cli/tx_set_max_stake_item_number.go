package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/mining/types"
)

var _ = strconv.Itoa(0)

func CmdSetMaxStakeItemNumber() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-max-stake-item-number [number]",
		Short: "Set max stake item number",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argNumber, err := sdk.ParseUint(args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetMaxStakeItemNumber(
				clientCtx.GetFromAddress().String(),
				uint32(argNumber.Uint64()),
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
