package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/claim/types"
)

var _ = strconv.Itoa(0)

func CmdWithdrawToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw-token [recipient] [token]",
		Short: "Withdraw token from claim module",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argRecipient := args[0]
			argToken, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgWithdrawToken(
				clientCtx.GetFromAddress().String(),
				argRecipient,
				argToken,
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
