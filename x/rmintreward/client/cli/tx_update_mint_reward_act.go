package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/rmintreward/types"
)

var _ = strconv.Itoa(0)

func CmdUpdateMintRewardAct() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-mint-reward-act [denom] [cycle] [path-to-act]",
		Short: "Update mint reward act",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Broadcast message add mint reward act which can be given through a JSON file.

Example:
$ %s tx rmintreward update-mint-reward-act uratom cycle path/to/unbondings.json --from admin

Where act.json could be like this:
{
    "begin": 1000,
    "end": 2000,
    "lockedBlocks": 100,
    "tokenRewardInfos": [
        {
            "denom": "ufis",
            "rewardRate": "0.000001",
            "totalRewardAmount": "10000000",
            "leftAmount": "10000000",
            "userLimit": "100"
        }
    ]
}
`, version.AppName)),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			denom := args[0]
			cycleUint, err := sdk.ParseUint(args[1])
			if err != nil {
				return err
			}
			argPathToAct := args[2]
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			act, err := parseAct(argPathToAct)
			if err != nil {
				return err
			}
			msg := types.NewMsgUpdateMintRewardAct(
				clientCtx.GetFromAddress().String(),
				denom,
				cycleUint.Uint64(),
				act,
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
