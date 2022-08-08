package cli

import (
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/utils"
	"github.com/stafihub/stafihub/x/ledger/types"
)

var _ = strconv.Itoa(0)

func CmdSetEraUnbondLimit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-era-unbond-limit [denom] [limit]",
		Short: "Broadcast message set era unbond limit",
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

func CmdSetPoolDetail() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-pool-detail [denom] [pool] [sub-accounts] [threshold]",
		Short: "Broadcast message set pool detail",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argPool := args[1]

			argSubAccounts := []string{}
			if len(args[2]) != 0 {
				argSubAccounts = strings.Split(args[2], ":")
			}
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
		Use:   "set-least-bond [denom] [least bond]",
		Short: "Broadcast message set least bond",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argLeastBond := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetLeastBond(
				clientCtx.GetFromAddress(),
				argDenom,
				argLeastBond,
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
		Use:   "clear-current-era-snapshots [denom]",
		Short: "Broadcast message clear current era snapshots",
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
			argCommission, err := utils.NewDecFromStr(args[1])
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
		Short: "Broadcast message set protocol fee receiver",
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
		Short: "Broadcast message set unbond relay fee",
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
		Short: "Broadcast message set unbond commission",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argCommission, err := utils.NewDecFromStr(args[1])
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
		Use:   "set-r-params [denom] [gas-price] [era-seconds] [offset] [bonding-duration] [least-bond]",
		Short: "Broadcast message set common params of rtoken relayers",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argGasPrice := args[1]
			argEraSeconds, err := strconv.ParseUint(args[2], 10, 32)
			if err != nil {
				return err
			}

			argOffset, err := strconv.ParseInt(args[3], 10, 32)
			if err != nil {
				return err
			}

			argBondingDuration, err := strconv.ParseUint(args[4], 10, 32)
			if err != nil {
				return err
			}
			argLeastBond := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetRParams(
				clientCtx.GetFromAddress().String(),
				argDenom,
				argGasPrice,
				uint32(argEraSeconds),
				int32(argOffset),
				uint32(argBondingDuration),
				argLeastBond,
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
