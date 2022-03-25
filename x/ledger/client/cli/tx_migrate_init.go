package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/utils"
	"github.com/stafihub/stafihub/x/ledger/types"
)

var _ = strconv.Itoa(0)

func CmdMigrateInit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate-init [denom] [pool] [total-supply] [active] [bond] [unbond] [exchange-rate] [total-protocol-fee]",
		Short: "Migrate init",

		Args: cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argPool := args[1]
			argTotalSupply, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return fmt.Errorf("argDenom format err")
			}
			argActive, ok := sdk.NewIntFromString(args[3])
			if !ok {
				return fmt.Errorf("argActive format err")
			}
			argBond, ok := sdk.NewIntFromString(args[4])
			if !ok {
				return fmt.Errorf("argBond format err")
			}
			argUnbond, ok := sdk.NewIntFromString(args[5])
			if !ok {
				return fmt.Errorf("argUnbond format err")
			}
			argExchangeRate, err := utils.NewDecFromStr(args[6])
			if err != nil {
				return err
			}
			argTotalProtocolFee, ok := sdk.NewIntFromString(args[7])
			if !ok {
				return fmt.Errorf("argUnbond format err")
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgMigrateInit(
				clientCtx.GetFromAddress().String(),
				argDenom,
				argPool,
				argTotalSupply,
				argActive,
				argBond,
				argUnbond,
				argExchangeRate,
				argTotalProtocolFee,
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
