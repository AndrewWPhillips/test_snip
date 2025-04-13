//go:build go1.24

package __

import (
	"testing"
)

func BenchmarkNotAliveInlineGo124(b *testing.B) { // 4
	for b.Loop() {
		_ = f42()
	}
}

func BenchmarkNotAliveNotInlineGo124(b *testing.B) { // 4
	for b.Loop() {
		_ = g42()
	}
}

type MyType string

//go:noinline
func (mt MyType) String() string {
	return string(mt)
}

func (mt MyType) Close() {
}

func New() MyType {
	return "42"
}

func BenchmarkOld(b *testing.B) {
	var s string
	n := New() // setup
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = n.String()
	}
	b.StopTimer()
	n.Close() // teardown
	TestGlobal = s
}

func BenchmarkNew(b *testing.B) {
	n := New() // setup
	for b.Loop() {
		n.String()
	}
	n.Close() // teardown
}

// Tests of what gets inlined now

func i1() string {
	return "42"
}

func i0() string {
	return i1()
}

//go:noinline
func j1() string {
	return "42"
}

//go:noinline
func j0() string {
	return j1()
}

func k() string {
	return j1() // no inline
}

func BenchmarkINew(b *testing.B) { // 2.4 = i0, assignment
	b.Context()
	for b.Loop() {
		i0()
	}
}
func BenchmarkJNew(b *testing.B) { // 2.7 = j0, j1, assignment
	for b.Loop() {
		j0()
	}
}
func BenchmarkKNew(b *testing.B) { // 2.7 = k, j1, assignment
	for b.Loop() {
		k()
	}
}

func BenchmarkIOld(b *testing.B) { // 0.5 = assignment
	var s string
	for i := 0; i < b.N; i++ {
		s = i0()
	}
	TestGlobal = s
}
func BenchmarkJOld(b *testing.B) { // 2.4 = j0, j1, assignment
	var s string
	for i := 0; i < b.N; i++ {
		s = j0()
	}
	TestGlobal = s
}
func BenchmarkKOld(b *testing.B) { // 2.25 = j1, assignment
	var s string
	for i := 0; i < b.N; i++ {
		s = k()
	}
	TestGlobal = s
}
