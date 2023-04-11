package cli

import (
	"strconv"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/rmintreward/types"
)

var _ = strconv.Itoa(0)

func CmdClaimMintReward() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim-mint-reward [denom] [cycle] [mint-index]",
		Short: "Claim mint reward",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argCycle, err := math.ParseUint(args[1])
			if err != nil {
				return err
			}

			argMintIndex, err := math.ParseUint(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgClaimMintReward(
				clientCtx.GetFromAddress().String(),
				argDenom,
				argCycle.Uint64(),
				argMintIndex.Uint64(),
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
