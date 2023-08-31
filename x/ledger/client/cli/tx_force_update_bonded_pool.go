package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/ledger/types"
)

var _ = strconv.Itoa(0)

func CmdForceUpdateBondedPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "force-update-bonded-pool [denom] [address] [active] [bond] [unbond]",
		Short: "Force update bonded pool",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argAddress := args[1]

			argActive, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return fmt.Errorf("argActive format err")
			}
			argBond, ok := sdk.NewIntFromString(args[3])
			if !ok {
				return fmt.Errorf("argBond format err")
			}
			argUnbond, ok := sdk.NewIntFromString(args[4])
			if !ok {
				return fmt.Errorf("argUnbond format err")
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgForceUpdateBondedPool(
				clientCtx.GetFromAddress().String(),
				argDenom,
				argAddress,
				argActive,
				argBond,
				argUnbond,
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
