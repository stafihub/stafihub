package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Constants pertaining to a Content object
const (
	PropIdLength = 64
)

// Content defines an interface that a proposal must implement. It contains
// information such as the title and description along with the type and routing
// information for the appropriate handler to process the proposal. Content can
// have additional fields, which will handled by a proposal's Handler.
// TODO Try to unify this interface with types/module/simulation
// https://github.com/cosmos/cosmos-sdk/issues/5853
type Content interface {
	GetPropId() string
	GetDenom() string
	ProposalRoute() string
	ProposalType() string
	ValidateBasic() error
	String() string
}

// Handler defines a function that handles a proposal after it has passed the
// governance process.
type Handler func(ctx sdk.Context, content Content) error

// ValidateAbstract validates a proposal's abstract contents returning an error
// if invalid.
func ValidateAbstract(c Content) error {
	propId := c.GetPropId()
	if len(propId) != PropIdLength {
		return sdkerrors.Wrapf(ErrInvalidProposalContent, "propId's length %d != %d", len(propId), PropIdLength)
	}

	denom := c.GetDenom()
	if len(denom) == 0 {
		return sdkerrors.Wrap(ErrInvalidProposalContent, "proposal denom cannot be blank")
	}

	return nil
}
