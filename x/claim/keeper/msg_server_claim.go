package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/claim/types"
)

func (k msgServer) Claim(goCtx context.Context, msg *types.MsgClaim) (*types.MsgClaimResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if k.IsIndexClaimed(ctx, msg.Round, msg.Index) {
		return nil, types.ErrAlreadyClaimed
	}

	proof := make([]NodeHash, len(msg.Proof))
	for i, p := range msg.Proof {
		nodeHash, err := NodeHashFromHexString(p)
		if err != nil {
			return nil, types.ErrNodeHashFormatNotMatch
		}
		proof[i] = nodeHash
	}
	account, err := sdk.AccAddressFromBech32(msg.Account)
	if err != nil {
		return nil, types.ErrAccountFormatNotMatch
	}

	rootNode, found := k.Keeper.GetMerkleRoot(ctx, msg.Round)
	if !found {
		return nil, types.ErrMerkleRootNotExist
	}
	userNode := GetNodeHash(msg.Round, msg.Index, account, msg.Coin)
	if !VerifyProof(userNode, proof, rootNode) {
		return nil, types.ErrMerkleProofNotMatch
	}

	if msg.Coin.IsPositive() {
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, account, sdk.NewCoins(msg.Coin)); err != nil {
			return nil, err
		}
	}

	k.Keeper.SetIndexClaimed(ctx, msg.Round, msg.Index)

	return &types.MsgClaimResponse{}, nil
}
