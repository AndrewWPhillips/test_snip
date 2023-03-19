package __

import (
	"log"
	"testing"
)

func TestNilFromSlice(t *testing.T) {
	i := byte(42)
	ss := []*byte{nil, &i, nil}
	i = 43
	ss[0] = &i

	for _, p := range ss {
		log.Println(*p) // panics if p is nil
	}
}
