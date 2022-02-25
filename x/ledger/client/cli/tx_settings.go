package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/ledger/types"
)

var _ = strconv.Itoa(0)

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
		Use:   "set-init-bond [pool] [coin] [receiver]",
		Short: "Broadcast message set_init_bond",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPool := args[0]
			argCoin, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			argReceiver, err := sdk.AccAddressFromBech32(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetInitBond(
				clientCtx.GetFromAddress(),
				argPool,
				argCoin,
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

			argSubAccounts := strings.Split(args[2], "+")
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
				return fmt.Errorf("amount %s cast error", args[1])
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

func CmdSetStakingRewardCommission() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-staking-reward-commission [denom] [rate]",
		Short: "Broadcast message set staking reward commission",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argCommission, err := sdk.NewDecFromStr(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetStakingRewardCommission(
				clientCtx.GetFromAddress(),
				argDenom,
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

func CmdSetProtocolFeeReceiver() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-protocol-fee-receiver [receiver]",
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

			msg := types.NewMsgSetProtocolFeeReceiver(
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

func CmdSetUnbondRelayFee() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-unbond-relay-fee [denom] [value]",
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

			msg := types.NewMsgSetUnbondRelayFee(
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
		Use:   "set-unbond-commission [denom] [commission]",
		Short: "Broadcast message set_unbond_commission",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argCommission, err := sdk.NewDecFromStr(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetUnbondCommission(
				clientCtx.GetFromAddress(),
				argDenom,
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

func CmdSetRParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-r-params [denom] [chain-id] [native-denom] [gas-price] [era-seconds] [least-bond] [validators]",
		Short: "Broadcast message set common params of relayers",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argChainId := args[1]
			argNativeDenom := args[2]
			argGasPrice := args[3]
			argEraSeconds := args[4]
			argLeastBond, ok := sdk.NewIntFromString(args[5])
			if !ok {
				return fmt.Errorf("amount %s cast error", args[5])
			}
			argValidators := strings.Split(args[6], "+")

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetRParams(
				clientCtx.GetFromAddress().String(),
				argDenom,
				argChainId,
				argNativeDenom,
				argGasPrice,
				argEraSeconds,
				argLeastBond,
				argValidators,
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
