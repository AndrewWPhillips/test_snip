//go:build go1.24

package __

import (
	"strconv"
	"strings"
	"testing"
	"unique"
)

// BenchmarkUniqueAddSame how long to add an existing value
func BenchmarkUniqueAddSame(b *testing.B) { // 42 ns/op
	for b.Loop() {
		unique.Make("1234567")
	}
}

func BenchmarkMapAddSame(b *testing.B) { // 20 ns/op
	m := make(map[string]string)
	str := "1234567"
	for b.Loop() {
		m[str] = str
	}
}

// BenchmarkUniqueAddSameLong - same using a longer string
func BenchmarkUniqueAddSameLong(b *testing.B) { // 66 ns/op
	str := strings.Repeat("abcdefghijklmnopqrstuvwxyz123456", 16)
	for b.Loop() {
		unique.Make(str)
	}
}

func BenchmarkMapAddSameLong(b *testing.B) { // 29
	m := make(map[string]string)
	str := strings.Repeat("abcdefghijklmnopqrstuvwxyz123456", 16)
	for b.Loop() {
		m[str] = str
	}
}

// BenchmarkUniqueAddSame how long to add a different value
// Time = 1700 - 30 (itoa cost) = 1670 but **this make no sense** - see BenchmarkUniqueAddDiff2
func BenchmarkUniqueAddDiff(b *testing.B) { // 1700
	i := 0
	for b.Loop() {
		_ = unique.Make(strconv.Itoa(i))
		i++
	}
}

func BenchmarkMapAddDiff(b *testing.B) { // 410
	m := make(map[string]string)
	i := 0
	for b.Loop() {
		str := strconv.Itoa(i)
		m[str] = str
		i++
	}
}

func BenchmarkUniqueAddDiff2(b *testing.B) { // 1165
	const count = 10_000_000
	var i int

	si := make([]string, count)
	for i = range count {
		//si[i] = strconv.Itoa(rand.Int())
		si[i] = strconv.Itoa(i)
	}
	i = 0

	b.ResetTimer()
	for b.Loop() {
		_ = unique.Make(si[i%count])
		i++
	}
}

// ------------------
// BenchmarkUniqueItoaOnly tests the overhead of ItoA
func BenchmarkUniqueItoaOnly(b *testing.B) { // 34
	i := 0
	for b.Loop() {
		i++
		_ = strconv.Itoa(i)
	}
}

func BenchmarkUniqueIncOnly(b *testing.B) { // 2 ns/op
	i := 0
	for b.Loop() {
		i++
	}
}
