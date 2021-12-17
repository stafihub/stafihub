package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/stafiprotocol/stafihub/x/ledger/types"
	"strconv"
)

var _ = strconv.Itoa(0)

func CmdSetChainEra() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-chain-era [denom] [era]",
		Short: "Broadcast message set_chain_era",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argEra, err := strconv.ParseUint(args[1], 10, 32)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetChainEra(
				clientCtx.GetFromAddress().String(),
				argDenom,
				uint32(argEra),
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