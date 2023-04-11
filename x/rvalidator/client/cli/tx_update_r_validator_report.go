package cli

import (
	"fmt"
	"strconv"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/rvalidator/types"
	rvotetypes "github.com/stafihub/stafihub/x/rvote/types"
)

var _ = strconv.Itoa(0)

func CmdUpdateRValidatorReport() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-r-validator-report [denom] [pool-address] [cycle-version] [cycle-number] [status]",
		Short: "Update rvalidator report",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argPoolAddress := args[1]
			argCycleVersion, err := math.ParseUint(args[2])
			if err != nil {
				return err
			}
			argCycleNumber, err := math.ParseUint(args[3])
			if err != nil {
				return err
			}
			status, exist := types.UpdateRValidatorStatus_value[args[4]]
			if !exist {
				return fmt.Errorf("status not exist")
			}
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			cycle := types.Cycle{
				Denom:       argDenom,
				PoolAddress: argPoolAddress,
				Version:     argCycleVersion.Uint64(),
				Number:      argCycleNumber.Uint64(),
			}
			content := types.NewUpdateRValidatorReportProposal(
				clientCtx.GetFromAddress().String(),
				argDenom,
				argPoolAddress,
				&cycle, types.UpdateRValidatorStatus(status))
			msg, err := rvotetypes.NewMsgSubmitProposal(clientCtx.GetFromAddress(), content)
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
