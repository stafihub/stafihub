package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/bridge/types"
)

var _ = strconv.Itoa(0)

func CmdVoteProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vote-proposal [chain-id] [deposit-nonce] [resource-id] [amount] [receiver]",
		Short: "Vote proposal",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argChainId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			argDepositNonce, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}
			argResourceId := args[2]
			argAmount, ok := sdk.NewIntFromString(args[3])
			if !ok {
				return fmt.Errorf("amount format err")
			}
			argReceiver := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgVoteProposal(
				clientCtx.GetFromAddress().String(),
				uint32(argChainId),
				argDepositNonce,
				argResourceId,
				argAmount,
				argReceiver,
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
