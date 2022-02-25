package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/rstaking/types"
)

var _ = strconv.Itoa(0)

func CmdAddValToWhitelist() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-val-to-whitelist [val-address]",
		Short: "Broadcast message add validator address to whitelist",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argValAddress := args[0]
			_, err = sdk.ValAddressFromBech32(argValAddress)
			if err != nil {
				return err
			}
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddValToWhitelist(
				clientCtx.GetFromAddress().String(),
				argValAddress,
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
