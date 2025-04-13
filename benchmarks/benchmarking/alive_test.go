package __

import (
	"strconv"
	"testing"
)

func BenchmarkAlmostNothing(b *testing.B) { // 0.3 ns/op
	const s = "abc"
	for i := 0; i < b.N; i++ {
		_ = s + s
	}
}

func BenchmarkAlmostNothing2(b *testing.B) { // 0.4 ns/op
	const s = "abc"
	var result string
	for i := 0; i < b.N; i++ {
		result = s + s
	}
	_ = result
}

func BenchmarkAlmostNothing3(b *testing.B) { // 0.3 ns/op
	const s = "abc"
	var result string
	for i := 0; i < b.N; i++ {
		result = s + s
	}
	println(result)
}

func BenchmarkAlmostSomething(b *testing.B) { // 21 ns/op
	var s = "abc"
	for i := 0; i < b.N; i++ {
		_ = s + s
	}
}

func BenchmarkAlmostSomething2(b *testing.B) { // 48 ns/op
	var s, s2 = "abc", ""
	for i := 0; i < b.N; i++ {
		s2 = s + s
	}
	println(s2)
}

func BenchmarkNotKeptAlive(b *testing.B) { // 28 ns/op
	s1, s2 := "abc", "def"
	for i := 0; i < b.N; i++ {
		_ = s1 + s2
	}
}

func BenchmarkKeptAlive(b *testing.B) { // 49 ns/op
	s1, s2 := "abc", "def"
	var s string
	for i := 0; i < b.N; i++ {
		s = s1 + s2
	}
	println(s)
}

var i42 = 42

func f42() string { return strconv.Itoa(i42) }

//go:noinline
func g42() string { return strconv.Itoa(i42) }

func BenchmarkNotAliveInline(b *testing.B) { // 3
	for i := 0; i < b.N; i++ {
		_ = f42()
	}
}

func BenchmarkAliveInline2(b *testing.B) { // 3
	var s string
	for i := 0; i < b.N; i++ {
		s = f42()
	}
	println(s)
}

func BenchmarkNotAliveNotInline(b *testing.B) { // 4
	for i := 0; i < b.N; i++ {
		_ = g42()
	}
}

func BenchmarkAliveNotInline2(b *testing.B) { // 4
	var s string
	for i := 0; i < b.N; i++ {
		s = g42()
	}
	println(s)
}
