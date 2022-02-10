package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/ledger/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

func (q Querier) GetCurrentEraSnapshot(goCtx context.Context, req *types.QueryGetCurrentEraSnapshotRequest) (*types.QueryGetCurrentEraSnapshotResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	shot := q.Keeper.CurrentEraSnapshots(ctx, req.Denom)

	return &types.QueryGetCurrentEraSnapshotResponse{ShotIds: shot.ShotIds}, nil
}

func (q Querier) PoolsByDenom(goCtx context.Context, req *types.QueryPoolsByDenomRequest) (*types.QueryPoolsByDenomResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	pool, found := q.GetPool(ctx, req.Denom)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryPoolsByDenomResponse{Addrs: pool.Addrs}, nil
}

func (q Querier) BondedPoolsByDenom(goCtx context.Context, req *types.QueryBondedPoolsByDenomRequest) (*types.QueryBondedPoolsByDenomResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	bondedPool, found := q.GetBondedPool(ctx, req.Denom)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryBondedPoolsByDenomResponse{Addrs: bondedPool.Addrs}, nil
}

func (q Querier) GetPoolDetail(goCtx context.Context, req *types.QueryGetPoolDetailRequest) (*types.QueryGetPoolDetailResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	detail, found := q.Keeper.GetPoolDetail(ctx, req.Denom, req.Pool)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetPoolDetailResponse{Detail: detail}, nil
}

func (q Querier) GetChainEra(goCtx context.Context, req *types.QueryGetChainEraRequest) (*types.QueryGetChainEraResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	ce, found := q.Keeper.GetChainEra(ctx, req.Denom)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetChainEraResponse{Era: ce.Era}, nil
}

func (q Querier) GetChainBondingDuration(goCtx context.Context, req *types.QueryGetChainBondingDurationRequest) (*types.QueryGetChainBondingDurationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	cbd, found := q.Keeper.GetChainBondingDuration(ctx, req.Denom)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetChainBondingDurationResponse{Era: cbd.Era}, nil
}

func (q Querier) GetReceiver(goCtx context.Context, req *types.QueryGetReceiverRequest) (*types.QueryGetReceiverResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	rec := q.Keeper.GetReceiver(ctx)
	if rec == nil {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetReceiverResponse{Receiver: rec.String()}, nil
}

func (q Querier) GetCommission(goCtx context.Context, req *types.QueryGetCommissionRequest) (*types.QueryGetCommissionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	cms := q.Keeper.Commission(ctx)

	return &types.QueryGetCommissionResponse{Commission: cms.String()}, nil
}

func (q Querier) GetUnbondFee(goCtx context.Context, req *types.QueryGetUnbondFeeRequest) (*types.QueryGetUnbondFeeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	fee, found := q.Keeper.GetUnbondFee(ctx, req.Denom)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetUnbondFeeResponse{Fee: fee}, nil
}

func (q Querier) GetUnbondCommission(goCtx context.Context, req *types.QueryGetUnbondCommissionRequest) (*types.QueryGetUnbondCommissionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	cms := q.Keeper.GetUnbondCommission(ctx)

	return &types.QueryGetUnbondCommissionResponse{Commission: cms.String()}, nil
}

func (q Querier) GetLeastBond(goCtx context.Context, req *types.QueryGetLeastBondRequest) (*types.QueryGetLeastBondResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	lb, found := q.Keeper.LeastBond(ctx, req.Denom)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetLeastBondResponse{LeastBond: &lb}, nil
}

func (q Querier) GetEraUnbondLimit(goCtx context.Context, req *types.QueryGetEraUnbondLimitRequest) (*types.QueryGetEraUnbondLimitResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	eul, found := q.Keeper.GetEraUnbondLimit(ctx, req.Denom)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetEraUnbondLimitResponse{Limit: eul.Limit}, nil
}

func (q Querier) GetBondPipeline(goCtx context.Context, req *types.QueryGetBondPipelineRequest) (*types.QueryGetBondPipelineResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	pipe, _ := q.Keeper.GetBondPipeline(ctx, req.Denom, req.Pool)
	return &types.QueryGetBondPipelineResponse{Pipeline: pipe}, nil
}

func (q Querier) GetEraSnapshot(goCtx context.Context, req *types.QueryGetEraSnapshotRequest) (*types.QueryGetEraSnapshotResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	shot := q.Keeper.EraSnapshot(ctx, req.Denom, req.Era)

	return &types.QueryGetEraSnapshotResponse{ShotIds: shot.ShotIds}, nil
}

func (q Querier) GetSnapshot(goCtx context.Context, req *types.QueryGetSnapshotRequest) (*types.QueryGetSnapshotResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	shot, found := q.Keeper.Snapshot(ctx, req.ShotId)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetSnapshotResponse{Shot: shot}, nil
}

func (q Querier) GetTotalExpectedActive(goCtx context.Context, req *types.QueryGetTotalExpectedActiveRequest) (*types.QueryGetTotalExpectedActiveResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	active := q.Keeper.TotalExpectedActive(ctx, req.Denom, req.Era)

	return &types.QueryGetTotalExpectedActiveResponse{Active: active}, nil
}

func (q Querier) GetPoolUnbond(goCtx context.Context, req *types.QueryGetPoolUnbondRequest) (*types.QueryGetPoolUnbondResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	unbond, found := q.Keeper.GetPoolUnbond(ctx, req.Denom, req.Pool, req.Era)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetPoolUnbondResponse{Unbond: unbond}, nil
}

func (q Querier) GetAccountUnbond(goCtx context.Context, req *types.QueryGetAccountUnbondRequest) (*types.QueryGetAccountUnbondResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	unbond, found := q.Keeper.GetAccountUnbond(ctx, req.Denom, req.Unbonder)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetAccountUnbondResponse{Unbond: unbond}, nil
}

func (q Querier) GetBondRecord(goCtx context.Context, req *types.QueryGetBondRecordRequest) (*types.QueryGetBondRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	record, found := q.Keeper.GetBondRecord(ctx, req.Denom, req.Blockhash, req.Txhash)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetBondRecordResponse{BondRecord: record}, nil
}

func (q Querier) GetSignature(goCtx context.Context, req *types.QueryGetSignatureRequest) (*types.QueryGetSignatureResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	sig, found := q.Keeper.GetSignature(ctx, req.Denom, req.Era, req.Pool, req.TxType, req.PropId)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetSignatureResponse{Signature: sig}, nil
}

func (q Querier) GetRParams(goCtx context.Context, req *types.QueryGetRParamsRequest) (*types.QueryGetRParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	rParams, found := q.Keeper.GetRParams(ctx, req.GetDenom())
	if !found {
		return nil, status.Error(codes.NotFound, codes.NotFound.String())
	}

	return &types.QueryGetRParamsResponse{RParams: rParams}, nil
}
