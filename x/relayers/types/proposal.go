package types

func (p *Proposal) IsExpired(block int64) bool {
	return p.ExpireBlock != p.StartBlock && block > p.ExpireBlock
}

func (p *Proposal) ProposalRoute() string {
	return p.Content.ProposalRoute
}
