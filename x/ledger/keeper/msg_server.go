package keeper

import (
	"context"

	"github.com/stafiprotocol/stafihub/x/ledger/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

func (k msgServer) LiquidityUnbond(goCtx context.Context,  msg *types.MsgLiquidityUnbond) (*types.MsgLiquidityUnbondResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	ce, ok := k.Keeper.GetChainEra(ctx, msg.Denom)
	if !ok {
		return nil, types.ErrChainEraNotFound
	}

	_, ok = k.Keeper.GetChainBondingDuration(ctx, msg.Denom)
	if !ok {
		return nil, types.ErrBondingDurationNotSet
	}

	receiver := k.Keeper.GetReceiver(ctx)
	if receiver == nil {
		return nil, types.ErrNoReceiver
	}

	unbonder, _ := sdk.AccAddressFromBech32(msg.Creator)
	rbalance := k.bankKeeper.GetBalance(ctx, unbonder, msg.Denom)
	if rbalance.Amount.LT(msg.Value) {
		return nil, sdkerrors.ErrInsufficientFunds
	}

	ucms := k.Keeper.GetUnbondCommission(ctx)
	cmsFee := ucms.MulInt(msg.Value).TruncateInt()
	leftValue := msg.Value.Sub(cmsFee)
	if leftValue.LTE(sdk.ZeroInt()) {
		return nil, sdkerrors.ErrInsufficientFunds
	}
	balance := k.RtokenToToken(ctx, msg.Denom, leftValue)

	pipe, ok := k.Keeper.GetBondPipeLine(ctx, msg.Denom, msg.Pool)
	if !ok {
		pipe = types.NewBondPipeline(msg.Denom, msg.Pool)
	}
	pipe.Chunk.Active = pipe.Chunk.Active.Sub(balance)
	if pipe.Chunk.Active.LT(sdk.ZeroInt()) {
		return nil, sdkerrors.ErrInsufficientFunds
	}
	pipe.Chunk.Unbond = pipe.Chunk.Unbond.Add(balance)

	chunk := types.UserUnlockChunk{Pool: msg.Pool, UnlockEra: ce.Era, Value: balance, Recipient: msg.Recipient}
	unbonds, ok := k.Keeper.GetAccountUnbond(ctx, msg.Denom, msg.Creator)
	if !ok {
		unbonds = types.NewAccountUnbond(msg.Denom, msg.Creator, []types.UserUnlockChunk{chunk})
	} else {
		unbonds.Chunks = append(unbonds.Chunks, chunk)
	}

	unbonding := types.NewUnbonding(msg.Creator, msg.Recipient, balance)
	poolUnbonds, ok := k.Keeper.GetPoolUnbond(ctx, msg.Denom, msg.Pool, ce.Era)
	eul, ok := k.Keeper.GetEraUnbondLimit(ctx, msg.Denom)
	if !ok {
		poolUnbonds = types.NewPoolUnbond(msg.Denom, msg.Pool, ce.Era, []types.Unbonding{unbonding})
	} else {
		if uint32(len(poolUnbonds.Unbondings)) > eul.Limit {
			return nil, types.ErrPoolLimitReached
		}
		poolUnbonds.Unbondings = append(poolUnbonds.Unbondings, unbonding)
	}

	cacheCtx, writeCache := ctx.CacheContext()
	unbondFee, ok := k.Keeper.GetUnbondFee(ctx, msg.Denom)
	if ok && unbondFee.Value.IsPositive() {
		if err := k.bankKeeper.SendCoins(cacheCtx, unbonder, receiver, sdk.Coins{unbondFee.Value}); err != nil {
			return nil, err
		}
	}

	if cmsFee.LT(sdk.ZeroInt()) {
		cmsFeeCoin := sdk.NewCoin(msg.Denom, cmsFee)
		if err := k.bankKeeper.SendCoins(cacheCtx, unbonder, receiver, sdk.Coins{cmsFeeCoin}); err != nil {
			return nil, err
		}
	}

	burnCoins := sdk.Coins{sdk.NewCoin(msg.Denom, leftValue)}
	if err := k.bankKeeper.SendCoinsFromAccountToModule(cacheCtx, unbonder, types.ModuleName, burnCoins); err != nil {
		return nil, err
	}

	if err := k.bankKeeper.BurnCoins(cacheCtx, types.ModuleName, burnCoins); err != nil {
		return nil, err
	}

	k.Keeper.SetBondPipeline(ctx, pipe)
	k.Keeper.SetAccountUnbond(ctx, unbonds)
	k.Keeper.SetPoolUnbond(ctx, poolUnbonds)

	// write state to the underlying multi-store
	writeCache()

	//todo add event and check cacheCtx

	return &types.MsgLiquidityUnbondResponse{}, nil
}
