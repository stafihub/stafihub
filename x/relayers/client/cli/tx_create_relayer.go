package cli

import (
    "strconv"



	"github.com/spf13/cobra"

    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/stafiprotocol/stafihub/x/relayers/types"
)

var _ = strconv.Itoa(0)

func CmdCreateRelayer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-relayer [denom] [address]",
		Short: "Broadcast message create_relayer",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
      		 argDenom := args[0]
             argAddress := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateRelayer(
				argDenom,
				argAddress,
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