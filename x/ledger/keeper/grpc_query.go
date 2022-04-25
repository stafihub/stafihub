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
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	shot := q.Keeper.CurrentEraSnapshots(ctx, req.Denom)

	return &types.QueryGetCurrentEraSnapshotResponse{ShotIds: shot.ShotIds}, nil
}

func (q Querier) BondedPoolsByDenom(goCtx context.Context, req *types.QueryBondedPoolsByDenomRequest) (*types.QueryBondedPoolsByDenomResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	bondedPool, found := q.GetBondedPool(ctx, req.Denom)
	if !found {
		return nil, status.Error(codes.NotFound, codes.NotFound.String())
	}

	return &types.QueryBondedPoolsByDenomResponse{Addrs: bondedPool.Addrs}, nil
}

func (q Querier) GetPoolDetail(goCtx context.Context, req *types.QueryGetPoolDetailRequest) (*types.QueryGetPoolDetailResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	detail, found := q.Keeper.GetPoolDetail(ctx, req.Denom, req.Pool)
	if !found {
		return nil, status.Error(codes.NotFound, codes.NotFound.String())
	}

	return &types.QueryGetPoolDetailResponse{Detail: detail}, nil
}

func (q Querier) GetChainEra(goCtx context.Context, req *types.QueryGetChainEraRequest) (*types.QueryGetChainEraResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	ce, found := q.Keeper.GetChainEra(ctx, req.Denom)
	if !found {
		return nil, status.Error(codes.NotFound, codes.NotFound.String())
	}

	return &types.QueryGetChainEraResponse{Era: ce.Era}, nil
}

func (q Querier) GetProtocolFeeReceiver(goCtx context.Context, req *types.QueryGetProtocolFeeReceiverRequest) (*types.QueryGetProtocolFeeReceiverResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	rec, found := q.Keeper.GetProtocolFeeReceiver(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, codes.NotFound.String())
	}

	return &types.QueryGetProtocolFeeReceiverResponse{Receiver: rec.String()}, nil
}

func (q Querier) GetStakingRewardCommission(goCtx context.Context, req *types.QueryGetStakingRewardCommissionRequest) (*types.QueryGetStakingRewardCommissionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	cms := q.Keeper.GetStakingRewardCommission(ctx, req.Denom)

	return &types.QueryGetStakingRewardCommissionResponse{Commission: cms.String()}, nil
}

func (q Querier) GetUnbondRelayFee(goCtx context.Context, req *types.QueryGetUnbondRelayFeeRequest) (*types.QueryGetUnbondRelayFeeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	fee := q.Keeper.GetUnbondRelayFee(ctx, req.Denom)

	return &types.QueryGetUnbondRelayFeeResponse{Fee: fee}, nil
}

func (q Querier) GetUnbondCommission(goCtx context.Context, req *types.QueryGetUnbondCommissionRequest) (*types.QueryGetUnbondCommissionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	cms := q.Keeper.GetUnbondCommission(ctx, req.Denom)

	return &types.QueryGetUnbondCommissionResponse{Commission: cms.String()}, nil
}

func (q Querier) GetEraUnbondLimit(goCtx context.Context, req *types.QueryGetEraUnbondLimitRequest) (*types.QueryGetEraUnbondLimitResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	eul := q.Keeper.GetEraUnbondLimit(ctx, req.Denom)

	return &types.QueryGetEraUnbondLimitResponse{Limit: eul.Limit}, nil
}

func (q Querier) GetBondPipeline(goCtx context.Context, req *types.QueryGetBondPipelineRequest) (*types.QueryGetBondPipelineResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	pipe, _ := q.Keeper.GetBondPipeline(ctx, req.Denom, req.Pool)
	return &types.QueryGetBondPipelineResponse{Pipeline: pipe}, nil
}

func (q Querier) GetEraSnapshot(goCtx context.Context, req *types.QueryGetEraSnapshotRequest) (*types.QueryGetEraSnapshotResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	shot := q.Keeper.EraSnapshot(ctx, req.Denom, req.Era)

	return &types.QueryGetEraSnapshotResponse{ShotIds: shot.ShotIds}, nil
}

func (q Querier) GetSnapshot(goCtx context.Context, req *types.QueryGetSnapshotRequest) (*types.QueryGetSnapshotResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	shot, found := q.Keeper.Snapshot(ctx, req.ShotId)
	if !found {
		return nil, status.Error(codes.NotFound, codes.NotFound.String())
	}

	return &types.QueryGetSnapshotResponse{Shot: shot}, nil
}

func (q Querier) GetTotalExpectedActive(goCtx context.Context, req *types.QueryGetTotalExpectedActiveRequest) (*types.QueryGetTotalExpectedActiveResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	active := q.Keeper.TotalExpectedActive(ctx, req.Denom, req.Era)

	return &types.QueryGetTotalExpectedActiveResponse{Active: active}, nil
}

func (q Querier) GetPoolUnbond(goCtx context.Context, req *types.QueryGetPoolUnbondRequest) (*types.QueryGetPoolUnbondResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	unbond, found := q.Keeper.GetPoolUnbond(ctx, req.Denom, req.Pool, req.Era)
	if !found {
		return nil, status.Error(codes.NotFound, codes.NotFound.String())
	}

	return &types.QueryGetPoolUnbondResponse{Unbond: unbond}, nil
}

func (q Querier) GetBondRecord(goCtx context.Context, req *types.QueryGetBondRecordRequest) (*types.QueryGetBondRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	record, found := q.Keeper.GetBondRecord(ctx, req.Denom, req.Txhash)
	if !found {
		return nil, status.Error(codes.NotFound, codes.NotFound.String())
	}

	return &types.QueryGetBondRecordResponse{BondRecord: record}, nil
}

func (q Querier) GetSignature(goCtx context.Context, req *types.QueryGetSignatureRequest) (*types.QueryGetSignatureResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	sig, found := q.Keeper.GetSignature(ctx, req.Denom, req.Era, req.Pool, req.TxType, req.PropId)
	if !found {
		return nil, status.Error(codes.NotFound, codes.NotFound.String())
	}

	return &types.QueryGetSignatureResponse{Signature: sig}, nil
}

func (q Querier) GetRParams(goCtx context.Context, req *types.QueryGetRParamsRequest) (*types.QueryGetRParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	rParams, found := q.Keeper.GetRParams(ctx, req.GetDenom())
	if !found {
		return nil, status.Error(codes.NotFound, codes.NotFound.String())
	}

	return &types.QueryGetRParamsResponse{RParams: rParams}, nil
}
