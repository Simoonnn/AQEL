package common

import (
	"fmt"
	"github.com/raidoNetwork/RDO_v2/proto/prototype"
	"time"
)

// Outputs types
const (
	UnspentTxO = 0
	SpentTxO   = 1
)

// Transaction types
const (
	NormalTxType  = 1
	GenesisTxType = 2
	FeeTxType     = 3
	RewardTxType  = 4
	StakeTxType   = 5
	UnstakeTxType = 6
)

const (
	BlackHoleAddress = "0x0000000000000000000000000000000000000000"
)

// Test settings
const (
	AccountNum  = 700
	StartAmount = 1e12 //10000000000000 // 1 * 10e12
)

// StatFmt parse time.Duration to the needed string format
func StatFmt(d time.Duration) string {
	return fmt.Sprintf("%d μs", int64(d/time.Microsecond))
}

// IsLegacyTx check transaction type and return true if transaction is legacy:
// send coins from one address to another or stake.
func IsLegacyTx(tx *prototype.Transaction) bool {
	switch tx.Type {
	case NormalTxType:
		fallthrough
	case StakeTxType:
		fallthrough
	case UnstakeTxType:
		return true
	case FeeTxType:
		fallthrough
	case RewardTxType:
		fallthrough
	case GenesisTxType:
		return false
	default:
		return false
	}
}
