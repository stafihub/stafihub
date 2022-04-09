package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rbank/types"
	sudoTypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) AddDenom(goCtx context.Context, msg *types.MsgAddDenom) (*types.MsgAddDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	k.SetAddressPrefix(ctx, msg.Metadata.Base, msg.AddressPrefix)
	k.bankKeeper.SetDenomMetaData(ctx, msg.Metadata)

	return &types.MsgAddDenomResponse{}, nil
}
