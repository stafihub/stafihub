package keeper

import (
	"context"
	"encoding/hex"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/bridge/types"
)

func (k msgServer) VoteProposal(goCtx context.Context, msg *types.MsgVoteProposal) (*types.MsgVoteProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	resourceIdBts, err := hex.DecodeString(msg.ResourceId)
	if err != nil || len(resourceIdBts) != 32 {
		return nil, types.ErrResourceIdFormatNotRight
	}

	var resourceId [32]byte
	copy(resourceId[:], resourceIdBts)

	resourceIdToDenom, found := k.Keeper.GetResourceIdToDenomByResourceId(ctx, msg.ResourceId)
	if !found {
		return nil, types.ErrResourceIdNotFound
	}
	receiver, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return nil, types.ErrReceiverFormatNotRight
	}

	if msg.Amount.LTE(sdk.ZeroInt()) {
		return nil, types.ErrVoteProposalAmountZero
	}

	shouldMintOrUnlockCoins := sdk.NewCoins(sdk.NewCoin(resourceIdToDenom.Denom, msg.Amount))

	content := types.ProposalContent{
		Amount:   msg.Amount,
		Receiver: msg.Receiver,
	}

	chainId := uint8(msg.ChainId)
	chainIdStr := fmt.Sprintf("%d", chainId)

	if !k.Keeper.HasChainId(ctx, chainId) {
		return nil, types.ErrChainIdNotSupport
	}
	hasRelayer := k.Keeper.relayersKeeper.HasRelayer(ctx, types.ModuleName, chainIdStr, msg.Creator)
	if !hasRelayer {
		return nil, types.ErrRelayerNotExist
	}
	proposal, found := k.Keeper.GetProposal(ctx, chainId, msg.DepositNonce, resourceId, content)
	if !found {
		proposal = &types.Proposal{
			Content: &content,
		}
	}

	if proposal.Executed {
		return nil, types.ErrAlreadyExecuted
	}

	for _, voter := range proposal.Voters {
		if msg.Creator == voter {
			return nil, types.ErrAlreadyVoted
		}
	}
	proposal.Voters = append(proposal.Voters, msg.Creator)
	threshold, found := k.Keeper.relayersKeeper.GetThreshold(ctx, types.ModuleName, chainIdStr)
	if !found {
		return nil, types.ErrThresholdNotSet
	}
	if len(proposal.Voters) >= int(threshold.Value) {
		switch resourceIdToDenom.DenomType {
		case types.External:
			err := k.bankKeeper.MintCoins(ctx, types.ModuleName, shouldMintOrUnlockCoins)
			if err != nil {
				return nil, err
			}
		case types.Native, types.InNativeOutExternal:
		default:
			return nil, types.ErrDenomTypeUnmatch
		}

		err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiver, shouldMintOrUnlockCoins)
		if err != nil {
			return nil, err
		}
		proposal.Executed = true
	}
	k.Keeper.SetProposal(ctx, chainId, msg.DepositNonce, resourceId, proposal)

	return &types.MsgVoteProposalResponse{}, nil
}
