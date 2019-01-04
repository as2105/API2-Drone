package ethereum

import (
	"github.com/pborman/uuid"
	"math/big"
	"time"
)

func bytesToUUID(bin [16]byte) uuid.UUID {
	return uuid.Array(bin).UUID()
}

func bigintToTime(i *big.Int) time.Time {
	i64 := i.Int64()
	return time.Unix(i64, 0)
}

func timeToBigint(t time.Time) *big.Int {
	i64 := t.Unix()
	return big.NewInt(i64)
}
