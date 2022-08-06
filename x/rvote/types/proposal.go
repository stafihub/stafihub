package types

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/gogo/protobuf/proto"
	yaml "gopkg.in/yaml.v2"
)

func (p *Proposal) SetContent(content Content) error {
	msg, ok := content.(proto.Message)
	if !ok {
		return fmt.Errorf("%T does not implement proto.Message", content)
	}

	any, err := types.NewAnyWithValue(msg)
	if err != nil {
		return err
	}

	p.Content = any
	return nil
}

// GetContent returns the proposal Content
func (p *Proposal) GetContent() Content {
	content, ok := p.Content.GetCachedValue().(Content)
	if !ok {
		return nil
	}
	return content
}

func (p *Proposal) PropId() string {
	content := p.GetContent()
	if content == nil {
		return ""
	}
	return content.GetPropId()
}

func (p *Proposal) ProposalType() string {
	content := p.GetContent()
	if content == nil {
		return ""
	}
	return content.ProposalType()
}

func (p *Proposal) ProposalRoute() string {
	content := p.GetContent()
	if content == nil {
		return ""
	}
	return content.ProposalRoute()
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (p *Proposal) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	var content Content
	return unpacker.UnpackAny(p.Content, &content)
}

func (p *Proposal) HasVoted(proposer string) bool {
	for _, v := range p.Voted {
		if v == proposer {
			return true
		}
	}

	return false
}

func (p *Proposal) IsExpired(block int64) bool {
	return p.ExpireBlock != 0 && block > p.ExpireBlock
}

// String implements stringer interface
func (p *Proposal) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

var validProposalTypes = map[string]struct{}{}

// RegisterProposalType registers a proposal type. It will panic if the type is
// already registered.
func RegisterProposalType(ty string) {
	if _, ok := validProposalTypes[ty]; ok {
		panic(fmt.Sprintf("already registered proposal type: %s", ty))
	}

	validProposalTypes[ty] = struct{}{}
}

// IsValidProposalType returns a boolean determining if the proposal type is
// valid.
//
// NOTE: Modules with their own proposal types must register them.
func IsValidProposalType(ty string) bool {
	_, ok := validProposalTypes[ty]
	return ok
}
