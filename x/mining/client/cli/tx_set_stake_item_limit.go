package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/utils"
	"github.com/stafihub/stafihub/x/mining/types"
)

var _ = strconv.Itoa(0)

func CmdSetStakeItemLimit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-stake-item-limit [max-lock-second] [max-power-reward-rate]",
		Short: "Set stake item limit",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argMaxLockSecond, err := sdk.ParseUint(args[0])
			if err != nil {
				return err
			}

			argMaxPowerRewardRate := utils.MustNewDecFromStr(args[1])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetStakeItemLimit(
				clientCtx.GetFromAddress().String(),
				argMaxLockSecond.Uint64(),
				argMaxPowerRewardRate,
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
