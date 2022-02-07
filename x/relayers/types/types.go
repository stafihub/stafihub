package types

func NewRelayer(denom, addr string) Relayer {
	return Relayer{
		Denom: denom,
		Addrs: map[string]bool{addr: true},
	}
}