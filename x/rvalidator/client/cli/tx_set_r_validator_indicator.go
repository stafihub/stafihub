package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rvalidator/types"
)

var _ = strconv.Itoa(0)

func CmdSetRValidatorIndicator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-r-validator-indicator [denom] [commission] [uptime] [locked]",
		Short: "Broadcast message SetRValidatorIndicator",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argCommission, err := sdk.NewDecFromStr(args[1])
			if err != nil {
				return err
			}

			argUptime, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}
			argLocked, err := sdk.ParseCoinNormalized(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetRValidatorIndicator(
				clientCtx.GetFromAddress(),
				argDenom,
				argCommission,
				uint32(argUptime),
				argLocked,
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
