package __

import (
	"fmt"
	"sync"
	"testing"
)

func BenchmarkSlice(b *testing.B) { // 25 ns  1 allocs/op
	b.ReportAllocs()
	var o1 []int
	for i := 0; i < b.N; i++ {
		a := make([]int, 4)
		o1 = a
	}
	fmt.Println(o1)
}

func BenchmarkSlice2(b *testing.B) { // 627 ns  3 allocs/op
	b.ReportAllocs()
	var o1 []int
	for i := 0; i < b.N; i++ {
		a := make([]int, 4)
		b := []int{1, 2, 3, 4, 5}
		c := make([]int, 0, 256)
		o1 = a
		o1 = b
		o1 = c
	}
	fmt.Println(o1)
}

func BenchmarkSlice3(b *testing.B) { // 622 ns  3 allocs/op
	b.ReportAllocs()
	var o1, o2, o3 []int
	for i := 0; i < b.N; i++ {
		a := make([]int, 4)
		b := []int{1, 2, 3, 4, 5}
		c := make([]int, 0, 256)
		o1 = a
		o2 = b
		o3 = c
	}
	fmt.Println(o1, o2, o3)
}

//---------------------------

type slicePool struct {
	data []int
	next []int
	mu   sync.Mutex
}

func NewSlicePool(size int) *slicePool {
	d := make([]int, size)
	return &slicePool{
		data: d,
		next: d[:],
	}
}

func (p *slicePool) Get(ln, cp int) []int {
	p.mu.Lock()
	defer p.mu.Unlock()
	// TODO check we have enough cap, check that ln <= cp
	r := p.next[:ln:cp]
	//p.next = p.next[cp:]
	return r
}

func BenchmarkSlicePool2(b *testing.B) { // 50 ns  0 allocs/op
	b.ReportAllocs()
	p := NewSlicePool(1000)
	var o1 []int
	for i := 0; i < b.N; i++ {
		a := p.Get(4, 4)
		b := p.Get(5, 5)
		c := p.Get(0, 256)
		o1 = a
		o1 = b
		o1 = c
	}
	fmt.Println(o1)
}
func BenchmarkSlicePool3(b *testing.B) { // 50 ns  0 allocs/op
	b.ReportAllocs()
	p := NewSlicePool(1000)
	var o1, o2, o3 []int
	for i := 0; i < b.N; i++ {
		a := p.Get(4, 4)
		b := p.Get(5, 5)
		c := p.Get(0, 256)
		o1 = a
		o2 = b
		o3 = c
	}
	fmt.Println(o1, o2, o3)
}
