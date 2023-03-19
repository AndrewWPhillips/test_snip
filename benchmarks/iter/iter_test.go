package __

import (
	"strconv"
	"testing"
)

type (
	Store     map[string]int
	StoreIter struct {
		ch   <-chan int
		curr int
	}
)

func NewStoreIter(target Store) StoreIter {
	outCh := make(chan int)
	go func() {
		for _, v := range target {
			outCh <- v
		}
		close(outCh)
	}()
	return StoreIter{ch: outCh}
}

func (s *StoreIter) Next() bool {
	var ok bool
	s.curr, ok = <-s.ch
	return ok
}

func (s *StoreIter) Value() int {
	return s.curr
}

// BenchmarkIter was a test of some sort of iterator using a chan
// TODO: check out if this is useful for something and if so make it generic - any map type OR any range type??
func BenchmarkIter(b *testing.B) { // 129 msec/op
	b.StopTimer()
	m := make(Store)
	for i := 0; i < 1e6; i++ {
		m[strconv.Itoa(i)] = i
	}
	b.StartTimer()

	var v int
	for n := 0; n < b.N; n++ {
		for iter := NewStoreIter(m); iter.Next(); {
			v = iter.Value()
		}
	}
	_ = v
}

func BenchmarkMap(b *testing.B) { // 60 msec/op
	b.StopTimer()
	m := make(Store)
	for i := 0; i < 1e6; i++ {
		m[strconv.Itoa(i)] = i
	}
	b.StartTimer()

	var v int
	for i := 0; i < b.N; i++ {
		for _, v = range m {
		}
	}
	_ = v
}
