package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/ledger/types"
	rvotetypes "github.com/stafihub/stafihub/x/rvote/types"
)

var _ = strconv.Itoa(0)

func CmdSetChainEra() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-chain-era [denom] [era]",
		Short: "Broadcast message set_chain_era",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argEra, err := strconv.ParseUint(args[1], 10, 32)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()
			content := types.NewSetChainEraProposal(from, argDenom, uint32(argEra))
			msg, err := rvotetypes.NewMsgSubmitProposal(from, content)
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

func CmdBondReport() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bond-report [denom] [shot-id] [action]",
		Short: "Broadcast message bond_report",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argShotId := args[1]
			argAction, ok := types.BondAction_value[args[2]]
			if !ok {
				return fmt.Errorf("cannot cast %s into bondAction", args[2])
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()
			content := types.NewBondReportProposal(from, argDenom, argShotId, types.BondAction(argAction))
			msg, err := rvotetypes.NewMsgSubmitProposal(from, content)
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

func CmdActiveReport() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "active-report [denom] [shot-id] [staked] [unstaked]",
		Short: "Broadcast message active_report",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argShotId := args[1]
			argStaked, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return fmt.Errorf("cast staked %s into Int error", args[2])
			}
			argUnstaked, ok := sdk.NewIntFromString(args[3])
			if !ok {
				return fmt.Errorf("cast unstaked %s into Int error", args[3])
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()
			content := types.NewActiveReportProposal(from, argDenom, argShotId, argStaked, argUnstaked)
			msg, err := rvotetypes.NewMsgSubmitProposal(from, content)
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

func CmdTransferReport() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-report [denom] [shot-id]",
		Short: "Broadcast message transfer_report",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argShotId := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()
			content := types.NewTransferReportProposal(from, argDenom, argShotId)
			msg, err := rvotetypes.NewMsgSubmitProposal(from, content)
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

func CmdExecuteBondProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "execute-bond-proposal [denom] [bonder] [pool] [txhash] [amount] [state]",
		Short: "Broadcast message execute_bond_proposal",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argBonder, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}
			argPool := args[2]
			argTxHash := args[3]
			argAmount, ok := sdk.NewIntFromString(args[4])
			if !ok {
				return fmt.Errorf("cast amount %s into Int error", args[4])
			}
			bondState, exist := types.LiquidityBondState_value[args[5]]
			if !exist {
				return fmt.Errorf("liquidityBondSate arg not found")
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()
			content := types.NewExecuteBondProposal(from, argDenom, argBonder, argPool, argTxHash, argAmount, types.LiquidityBondState(bondState))
			msg, err := rvotetypes.NewMsgSubmitProposal(from, content)
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

func CmdInterchainTxProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "interchain-tx-proposal [denom] [pool] [era] [txType] [factor] [path_to_msg.json]",
		Short: "Broadcast message interchain tx proposal",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argPool := args[1]
			argEra, err := sdk.ParseUint(args[2])
			if err != nil {
				return err
			}

			txType := types.OriginalTxType(types.OriginalTxType_value[args[3]])
			argFactor, err := sdk.ParseUint(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// check for file path if JSON input is not provided
			contents, err := ioutil.ReadFile(args[4])
			if err != nil {
				return errors.Wrap(err, "neither JSON input nor path to .json file for sdk msg were provided")
			}
			var msgs []interface{}

			err = json.Unmarshal(contents, &msgs)
			if err != nil {
				return err
			}
			cdc := codec.NewProtoCodec(clientCtx.InterfaceRegistry)

			txMsgs := make([]sdk.Msg, 0)

			for _, msg := range msgs {
				content, err := json.Marshal(msg)
				if err != nil {
					return err
				}
				var txMsg sdk.Msg
				if err := cdc.UnmarshalInterfaceJSON(content, &txMsg); err != nil {
					return errors.Wrap(err, "error unmarshalling sdk msg file")
				}
				txMsgs = append(txMsgs, txMsg)
			}
			fmt.Println(txMsgs)

			from := clientCtx.GetFromAddress()

			content, err := types.NewInterchainTxProposal(from, argDenom, argPool, uint32(argEra.Uint64()), txType, uint32(argFactor.Uint64()), txMsgs)
			if err != nil {
				return err
			}

			msg, err := rvotetypes.NewMsgSubmitProposal(from, content)
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
