package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/bridge/types"
)

var _ = strconv.Itoa(0)

func CmdProposalDetail() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proposal-detail [chain-id] [deposit-nonce] [resource-id] [amount] [receiver]",
		Short: "Query proposal detail",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqChainId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			reqDepositNonce, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}
			reqResourceId := args[2]
			reqAmount := args[3]
			reqReceiver := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryProposalDetailRequest{

				ChainId:      uint32(reqChainId),
				DepositNonce: reqDepositNonce,
				ResourceId:   reqResourceId,
				Amount:       reqAmount,
				Receiver:     reqReceiver,
			}

			res, err := queryClient.ProposalDetail(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
