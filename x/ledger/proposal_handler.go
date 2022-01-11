package ledger

import (
	rvotetypes "github.com/stafiprotocol/stafihub/x/rvote/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafiprotocol/stafihub/x/ledger/keeper"
	"github.com/stafiprotocol/stafihub/x/ledger/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewParamChangeProposalHandler creates a new governance Handler for a ParamChangeProposal
func NewProposalHandler(k keeper.Keeper) rvotetypes.Handler {
	return func(ctx sdk.Context, content rvotetypes.Content) error {
		switch c := content.(type) {
		case *types.SetChainEraProposal:
			return k.ProcessSetChainEraProposal(ctx, c)
		case *types.BondReportProposal:
			return k.ProcessBondReportProposal(ctx, c)
		case *types.BondAndReportActiveProposal:
			return k.ProcessBondAndReportActiveProposal(ctx, c)
		case *types.ActiveReportProposal:
			return k.ProcessActiveReportProposal(ctx, c)
		case *types.WithdrawReportProposal:
			return k.ProcessWithdrawReportProposal(ctx, c)
		case *types.TransferReportProposal:
			return k.ProcessTransferReportProposal(ctx, c)
		case *types.ExecuteBondProposal:
			return k.ProcessExecuteBondProposal(ctx, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized param proposal content type: %T", c)
		}
	}
}
