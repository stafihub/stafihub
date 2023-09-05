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

func CmdSetInterchainTxProposalStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-interchain-tx-proposal-status [proposal-id] [status]",
		Short: "Set interchain tx proposal status",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argProposalId := args[0]
			argStatus := args[1]

			if _, exist := types.InterchainTxStatus_value[argStatus]; !exist {
				return fmt.Errorf("status not exist")
			}
			status := types.InterchainTxStatus(types.InterchainTxStatus_value[argStatus])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetInterchainTxProposalStatus(
				clientCtx.GetFromAddress().String(),
				argProposalId,
				status,
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
