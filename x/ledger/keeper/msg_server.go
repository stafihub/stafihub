package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stafiprotocol/stafihub/x/ledger/types"
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

func (k msgServer) LiquidityUnbond(goCtx context.Context, msg *types.MsgLiquidityUnbond) (*types.MsgLiquidityUnbondResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	denom := msg.Value.Denom
	ce, ok := k.Keeper.GetChainEra(ctx, denom)
	if !ok {
		return nil, types.ErrChainEraNotFound
	}

	_, ok = k.Keeper.GetChainBondingDuration(ctx, denom)
	if !ok {
		return nil, types.ErrBondingDurationNotSet
	}

	receiver := k.Keeper.GetReceiver(ctx)
	if receiver == nil {
		return nil, types.ErrNoReceiver
	}

	unbonder, _ := sdk.AccAddressFromBech32(msg.Creator)
	rbalance := k.bankKeeper.GetBalance(ctx, unbonder, denom)
	if rbalance.IsLT(msg.Value) {
		return nil, sdkerrors.ErrInsufficientFunds
	}

	pipe, ok := k.Keeper.GetBondPipeLine(ctx, denom, msg.Pool)
	if !ok {
		pipe = types.NewBondPipeline(denom, msg.Pool)
	}

	cms := k.Keeper.GetUnbondCommission(ctx)
	cmsFee := cms.MulInt(msg.Value.Amount).TruncateInt()
	leftValue := msg.Value.SubAmount(cmsFee)
	balance := k.RtokenToToken(ctx, leftValue.Denom, leftValue.Amount)
	pipe.Chunk.Active = pipe.Chunk.Active.Sub(balance)
	if pipe.Chunk.Active.IsNegative() {
		return nil, sdkerrors.ErrInsufficientFunds
	}
	pipe.Chunk.Unbond = pipe.Chunk.Unbond.Add(balance)

	chunk := types.UserUnlockChunk{Pool: msg.Pool, UnlockEra: ce.Era, Value: balance, Recipient: msg.Recipient}
	unbonds, ok := k.Keeper.GetAccountUnbond(ctx, denom, msg.Creator)
	if !ok {
		unbonds = types.NewAccountUnbond(denom, msg.Creator, []types.UserUnlockChunk{chunk})
	} else {
		unbonds.Chunks = append(unbonds.Chunks, chunk)
	}

	unbonding := types.NewUnbonding(msg.Creator, msg.Recipient, balance)
	poolUnbonds, ok := k.Keeper.GetPoolUnbond(ctx, denom, msg.Pool, ce.Era)
	eul, ok := k.Keeper.GetEraUnbondLimit(ctx, denom)
	if !ok {
		poolUnbonds = types.NewPoolUnbond(denom, msg.Pool, ce.Era, []types.Unbonding{unbonding})
	} else {
		if uint32(len(poolUnbonds.Unbondings)) > eul.Limit {
			return nil, types.ErrPoolLimitReached
		}
		poolUnbonds.Unbondings = append(poolUnbonds.Unbondings, unbonding)
	}

	unbondFee, ok := k.Keeper.GetUnbondFee(ctx, denom)
	if ok && unbondFee.Value.IsPositive() {
		if err := k.bankKeeper.SendCoins(ctx, unbonder, receiver, sdk.Coins{unbondFee.Value}); err != nil {
			return nil, err
		}
	}

	if cmsFee.LT(sdk.ZeroInt()) {
		cmsFeeCoin := sdk.NewCoin(denom, cmsFee)
		if err := k.bankKeeper.SendCoins(ctx, unbonder, receiver, sdk.Coins{cmsFeeCoin}); err != nil {
			panic(err)
		}
	}

	burnCoins := sdk.Coins{leftValue}
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, unbonder, types.ModuleName, burnCoins); err != nil {
		panic(err)
	}

	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, burnCoins); err != nil {
		panic(err)
	}

	k.Keeper.SetBondPipeline(ctx, pipe)
	k.Keeper.SetAccountUnbond(ctx, unbonds)
	k.Keeper.SetPoolUnbond(ctx, poolUnbonds)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeLiquidityUnbond,
			sdk.NewAttribute(types.AttributeKeyDenom, denom),
			sdk.NewAttribute(types.AttributeKeyUnbonder, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyPool, msg.Pool),
			sdk.NewAttribute(types.AttributeKeyUnBondAmount, leftValue.String()),
			sdk.NewAttribute(types.AttributeKeyReceiver, balance.String()),
			sdk.NewAttribute(types.AttributeKeyReceiver, msg.Recipient),
		),
	)

	return &types.MsgLiquidityUnbondResponse{}, nil
}
