package crypto

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"testing"
)

func TestSha1Method(t *testing.T) {
	h1 := sha1.New()
	h1.Write([]byte("abc"))
	h1.Write([]byte("def"))

	h2 := sha1.New()
	h2.Write([]byte("abcdef"))
	if bytes.Compare(h1.Sum(nil), h2.Sum(nil)) != 0 {
		t.FailNow()
	}
	fmt.Println(h1.Sum(nil))
}

func TestSha1Func(t *testing.T) {
	fmt.Println(sha1.Sum([]byte("abcdef")))
}
