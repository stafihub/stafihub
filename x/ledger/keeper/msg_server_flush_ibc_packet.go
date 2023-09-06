package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/ledger/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) FlushIbcPacket(goCtx context.Context, msg *types.MsgFlushIbcPacket) (*types.MsgFlushIbcPacketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	k.IBCKeeper.ChannelKeeper.DeletePacketCommitment(ctx, msg.PortID, msg.ChannelID, uint64(msg.Sequence))

	nextSequenceAck := uint64(msg.Sequence + 1)
	k.IBCKeeper.ChannelKeeper.SetNextSequenceAck(ctx, msg.PortID, msg.ChannelID, nextSequenceAck)

	return &types.MsgFlushIbcPacketResponse{}, nil
}
