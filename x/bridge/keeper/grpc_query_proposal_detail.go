package keeper

import (
	"context"
	"encoding/hex"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/bridge/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ProposalDetail(goCtx context.Context, req *types.QueryProposalDetailRequest) (*types.QueryProposalDetailResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	resourceIdSlice, err := hex.DecodeString(req.ResourceId)
	if err != nil {
		return nil, err
	}
	var resourceId [32]byte
	copy(resourceId[:], resourceIdSlice)
	amount, ok := sdk.NewIntFromString(req.Amount)
	if !ok {
		return nil, fmt.Errorf("amount format err")
	}

	chainId := uint8(req.ChainId)
	proposal, found := k.GetProposal(ctx, chainId, req.DepositNonce, resourceId, types.ProposalContent{
		Amount:   amount,
		Receiver: req.Receiver,
	})
	if !found {
		return nil, status.Error(codes.NotFound, codes.NotFound.String())
	}

	return &types.QueryProposalDetailResponse{
		Proposal: proposal,
	}, nil
}
