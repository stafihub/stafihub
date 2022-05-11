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
		Use:   "remove-liquidity [rm-unit] [swap-unit] [min-out-tokens] [input-token-denom]",
		Short: "Remove liquidity",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argRmUnit, ok := sdk.NewIntFromString(args[0])
			if !ok {
				return fmt.Errorf("rm unit params err")
			}
			argSwapUnit, ok := sdk.NewIntFromString(args[1])
			if !ok {
				return fmt.Errorf("swap unit params err")
			}

			minOutTokens, err := sdk.ParseCoinsNormalized(args[2])
			if err != nil {
				return err
			}

			inputTokenDenom := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRemoveLiquidity(
				clientCtx.GetFromAddress().String(),
				argRmUnit,
				argSwapUnit,
				minOutTokens,
				inputTokenDenom,
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
