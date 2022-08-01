package keeper

import (
	"bytes"
	"encoding/hex"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"golang.org/x/crypto/sha3"
)

type NodeHash []byte

func (nodeHash *NodeHash) String() string {
	return hex.EncodeToString(*nodeHash)
}

func NodeHashFromHexString(hexStr string) (NodeHash, error) {
	bts, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, err
	}
	return bts, nil
}

func VerifyProof(leafNode NodeHash, proof []NodeHash, root NodeHash) bool {
	result := leafNode
	for _, p := range proof {
		result = ConbinedHash(result, p)
	}
	return bytes.EqualFold(result, root)
}

func ConbinedHash(b0, b1 NodeHash) NodeHash {
	bts := make([]byte, 64)
	if bytes.Compare(b0, b1) <= 0 { //a<=b
		copy(bts[0:], b0)
		copy(bts[32:], b1)
	} else {
		copy(bts[0:], b1)
		copy(bts[32:], b0)
	}
	h := sha3.NewLegacyKeccak256()
	h.Write(bts)
	return h.Sum(nil)
}

func GetNodeHash(round, index uint64, account sdk.AccAddress, coin sdk.Coin) NodeHash {
	coinBts := []byte(coin.String())
	accountLen := len(account)
	coinLen := len(coinBts)
	len := 16 + accountLen + coinLen

	bts := make([]byte, len)
	copy(bts[0:], sdk.Uint64ToBigEndian(round))
	copy(bts[8:], sdk.Uint64ToBigEndian(index))
	copy(bts[16:], account)
	copy(bts[16+accountLen:], coinBts)

	h := sha3.NewLegacyKeccak256()
	h.Write(bts)
	return h.Sum(nil)
}
