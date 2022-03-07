package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/stafihub/stafihub/x/rvalidator/types"
)

var _ = strconv.Itoa(0)

func CmdVoteStatusProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vote-status-proposal [denom] [addresses] [source-status] [dest-status]",
		Short: "Broadcast message voteStatusProposal",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argAddresses := strings.Split(args[1], ":")

			argSourceStatus, ok := types.RValidatorStatus_value[args[2]]
			if !ok {
				return fmt.Errorf("cannot cast argSourceStatus %s into RValidatorStatus", args[2])
			}

			argDestStatus, ok := types.RValidatorStatus_value[args[3]]
			if !ok {
				return fmt.Errorf("cannot cast argDestStatus %s into RValidatorStatus", args[3])
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgVoteStatusProposal(
				clientCtx.GetFromAddress(),
				argDenom,
				argAddresses,
				types.RValidatorStatus(argSourceStatus),
				types.RValidatorStatus(argDestStatus),
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
