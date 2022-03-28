package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/rmintreward/types"
)

var _ = strconv.Itoa(0)

func CmdProvideRewardToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "provide-reward-token [amount]",
		Short: "Provide reward token for mint reward",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAmount, err := sdk.ParseCoinsNormalized(args[0])
			if err != nil {
				return err
			}
			if len(argAmount) == 0 {
				return fmt.Errorf("amount is required")
			}
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgProvideRewardToken(
				clientCtx.GetFromAddress().String(),
				argAmount[0],
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
