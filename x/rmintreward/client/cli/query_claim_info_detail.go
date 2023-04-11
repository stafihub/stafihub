package cli

import (
	"strconv"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/rmintreward/types"
)

var _ = strconv.Itoa(0)

func CmdClaimInfoDetail() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim-info-detail [address] [denom] [cycle] [mint-index]",
		Short: "Query user claim info detail",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqAddress := args[0]
			reqDenom := args[1]
			reqCycle, err := math.ParseUint(args[2])
			if err != nil {
				return err
			}
			reqMintIndex, err := math.ParseUint(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryClaimInfoDetailRequest{
				Address:   reqAddress,
				Denom:     reqDenom,
				Cycle:     reqCycle.Uint64(),
				MintIndex: reqMintIndex.Uint64(),
			}

			res, err := queryClient.ClaimInfoDetail(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
