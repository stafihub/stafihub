package ledger

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafiprotocol/stafihub/x/ledger/keeper"
	"github.com/stafiprotocol/stafihub/x/ledger/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgAddNewPool:
					res, err := msgServer.AddNewPool(sdk.WrapSDKContext(ctx), msg)
					return sdk.WrapServiceResult(ctx, res, err)
case *types.MsgRemovePool:
					res, err := msgServer.RemovePool(sdk.WrapSDKContext(ctx), msg)
					return sdk.WrapServiceResult(ctx, res, err)
case *types.MsgSetEraUnbondLimit:
					res, err := msgServer.SetEraUnbondLimit(sdk.WrapSDKContext(ctx), msg)
					return sdk.WrapServiceResult(ctx, res, err)
case *types.MsgSetInitBond:
					res, err := msgServer.SetInitBond(sdk.WrapSDKContext(ctx), msg)
					return sdk.WrapServiceResult(ctx, res, err)
case *types.MsgSetChainBondingDuration:
					res, err := msgServer.SetChainBondingDuration(sdk.WrapSDKContext(ctx), msg)
					return sdk.WrapServiceResult(ctx, res, err)
case *types.MsgSetPoolDetail:
					res, err := msgServer.SetPoolDetail(sdk.WrapSDKContext(ctx), msg)
					return sdk.WrapServiceResult(ctx, res, err)
case *types.MsgSetLeastBond:
					res, err := msgServer.SetLeastBond(sdk.WrapSDKContext(ctx), msg)
					return sdk.WrapServiceResult(ctx, res, err)
case *types.MsgClearCurrentEraSnapShots:
					res, err := msgServer.ClearCurrentEraSnapShots(sdk.WrapSDKContext(ctx), msg)
					return sdk.WrapServiceResult(ctx, res, err)
case *types.MsgSetChainEra:
					res, err := msgServer.SetChainEra(sdk.WrapSDKContext(ctx), msg)
					return sdk.WrapServiceResult(ctx, res, err)
// this line is used by starport scaffolding # 1
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
