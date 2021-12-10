package cli

import (
	"encoding/hex"
	"fmt"
	"strconv"



	"github.com/spf13/cobra"

    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/stafiprotocol/stafihub/x/relayers/types"
)

var _ = strconv.Itoa(0)

func CmdSubmitProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-proposal [proposal_route] [name] [params] [in_favour]",
		Short: "Broadcast message submit_proposal, in_favour should only be true or false",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
      		 argProposalRoute := args[0]
             argName := args[1]
             argParams, err := hex.DecodeString(args[2])
             if err != nil {
             	return err
			 }
             var argInFavour bool
			switch args[3] {
			case "true":
				argInFavour = true
			case "false":
				argInFavour = false
			default:
				return fmt.Errorf("in_favour neither true of false")
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSubmitProposal(
				clientCtx.GetFromAddress().String(),
				argProposalRoute,
				argName,
				argParams,
				argInFavour,

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