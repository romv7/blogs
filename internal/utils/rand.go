package utils

import (
	crand "crypto/rand"
	"log"
	"math/big"
	"time"
)

func RandomUniqueId() uint64 {
	n, err := crand.Int(crand.Reader, big.NewInt(2^31))
	if err != nil {
		log.Panic(err)
	}

	return uint64(time.Now().UTC().UnixMilli()) - uint64(n.Int64())
}
