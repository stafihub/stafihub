package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/stafiprotocol/stafihub/x/ledger/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ = strconv.Itoa(0)

func CmdAddNewPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-new-pool [denom] [addr]",
		Short: "Broadcast message add_new_pool",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
      		 argDenom := args[0]
      		 argAddr := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddNewPool(
				clientCtx.GetFromAddress(),
				argDenom,
				argAddr,
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

func CmdRemovePool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-pool [denom] [addr]",
		Short: "Broadcast message remove_pool",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argAddr := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRemovePool(
				clientCtx.GetFromAddress(),
				argDenom,
				argAddr,

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

func CmdSetEraUnbondLimit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-era-unbond-limit [denom] [limit]",
		Short: "Broadcast message set_era_unbond_limit",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argLimit, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetEraUnbondLimit(
				clientCtx.GetFromAddress(),
				argDenom,
				uint32(argLimit),
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

func CmdSetInitBond() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-init-bond [denom] [pool] [amount] [receiver]",
		Short: "Broadcast message set_init_bond",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argPool := args[1]
			argAmount, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return fmt.Errorf("cast amount %s into Int error", args[2])
			}

			argReceiver, err := sdk.AccAddressFromBech32(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetInitBond(
				clientCtx.GetFromAddress(),
				argDenom,
				argPool,
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

func CmdSetChainBondingDuration() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-chain-bonding-duration [denom] [era]",
		Short: "Broadcast message set_chain_bonding_duration",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]

			argEra, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetChainBondingDuration(
				clientCtx.GetFromAddress(),
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

func CmdSetPoolDetail() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-pool-detail [denom] [pool] [sub-accounts] [threshold]",
		Short: "Broadcast message set_pool_detail",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argPool := args[1]

			argSubAccounts := strings.Split(args[2], "|")
			argThreshold, err := strconv.ParseUint(args[3], 10, 32)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetPoolDetail(
				clientCtx.GetFromAddress(),
				argDenom,
				argPool,
				argSubAccounts,
				uint32(argThreshold),
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

func CmdSetLeastBond() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-least-bond [denom] [amount]",
		Short: "Broadcast message set_least_bond",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argAmount, ok := sdk.NewIntFromString(args[1])
			if !ok {
				return fmt.Errorf("amount %s cast error")
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetLeastBond(
				clientCtx.GetFromAddress(),
				argDenom,
				argAmount,
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

func CmdClearCurrentEraSnapShots() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clear-current-era-snap-shots [denom]",
		Short: "Broadcast message clear_current_era_snap_shots",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgClearCurrentEraSnapShots(
				clientCtx.GetFromAddress(),
				argDenom,

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

func CmdSetCommission() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-commission [rate]",
		Short: "Broadcast message set_commission",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCommission, err := sdk.NewDecFromStr(args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetCommission(
				clientCtx.GetFromAddress(),
				argCommission,
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

func CmdSetReceiver() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-receiver [receiver]",
		Short: "Broadcast message set_receiver",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argReceiver, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetReceiver(
				clientCtx.GetFromAddress(),
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

func CmdSetUnbondFee() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-unbond-fee [denom] [value]",
		Short: "Broadcast message set_unbond_fee",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argValue, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetUnbondFee(
				clientCtx.GetFromAddress(),
				argDenom,
				argValue,

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

func CmdSetUnbondCommission() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-unbond-commission [commission]",
		Short: "Broadcast message set_unbond_commission",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCommission, err := sdk.NewDecFromStr(args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetUnbondCommission(
				clientCtx.GetFromAddress(),
				argCommission,

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