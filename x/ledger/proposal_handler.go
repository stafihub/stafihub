package ledger

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stafihub/stafihub/x/ledger/keeper"
	"github.com/stafihub/stafihub/x/ledger/types"
	rvotetypes "github.com/stafihub/stafihub/x/rvote/types"
)

// NewParamChangeProposalHandler creates a new governance Handler for a ParamChangeProposal
func NewProposalHandler(k keeper.Keeper) rvotetypes.Handler {
	return func(ctx sdk.Context, content rvotetypes.Content) error {
		switch c := content.(type) {
		case *types.SetChainEraProposal:
			return k.ProcessSetChainEraProposal(ctx, c)
		case *types.BondReportProposal:
			return k.ProcessBondReportProposal(ctx, c)
		case *types.ActiveReportProposal:
			return k.ProcessActiveReportProposal(ctx, c)
		case *types.TransferReportProposal:
			return k.ProcessTransferReportProposal(ctx, c)
		case *types.ExecuteBondProposal:
			return k.ProcessExecuteBondProposal(ctx, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized param proposal content type: %T", c)
		}
	}
}
