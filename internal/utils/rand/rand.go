package rand

import (
	"crypto/rand"
	"math"
	"math/big"
)

func Rand() (uint64, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(math.Pow(2, 32))))
	if err != nil {
		return 0, err
	}

	return n.Uint64(), nil
}
