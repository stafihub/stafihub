package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/stakextra/types"
)

var _ = strconv.Itoa(0)

func CmdSetInflationBase() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-inflation-base [inflation-base]",
		Short: "set inflation base",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argInflationBase, ok := sdk.NewIntFromString(args[0])
			if !ok {
				return fmt.Errorf("args[0] cast to int failed")
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetInflationBase(
				clientCtx.GetFromAddress().String(),
				argInflationBase,
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
