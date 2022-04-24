package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/rmintreward/types"
)

var _ = strconv.Itoa(0)

func CmdAddMintRewardAct() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-mint-reward-act [denom] [path-to-act]",
		Short: "Add mint reward act",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Broadcast message add mint reward act which can be given through a JSON file.

Example:
$ %s tx rmintreward add-mint-reward-act uratom path/to/mint_reward_act.json --from admin

Where mint_reward_act.json could be like this:
{
    "begin": 1000,
    "end": 2000,
    "lockedBlocks": 100,
    "tokenRewardInfos": [
        {
            "denom": "ufis",
            "rewardRate": "0.000001",
            "totalRewardAmount": "10000000",
            "userLimit": "100"
        }
    ]
}
`, version.AppName)),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			denom := args[0]
			argPathToAct := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			act, err := parseAct(argPathToAct)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddMintRewardAct(
				clientCtx.GetFromAddress().String(),
				denom,
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

func parseAct(path string) (*types.MintRewardActPost, error) {
	act := types.MintRewardActPost{}
	if path == "" {
		return nil, fmt.Errorf("act json file path not give")
	}

	contents, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(contents, &act)
	if err != nil {
		return nil, err
	}
	return &act, nil
}
