package types

func (p *Proposal) IsExpired(block int64) bool {
	return p.ExpireBlock != p.StartBlock && block > p.ExpireBlock
}
