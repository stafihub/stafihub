package keeper

import (
	"context"
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/claim/types"
	sudoTypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) SetMerkleRoot(goCtx context.Context, msg *types.MsgSetMerkleRoot) (*types.MsgSetMerkleRootResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	bts, err := hex.DecodeString(msg.MerkleRoot)
	if err != nil {
		return nil, types.ErrMerkleRootFormat
	}
	if len(bts) != 32 {
		return nil, types.ErrMerkleRootFormat
	}

	var root [32]byte
	copy(root[:], bts)

	willUseRound := k.GetClaimRound(ctx) + 1
	k.Keeper.SetMerkleRoot(ctx, willUseRound, root)
	k.Keeper.SetClaimRound(ctx, willUseRound)

	return &types.MsgSetMerkleRootResponse{}, nil
}
