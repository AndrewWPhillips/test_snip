package time

import (
	"fmt"
	"testing"
	"time"
)

// BenchmarkNow - about 45 ns/op on home (Windows 10) computer with Go 1.16.6, 6 ns/op on work computer
// time.Now() is 30% faster if we remove the (unnec.) non-monotonic test (3 lines):
//	if uint64(sec)>>33 != 0 {
//		return Time{uint64(nsec), sec + minWall, Local}
//	}
func BenchmarkNow(b *testing.B) {
	var now time.Time
	for i := 0; i < b.N; i++ {
		now = time.Now()
	}
	fmt.Println(now)
}

// BenchmarkNowUTC - about 111 ns/op on home computer with Go 1.16.6
func BenchmarkNowUTC(b *testing.B) {
	var now time.Time
	for i := 0; i < b.N; i++ {
		now = time.Now().UTC()
	}
	fmt.Println(now)
}

// BenchmarkNowCached - about 11 ns/op on home computer with Go 1.16.6
func BenchmarkNowCached(b *testing.B) {
	var now time.Time
	nowCached := time.Now()
	for i := 0; i < b.N; i++ {
		now = nowCached
	}
	fmt.Println(now)
}
