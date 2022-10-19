package chain

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var wei = big.NewFloat(1_000_000_000_000_000_000)

func IsValidAddress(address string, checksummed bool) bool {
	if !common.IsHexAddress(address) {
		return false
	}
	return !checksummed || common.HexToAddress(address).Hex() == address
}

func EtherToWei(amount float64) *big.Int {
	bigFloat := new(big.Float).Mul(wei, big.NewFloat(amount))
	bigInt, _ := bigFloat.Int(nil)
	return bigInt
}
