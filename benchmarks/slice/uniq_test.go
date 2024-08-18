package __

import (
	"math/rand"
	"slices"
	"strconv"
	"testing"
)

// uniq_test.go compares removing duplicates from a slice (a bit like the UNIX uniq command
// except that only removes consecutive repeated lines).  I look at using a set to keep track
// of all values seen + also sort and compact (if changing the order does not matter or you
// have/want the slice sorted)

var uniqTestData = []string{"x", "b", "z", "a", "z", "a", "a", "c", "a", "mnmnmnm"}

func randomStrings(n int) []string {
	r := make([]string, 0, n)
	for range n {
		r = append(r, strconv.Itoa(rand.Intn(n)))
	}
	return r
}

func BenchmarkCloneSortAndCompact(b *testing.B) { // 600 ns/op
	var sz int
	for i := 0; i < b.N; i++ {
		d := slices.Clone(uniqTestData)
		slices.Sort(d)
		d = slices.Compact(d)
		sz = len(d)
	}
	println(sz)
}

// BenchmarkCloneSortAndCompact benchmarks removing duplicates by sorting then compacting
// Times (nanoseconds/operation):
// 300 ns           300 - default (10 strings)
// 125 us       125,000 - 1000 strings
// 25 ms     25,009,980 - 100,000
// 7 secs 6,678,440,600 - 10,000,000   9,999,923 allocs/op
func BenchmarkCopySortAndCompact(b *testing.B) {
	//uniqTestData = randomStrings(100_000)
	var sz int
	d := make([]string, len(uniqTestData))
	for i := 0; i < b.N; i++ {
		copy(d, uniqTestData)
		slices.Sort(d)
		sz = len(slices.Compact(d))
	}
	println(sz)
}

// BenchmarkMapAsSet remove duplicates by saving each element as a set (map)
// Times (nanoseconds/operation):
// 320 ns           320 - default (10 string)
// 57 us         57,000 - 1000 strings
// 7 ms       6,911,784 - 100,000
// 4 secs 3,923,599,300 - 10,000,000   10,250,711 allocs/op
func BenchmarkMapAsSet(b *testing.B) {
	//uniqTestData = randomStrings(100_000)
	var sz int
	set := make(map[string]struct{})
	list := make([]string, 0, len(uniqTestData))

	for i := 0; i < b.N; i++ {
		clear(set)
		list = list[:0]

		for _, item := range uniqTestData {
			if _, ok := set[item]; !ok {
				set[item] = struct{}{}
				list = append(list, item)
			}
		}
		sz = len(list)
	}
	println(sz)
}
