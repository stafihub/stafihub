package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/mining/types"
)

var _ = strconv.Itoa(0)

func CmdAddStakePool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-stake-pool [stake-token-denom] [path-to-add_stake_pool.json]",
		Short: "Add stake pool",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argStakeTokenDenom := args[0]

			createInfo, err := parseCreateInfo(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddStakePool(
				clientCtx.GetFromAddress().String(),
				argStakeTokenDenom,
				createInfo.RewardPoolList,
				createInfo.StakeItemList,
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

type CreateInfo struct {
	RewardPoolList []*types.CreateRewardPoolInfo `json:"rewardPoolList"`
	StakeItemList  []*types.CreateStakeItemInfo  `json:"stakeItemList"`
}

func parseCreateInfo(path string) (*CreateInfo, error) {
	var createInfo *CreateInfo
	if path == "" {
		return nil, fmt.Errorf("reward pool list json file path not give")
	}

	contents, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(contents, createInfo)
	if err != nil {
		return nil, err
	}
	return createInfo, nil
}
