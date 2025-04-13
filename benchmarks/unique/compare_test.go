package __

// test speed of comparing unique handles vs strings

import (
	"crypto/rand"
	"testing"
	"unique"
)

func BenchmarkCompareUniqueHandles(b *testing.B) { // 0.27
	h1 := unique.Make(rand.Text())
	h2 := unique.Make(rand.Text())
	for i := 0; i < b.N; i++ {
		_ = h1 == h2
	}
}

func BenchmarkCompareUniqueStrings(b *testing.B) { // 3 ns/op
	s1 := rand.Text()
	s2 := rand.Text()
	for i := 0; i < b.N; i++ {
		_ = s1 == s2
	}
}
