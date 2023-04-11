package cli

import (
	"strconv"
	"strings"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/claim/types"
)

var _ = strconv.Itoa(0)

func CmdClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim [round] [index] [account] [coin] [proof]",
		Short: "Claim",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argRound, err := math.ParseUint(args[0])
			if err != nil {
				return err
			}
			argIndex, err := math.ParseUint(args[1])
			if err != nil {
				return err
			}
			argAccount := args[2]
			argCoin, err := sdk.ParseCoinNormalized(args[3])
			if err != nil {
				return err
			}
			argProof := strings.Split(args[4], ":")

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgClaim(
				clientCtx.GetFromAddress().String(),
				argRound.Uint64(),
				argIndex.Uint64(),
				argAccount,
				argCoin,
				argProof,
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
