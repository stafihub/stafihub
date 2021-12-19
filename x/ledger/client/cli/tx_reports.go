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

func CmdActiveReport() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "active-report [denom] [shot-id] [staked] [unstaked]",
		Short: "Broadcast message active_report",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argShotId := args[1]
			argStaked := args[2]
			argUnstaked := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgActiveReport(
				clientCtx.GetFromAddress().String(),
				argDenom,
				argShotId,
				argStaked,
				argUnstaked,

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