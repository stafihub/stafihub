package keeper

import (
	"context"
	"encoding/hex"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/bridge/types"
)

func (k msgServer) Deposit(goCtx context.Context, msg *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	resourceIdSlice, err := hex.DecodeString(msg.ResourceId)
	if err != nil {
		return nil, err
	}
	var resourceId [32]byte
	copy(resourceId[:], resourceIdSlice)

	denom, found := k.Keeper.GetDenomByResourceId(ctx, resourceId)
	if !found {
		return nil, types.ErrResourceIdNotFound
	}
	userAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	chainId := uint8(msg.DestChainId)
	count := k.Keeper.GetDepositCounts(ctx, chainId)

	balance := k.bankKeeper.GetBalance(ctx, userAddress, denom)
	if balance.Amount.LT(msg.Amount) {
		return nil, types.ErrBalanceNotEnough
	}

	burnedCoins := sdk.NewCoins(sdk.NewCoin(denom, msg.Amount))

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, userAddress, types.ModuleName, burnedCoins)
	if err != nil {
		return nil, err
	}
	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, burnedCoins)
	if err != nil {
		return nil, err
	}
	k.Keeper.SetDepositCounts(ctx, chainId, count+1)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeDeposit,
			sdk.NewAttribute(types.AttributeKeyDestChainId, fmt.Sprintf("%d", chainId)),
			sdk.NewAttribute(types.AttributeKeyResourceId, msg.ResourceId),
			sdk.NewAttribute(types.AttributeKeyDepositNonce, fmt.Sprintf("%d", count)),
			sdk.NewAttribute(types.AttributeKeyAmount, msg.Amount.String()),
			sdk.NewAttribute(types.AttributeKeyReceiver, msg.Receiver),
		),
	)

	return &types.MsgDepositResponse{}, nil
}
