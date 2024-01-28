package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/ledger/types"
)

var _ = strconv.Itoa(0)

func CmdOpenIcaChannel() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "open-ica-channel [pool-address] [account-type]",
		Short: "Open ica channel",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPoolAddress := args[0]
			argAccountType := args[1]

			if _, exist := types.AccountType_value[argAccountType]; !exist {
				return fmt.Errorf("status not exist")
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgOpenIcaChannel(
				clientCtx.GetFromAddress().String(),
				argPoolAddress,
				types.AccountType(types.AccountType_value[argAccountType]),
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
