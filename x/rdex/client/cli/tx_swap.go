package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/rdex/types"
)

var _ = strconv.Itoa(0)

func CmdSwap() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swap [denom] [input-amount] [min-out-amount] [input-is-fis(true/false)]",
		Short: "Swap ",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argInputAmount, ok := sdk.NewIntFromString(args[1])
			if !ok {
				return fmt.Errorf("argInputAmount format err")
			}
			argMinOutAmount, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return fmt.Errorf("argMinOutAmount format err")
			}

			argInputIsFis := false
			if args[3] == "true" {
				argInputIsFis = true
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSwap(
				clientCtx.GetFromAddress().String(),
				argDenom,
				argInputAmount,
				argMinOutAmount,
				argInputIsFis,
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
