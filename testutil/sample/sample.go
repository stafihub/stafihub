package sample

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/cosmoscmd"
)

var (
	TestDenom      = "TestDenom"
	TestDenom1     = "TestDenom1"
	TestDenom2     = "TestDenom2"
	TestAddrPrefix = "Tstafi"
	TestAdmin      string
	TestAdminAcc   sdk.AccAddress
)

func init() {
	cosmoscmd.SetPrefixes("stafi")
	TestAdminAcc = OriginAccAddress()
	TestAdmin = TestAdminAcc.String()
}

// AccAddress returns a sample account address
func AccAddress() string {
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.AccAddress(addr).String()
}

func OriginAccAddress() sdk.AccAddress {
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.AccAddress(addr)
}
