package crypto

import (
	crand "crypto/rand"
	"log"
	"math/big"
	"testing"
)

func TestBigRand(t *testing.T) {
	limit, _ := new(big.Int).SetString("1234566884078423975497543987543987543987543987", 10)
	//randInt, err := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	randInt, err := crand.Int(crand.Reader, limit)
	log.Println(randInt, err)
}
