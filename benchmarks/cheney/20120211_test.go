package __test

import "testing"

// SIZE=1000 results (core i5 late 2011 mac mini, 10.7.3)
// % go test -v -run='XXX' -bench='.'
// PASS
// BenchmarkUpdate 500000 2996 ns/op
// BenchmarkManual 500000 4642 ns/op
// BenchmarkUnroll 1000000 2824 ns/op

type E struct {
	A, B, C, D int
}

func (e *E) update(a, b, c, d int) {
	e.A += a
	e.B += b
	e.C += c
	e.D += d
}

var SIZE = 1000 // needed to make a valid testable package

func TestNothing(t *testing.T) {}

func assert(e []E, b *testing.B) {
	for _, v := range e {
		//if v.A != b.N || v.B != b.N2 || v.C != b.N3 || v.D != b.N4 {
		//	b.Errorf("Expected: %d, %d, %d, %d; actual: %d, %d, %d, %d",
		//		b.N, b.N2, b.N3, b.N*4, v.A, v.B, v.C, v.D)
		//}
		if v.A != b.N {
			b.Errorf("Expected: %d; actual: %d, %d, %d, %d",
				b.N, v.A, v.B, v.C, v.D)
		}
	}
}

func BenchmarkUpdate(b *testing.B) {
	var e = make([]E, SIZE)
	for j := 0; j < b.N; j++ {
		for i := range e {
			e[i].update(1, 2, 3, 4)
		}
	}
	b.StopTimer()
	assert(e, b)
}

func BenchmarkManual(b *testing.B) {
	var e = make([]E, SIZE)
	for j := 0; j < b.N; j++ {
		for i := range e {
			e[i].A += 1
			e[i].B += 2
			e[i].C += 3
			e[i].D += 4
		}
	}
	b.StopTimer()
	assert(e, b)
}

func BenchmarkUnroll(b *testing.B) {
	var e = make([]E, SIZE)
	for j := 0; j < b.N; j++ {
		for i := range e {
			v := &e[i]
			v.A += 1
			v.B += 2
			v.C += 3
			v.D += 4
		}
	}
	b.StopTimer()
	assert(e, b)
}
