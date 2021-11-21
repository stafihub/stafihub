module github.com/stafiprotocol/stafihub

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.44.3
	github.com/cosmos/ibc-go v1.2.2
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/spf13/cast v1.3.1
	github.com/tendermint/spm v0.1.8
	github.com/tendermint/tendermint v0.34.14
	github.com/tendermint/tm-db v0.6.4
	google.golang.org/genproto v0.0.0-20211118181313-81c1377c94b1 // indirect
	google.golang.org/grpc v1.42.0 // indirect
)

replace (
	github.com/99designs/keyring => github.com/cosmos/keyring v1.1.7-0.20210622111912-ef00f8ac3d76
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)
