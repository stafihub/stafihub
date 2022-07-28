package keeper

import (
	"context"
	"encoding/hex"
	"fmt"

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

	root, err := hex.DecodeString(msg.MerkleRoot)
	if err != nil {
		return nil, types.ErrMerkleRootFormatNotMatch
	}
	if len(root) != 32 {
		return nil, types.ErrMerkleRootFormatNotMatch
	}

	willUseRound := k.GetClaimRound(ctx) + 1
	k.Keeper.SetMerkleRoot(ctx, willUseRound, root)
	k.Keeper.SetClaimRound(ctx, willUseRound)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSetMerkleRoot,
			sdk.NewAttribute(types.AttributeKeyClaimRound, fmt.Sprint(willUseRound)),
			sdk.NewAttribute(types.AttributeKeyMerkleRoot, msg.MerkleRoot),
		),
	)

	return &types.MsgSetMerkleRootResponse{}, nil
}
