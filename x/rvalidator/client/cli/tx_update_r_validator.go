package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/rvalidator/types"
	rvotetypes "github.com/stafihub/stafihub/x/rvote/types"
)

var _ = strconv.Itoa(0)

func CmdUpdateRValidator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-r-validator [denom] [old-address] [new-address] [cycleVersion] [cycleNumber]",
		Short: "Update rvalidator",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argOldAddress := args[1]
			argNewAddress := args[2]
			argCycleVersion, err := sdk.ParseUint(args[3])
			if err != nil {
				return err
			}
			argCycleNumber, err := sdk.ParseUint(args[4])
			if err != nil {
				return err
			}

			cycle := types.Cycle{
				Denom:   argDenom,
				Version: argCycleVersion.Uint64(),
				Number:  argCycleNumber.Uint64(),
			}
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			content := types.NewUpdateRValidatorProposal(clientCtx.GetFromAddress().String(),
				argDenom,
				argOldAddress,
				argNewAddress,
				&cycle)
			msg, err := rvotetypes.NewMsgSubmitProposal(clientCtx.GetFromAddress(), content)
			if err != nil {
				return err
			}

			if err := msg.ValidateBasic(); err != nil {
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
