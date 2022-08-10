package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rvalidator/types"
	sudoTypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) RmRValidator(goCtx context.Context, msg *types.MsgRmRValidator) (*types.MsgRmRValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}
	newVal := types.RValidator{
		Denom:       msg.Denom,
		PoolAddress: msg.PoolAddress,
		ValAddress:  msg.NewAddress,
	}
	if !k.Keeper.HasSelectedRValidator(ctx, &newVal) {
		return nil, types.ErrRValidatorNotExist
	}

	latestVotedCycle := k.GetLatestVotedCycle(ctx, msg.Denom, msg.PoolAddress)
	willUseCycle := types.Cycle{
		Denom:       msg.Denom,
		PoolAddress: msg.PoolAddress,
		Version:     latestVotedCycle.Version,
		Number:      latestVotedCycle.Number + 1,
	}

	proposal := types.NewUpdateRValidatorProposal(msg.Creator, msg.Denom, msg.PoolAddress, msg.OldAddress, msg.NewAddress, &willUseCycle)

	err := k.ProcessUpdateRValidatorProposal(ctx, proposal)
	if err != nil {
		return nil, err
	}
	return &types.MsgRmRValidatorResponse{}, nil
}
