package cli

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/relayers/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdAddRelayer())
	cmd.AddCommand(CmdDeleteRelayer())
	cmd.AddCommand(CmdSetThreshold())
	// this line is used by starport scaffolding # 1

	return cmd
}

func CmdAddRelayer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-relayers [arena] [denom] [addresses]",
		Short: "Add new relayers",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argArena := args[0]

			argDenom := args[1]
			if err := sdk.ValidateDenom(argDenom); err != nil {
				return err
			}

			argValidators := strings.Split(args[2], ":")
			for _, v := range argValidators {
				_, err := sdk.AccAddressFromBech32(v)
				if err != nil {
					return err
				}
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateRelayer(
				clientCtx.GetFromAddress(),
				argArena,
				argDenom,
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

func CmdDeleteRelayer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-relayer [arena] [denom] [address]",
		Short: "Delete a relayer",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argArena := args[0]

			argDenom := args[1]
			if sdk.ValidateDenom(argDenom) != nil {
				return nil
			}

			relAddr, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteRelayer(
				clientCtx.GetFromAddress(),
				argArena,
				argDenom,
				relAddr,
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

func CmdSetThreshold() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-threshold [arena] [denom] [value]",
		Short: "Set threshold",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argArena := args[0]

			argDenom := args[1]
			if sdk.ValidateDenom(argDenom) != nil {
				return nil
			}

			value, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetThreshold(
				clientCtx.GetFromAddress(),
				argArena,
				argDenom,
				uint32(value),
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
