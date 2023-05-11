package utils

import (
	crand "crypto/rand"
	"log"
	"math/big"
	"time"
)

func RandomUniqueId() uint32 {
	n, err := crand.Int(crand.Reader, big.NewInt(2^31))
	if err != nil {
		log.Panic(err)
	}

	return uint32(time.Now().Unix()) - uint32(n.Int64())
}
