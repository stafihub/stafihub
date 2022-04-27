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

func CmdRemoveLiquidity() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-liquidity [denom] [rm-unit] [swap-unit] [min-fis-out-amount] [min-rtoken-out-amount] [input-is-fis]",
		Short: "Remove liquidity",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argRmUnit, ok := sdk.NewIntFromString(args[1])
			if !ok {
				return fmt.Errorf("rm unit params err")
			}
			argSwapUnit, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return fmt.Errorf("swap unit params err")
			}
			argMinFisOutAmount, ok := sdk.NewIntFromString(args[3])
			if !ok {
				return fmt.Errorf("min fis out amount params err")
			}
			argMinRtokenOutAmount, ok := sdk.NewIntFromString(args[4])
			if !ok {
				return fmt.Errorf("min rtoken out amount params err")
			}

			argInputIsFis := false
			if args[5] == "true" {
				argInputIsFis = true
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRemoveLiquidity(
				clientCtx.GetFromAddress().String(),
				argDenom,
				argRmUnit,
				argSwapUnit,
				argMinFisOutAmount,
				argMinRtokenOutAmount,
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
