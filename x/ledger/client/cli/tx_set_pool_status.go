package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/ledger/types"
)

var _ = strconv.Itoa(0)

func CmdSetPoolStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-pool-status [denom] [pool] [status](support 'ACTIVE' 'NOT_ACTIVE')",
		Short: "set pool status",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argPool := args[1]
			argStatus := args[2]

			if _, exist := types.PoolStatus_value[argStatus]; !exist {
				return fmt.Errorf("status not exist")
			}
			status := types.PoolStatus_value[argStatus]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetPoolStatus(
				clientCtx.GetFromAddress().String(),
				argDenom,
				argPool,
				types.PoolStatus(status),
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
