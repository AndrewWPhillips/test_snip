package __

import (
	"crypto/sha1"
	"hash/crc64"
	"hash/fnv"
	"testing"
)

func BenchmarkCRC(b *testing.B) { // 119 ns/op
	for b.Loop() {
		h := crc64.New(crc64.MakeTable(crc64.ECMA))
		h.Sum([]byte("abcdef"))
	}
}

func BenchmarkSha1(b *testing.B) {
	for b.Loop() {
		h := sha1.New()
		h.Sum([]byte("abcdef"))
	}
}

func BenchmarkFNV(b *testing.B) { // 68 ns/op
	for b.Loop() {
		h := fnv.New64()
		_ = h.Sum([]byte("abcdef"))
	}
}
