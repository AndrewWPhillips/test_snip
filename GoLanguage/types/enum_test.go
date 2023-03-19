package __

import (
	"log"
	"testing"
)

// see also MyEnum type (myenum.go and myenum_test.go)

func TestEnumAsIndex(t *testing.T) {
	type IT int
	const (
		A IT = 1
	)
	var b [21]string

	b[A] = "abc"
	log.Println(b[0])
}
