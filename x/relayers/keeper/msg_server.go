package keeper

import (
	"context"
	"strconv"

	"github.com/stafiprotocol/stafihub/x/relayers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) CreateRelayer(goCtx context.Context,  msg *types.MsgCreateRelayer) (*types.MsgCreateRelayerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	relayer := types.Relayer{Denom: msg.Denom, Address: msg.Address}
	k.SetRelayer(ctx, relayer)

	return &types.MsgCreateRelayerResponse{}, nil
}

func (k msgServer) DeleteRelayer(goCtx context.Context,  msg *types.MsgDeleteRelayer) (*types.MsgDeleteRelayerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.RemoveRelayer(ctx, msg.Denom, msg.Address)

	return &types.MsgDeleteRelayerResponse{}, nil
}

func (k msgServer) SetThreshold(goCtx context.Context,  msg *types.MsgSetThreshold) (*types.MsgSetThresholdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)


	value, err := strconv.ParseUint(msg.Value, 10, 64)
	if err != nil {
		return nil, err
	}

	threshold := types.Threshold{Denom: msg.Denom, Value: uint32(value)}
	k.UpdateThreshold(ctx, threshold)

	return &types.MsgSetThresholdResponse{}, nil
}
