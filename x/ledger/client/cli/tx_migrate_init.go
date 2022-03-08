package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/stafihub/stafihub/utils"
	"github.com/stafihub/stafihub/x/ledger/types"
)

var _ = strconv.Itoa(0)
var FlagUnbondings = "unbondings"

func CmdMigrateInit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate-init [denom] [total-supply] [exchange-rate]",
		Short: "Migrate init",

		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argTotalSupply, ok := sdk.NewIntFromString(args[1])
			if !ok {
				return fmt.Errorf("argDenom format err")
			}
			argExchangeRate, err := utils.NewDecFromStr(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgMigrateInit(
				clientCtx.GetFromAddress().String(),
				argDenom,
				argTotalSupply,
				argExchangeRate,
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

func parseUnbondingsFlags(fs *pflag.FlagSet) ([]*types.PoolUnbond, error) {
	ud := make([]*types.PoolUnbond, 0)
	udFile, err := fs.GetString(FlagUnbondings)
	if err != nil {
		return nil, err
	}

	if udFile == "" {
		return nil, fmt.Errorf("unbondings json file not give")
	}

	contents, err := os.ReadFile(udFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &ud)
	if err != nil {
		return nil, err
	}

	return ud, nil
}
