package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/ledger/types"
	"strings"
)

var _ = strconv.Itoa(0)

func CmdSetRParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-r-params [denom] [chain-id] [native-denom] [gas-price] [era-seconds] [validators]",
		Short: "set common params of relayers",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argChainId := args[1]
			argNativeDenom := args[2]
			argGasPrice := args[3]
			argEraSeconds := args[4]
			argValidators := strings.Split(args[5], listSeparator)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetRParams(
				clientCtx.GetFromAddress().String(),
				argDenom,
				argChainId,
				argNativeDenom,
				argGasPrice,
				argEraSeconds,
				argValidators,
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
