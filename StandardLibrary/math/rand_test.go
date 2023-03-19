package __

import (
	"log"
	"math/rand"
	"testing"
)

func TestRandBytes(t *testing.T) {
	var a [10]byte
	rand.Read(a[:])
	log.Println(a)

}
