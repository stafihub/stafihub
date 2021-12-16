package cli

import (
    "strconv"
	

	
	"github.com/spf13/cobra"

    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/stafiprotocol/stafihub/x/ledger/types"
)

var _ = strconv.Itoa(0)

func CmdSetChainEra() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-chain-era [denom] [new-era]",
		Short: "Broadcast message set_chain_era",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
      		 argDenom := args[0]
             argNewEra := args[1]
            
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetChainEra(
				clientCtx.GetFromAddress().String(),
				argDenom,
				argNewEra,
				
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