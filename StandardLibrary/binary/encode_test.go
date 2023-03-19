package __

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"testing"
)

func TestStringEncode(t *testing.T) {
	// from SO 73914511
	b := make([]byte, 9)
	rand.Read(b)
	fmt.Println(b)

	ba := make([]byte, 12)
	base64.StdEncoding.Encode(ba, b)
	fmt.Println(string(ba))
}

type LicenseCards struct {
	cards []int
}

func (lc *LicenseCards) PopLicenseCard() int {
	l := len(lc.cards)
	ret := lc.cards[l-1]
	lc.cards = lc.cards[:l-1]
	return ret
}
