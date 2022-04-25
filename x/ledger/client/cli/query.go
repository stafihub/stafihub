package cli

import (
	"fmt"

	"strconv"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/stafihub/stafihub/x/ledger/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group ledger queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdListExchangeRate())
	cmd.AddCommand(CmdShowExchangeRate())
	cmd.AddCommand(CmdShowEraExchangeRate())
	cmd.AddCommand(CmdEraExchangeRatesByDenom())
	cmd.AddCommand(CmdBondedPoolsByDenom())
	cmd.AddCommand(CmdGetPoolDetail())
	cmd.AddCommand(CmdGetChainEra())
	cmd.AddCommand(CmdGetCurrentEraSnapshot())
	cmd.AddCommand(CmdGetProtocolFeeReceiver())
	cmd.AddCommand(CmdGetStakingRewardCommission())
	cmd.AddCommand(CmdGetUnbondRelayFee())
	cmd.AddCommand(CmdGetUnbondCommission())
	cmd.AddCommand(CmdGetEraUnbondLimit())
	cmd.AddCommand(CmdGetBondPipeline())
	cmd.AddCommand(CmdGetEraSnapshot())
	cmd.AddCommand(CmdGetSnapshot())
	cmd.AddCommand(CmdGetTotalExpectedActive())
	cmd.AddCommand(CmdGetPoolUnbond())
	cmd.AddCommand(CmdGetBondRecord())
	cmd.AddCommand(CmdGetSignature())
	cmd.AddCommand(CmdGetRParams())

	cmd.AddCommand(CmdTotalProtocolFee())

	cmd.AddCommand(CmdRelayFeeReceiver())

	cmd.AddCommand(CmdUnbondSwitch())

	// this line is used by starport scaffolding # 1

	return cmd
}

var _ = strconv.Itoa(0)

func CmdBondedPoolsByDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bonded-pools [denom]",
		Short: "Query bonded_pools_by_denom",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDenom := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryBondedPoolsByDenomRequest{
				Denom: reqDenom,
			}

			res, err := queryClient.BondedPoolsByDenom(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetPoolDetail() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pool-detail [denom] [pool]",
		Short: "Query get_pool_detail",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDenom := args[0]
			reqPool := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetPoolDetailRequest{
				Denom: reqDenom,
				Pool:  reqPool,
			}

			res, err := queryClient.GetPoolDetail(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetChainEra() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "chain-era [denom]",
		Short: "Query getChainEra",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDenom := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetChainEraRequest{
				Denom: reqDenom,
			}

			res, err := queryClient.GetChainEra(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetCurrentEraSnapshot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "current-era-snapshot [denom]",
		Short: "Query getCurrentEraSnapshot",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDenom := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetCurrentEraSnapshotRequest{
				Denom: reqDenom,
			}

			res, err := queryClient.GetCurrentEraSnapshot(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetProtocolFeeReceiver() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "protocol-fee-receiver",
		Short: "Query protocol fee receiver",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetProtocolFeeReceiverRequest{}

			res, err := queryClient.GetProtocolFeeReceiver(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetStakingRewardCommission() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "staking-reward-commission [denom]",
		Short: "Query staking reward commission",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetStakingRewardCommissionRequest{
				Denom: args[0],
			}

			res, err := queryClient.GetStakingRewardCommission(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetUnbondRelayFee() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unbond-relay-fee [denom]",
		Short: "Query getUnbondRelayFee",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDenom := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetUnbondRelayFeeRequest{
				Denom: reqDenom,
			}

			res, err := queryClient.GetUnbondRelayFee(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetUnbondCommission() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unbond-commission [denom]",
		Short: "Query getUnbondCommission",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			params := &types.QueryGetUnbondCommissionRequest{
				Denom: args[0],
			}

			res, err := queryClient.GetUnbondCommission(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetEraUnbondLimit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "era-unbond-limit [denom]",
		Short: "Query getEraUnbondLimit",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDenom := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetEraUnbondLimitRequest{

				Denom: reqDenom,
			}

			res, err := queryClient.GetEraUnbondLimit(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetBondPipeline() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bond-pipeline [denom] [pool]",
		Short: "Query GetBondPipeline",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDenom := args[0]
			reqPool := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			params := &types.QueryGetBondPipelineRequest{
				Denom: reqDenom,
				Pool:  reqPool,
			}

			res, err := queryClient.GetBondPipeline(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetEraSnapshot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "era-snapshot [denom] [era]",
		Short: "Query GetEraSnapshot",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDenom := args[0]
			reqEra, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			params := &types.QueryGetEraSnapshotRequest{
				Denom: reqDenom,
				Era:   uint32(reqEra),
			}

			res, err := queryClient.GetEraSnapshot(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetSnapshot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "snapshot [shot-id]",
		Short: "Query GetSnapShot",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqShotId := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			params := &types.QueryGetSnapshotRequest{
				ShotId: reqShotId,
			}

			res, err := queryClient.GetSnapshot(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetTotalExpectedActive() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "total-expected-active [denom] [era]",
		Short: "Query GetTotalExpectedActive",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDenom := args[0]
			reqEra, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetTotalExpectedActiveRequest{
				Denom: reqDenom,
				Era:   uint32(reqEra),
			}

			res, err := queryClient.GetTotalExpectedActive(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetPoolUnbond() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pool-unbond [denom] [pool] [era]",
		Short: "Query GetPoolUnbond",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDenom := args[0]
			reqPool := args[1]
			reqEra, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			params := &types.QueryGetPoolUnbondRequest{
				Denom: reqDenom,
				Pool:  reqPool,
				Era:   uint32(reqEra),
			}

			res, err := queryClient.GetPoolUnbond(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetBondRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bond-record [denom] [txhash]",
		Short: "Query GetBondRecord",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDenom := args[0]
			reqTxhash := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			params := &types.QueryGetBondRecordRequest{
				Denom:  reqDenom,
				Txhash: reqTxhash,
			}

			res, err := queryClient.GetBondRecord(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func CmdGetSignature() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "signature [denom] [era] [pool] [tx-type] [prop-id]",
		Short: "Query GetSignature",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDenom := args[0]
			reqEra, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}
			reqPool := args[2]
			reqTxType, ok := types.OriginalTxType_value[args[3]]
			if !ok {
				return fmt.Errorf("invalid txtype")
			}

			reqPropId := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			params := &types.QueryGetSignatureRequest{
				Denom:  reqDenom,
				Era:    uint32(reqEra),
				Pool:   reqPool,
				TxType: types.OriginalTxType(reqTxType),
				PropId: reqPropId,
			}

			res, err := queryClient.GetSignature(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetRParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "r-params [denom]",
		Short: "query rParams",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDenom := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetRParamsRequest{

				Denom: reqDenom,
			}

			res, err := queryClient.GetRParams(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
