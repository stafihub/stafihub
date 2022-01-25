package cli

import (
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/stafiprotocol/stafihub/x/ledger/types"
	rvotetypes "github.com/stafiprotocol/stafihub/x/rvote/types"
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

			from := clientCtx.GetFromAddress()
			content := types.NewSetChainEraProposal(from, argDenom, uint32(argEra))
			msg, err := rvotetypes.NewMsgSubmitProposal(from, content)
			if err != nil {
				return err
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdBondReport() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bond-report [denom] [shot-id] [action]",
		Short: "Broadcast message bond_report",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argShotId, err := hex.DecodeString(args[1])
			if err != nil {
				return err
			}
			argAction, ok := types.BondAction_value[args[2]]
			if !ok {
				return fmt.Errorf("cannot cast %s into bondAction", args[2])
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()
			content := types.NewBondReportProposal(from, argDenom, argShotId, types.BondAction(argAction))
			msg, err := rvotetypes.NewMsgSubmitProposal(from, content)
			if err != nil {
				return err
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdBondAndReportActive() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bond-and-report-active [denom] [shot-id] [action] [staked] [unstaked]",
		Short: "Broadcast message bond_and_report_active",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argShotId, err := hex.DecodeString(args[1])
			if err != nil {
				return err
			}
			argAction, ok := types.BondAction_value[args[2]]
			if !ok {
				return fmt.Errorf("cannot cast %s into bondAction", args[2])
			}

			argStaked, ok := sdk.NewIntFromString(args[3])
			if !ok {
				return fmt.Errorf("cast staked %s into Int error", args[3])
			}
			argUnstaked, ok := sdk.NewIntFromString(args[4])
			if !ok {
				return fmt.Errorf("cast unstaked %s into Int error", args[4])
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()
			content := types.NewBondAndReportActiveProposal(from, argDenom, argShotId, types.BondAction(argAction), argStaked, argUnstaked)
			msg, err := rvotetypes.NewMsgSubmitProposal(from, content)
			if err != nil {
				return err
			}

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
			argShotId, err := hex.DecodeString(args[1])
			if err != nil {
				return err
			}
			argStaked, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return fmt.Errorf("cast staked %s into Int error", args[2])
			}
			argUnstaked, ok := sdk.NewIntFromString(args[3])
			if !ok {
				return fmt.Errorf("cast unstaked %s into Int error", args[3])
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()
			content := types.NewActiveReportProposal(from, argDenom, argShotId, argStaked, argUnstaked)
			msg, err := rvotetypes.NewMsgSubmitProposal(from, content)
			if err != nil {
				return err
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdWithdrawReport() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw-report [denom] [shot-id]",
		Short: "Broadcast message withdraw_report",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argShotId, err := hex.DecodeString(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()
			content := types.NewWithdrawReportProposal(from, argDenom, argShotId)
			msg, err := rvotetypes.NewMsgSubmitProposal(from, content)
			if err != nil {
				return err
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdTransferReport() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-report [denom] [shot-id]",
		Short: "Broadcast message transfer_report",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argShotId, err := hex.DecodeString(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()
			content := types.NewTransferReportProposal(from, argDenom, argShotId)
			msg, err := rvotetypes.NewMsgSubmitProposal(from, content)
			if err != nil {
				return err
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdExecuteBondProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "execute-bond-proposal [denom] [bonder] [pool] [blockhash] [txhash] [amount]",
		Short: "Broadcast message execute_bond_proposal",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argBonder, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}
			argPool := args[2]
			argBlockhash := args[3]
			argTxHash := args[4]
			argAmount, ok := sdk.NewIntFromString(args[5])
			if !ok {
				return fmt.Errorf("cast amount %s into Int error", args[5])
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()
			content := types.NewExecuteBondProposal(from, argDenom, argBonder, argPool, argBlockhash, argTxHash, argAmount)
			msg, err := rvotetypes.NewMsgSubmitProposal(from, content)
			if err != nil {
				return err
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
