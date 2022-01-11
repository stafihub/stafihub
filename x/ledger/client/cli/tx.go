package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/stafiprotocol/stafihub/x/ledger/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

	cmd.AddCommand(CmdAddNewPool())
	cmd.AddCommand(CmdRemovePool())
	cmd.AddCommand(CmdSetEraUnbondLimit())
	cmd.AddCommand(CmdSetInitBond())
	cmd.AddCommand(CmdSetChainBondingDuration())
	cmd.AddCommand(CmdSetPoolDetail())
	cmd.AddCommand(CmdSetLeastBond())
	cmd.AddCommand(CmdClearCurrentEraSnapShots())
	cmd.AddCommand(CmdSetCommission())
	cmd.AddCommand(CmdSetReceiver())
	cmd.AddCommand(CmdSetChainEra())
	cmd.AddCommand(CmdActiveReport())
	cmd.AddCommand(CmdBondReport())
	cmd.AddCommand(CmdBondAndReportActive())
	cmd.AddCommand(CmdWithdrawReport())
	cmd.AddCommand(CmdTransferReport())
	cmd.AddCommand(CmdSetUnbondFee())
	cmd.AddCommand(CmdSetUnbondCommission())
	cmd.AddCommand(CmdLiquidityUnbond())


// this line is used by starport scaffolding # 1

	return cmd
}

func CmdLiquidityUnbond() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "liquidity-unbond [denom] [pool] [value] [recipient]",
		Short: "Broadcast message liquidity_unbond",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argPool := args[1]
			argValue, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return fmt.Errorf("cast value %s into Int error", args[2])
			}
			argRecipient := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgLiquidityUnbond(
				clientCtx.GetFromAddress(),
				argDenom,
				argPool,
				argValue,
				argRecipient,

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
