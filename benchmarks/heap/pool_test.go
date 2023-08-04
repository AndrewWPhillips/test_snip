package __

import (
	"sync"
	"testing"
)

func BenchmarkNoPool(b *testing.B) { // 13ns/op  1 allocs/op
	b.ReportAllocs()
	var outer *int
	for i := 0; i < b.N; i++ {
		ptr := new(int)
		*ptr = i
		outer = ptr
	}
	println(*outer)
}

//----------------------------------------------------------------

type chanPool struct {
	data []int
	ch   chan *int
}

func NewChanPool(size int) *chanPool {
	r := &chanPool{
		data: make([]int, size),
		ch:   make(chan *int, size),
	}
	for i := 0; i < size; i++ {
		r.ch <- &r.data[i]
	}
	return r
}

func (p *chanPool) Get() *int {
	return <-p.ch
}

func (p *chanPool) Put(pi *int) {
	*pi = 0
	p.ch <- pi
}

func BenchmarkChanPool(b *testing.B) { // 48ns/op  0 allocs/op
	b.ReportAllocs()
	p := NewChanPool(10000)
	outer := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ptr := p.Get()
		*ptr = i
		outer = *ptr
		p.Put(ptr)
	}
	println(outer)
}

//----------------------------------------------------------------

type slicePoolA struct {
	data []int  // "array" of all the elements
	free []*int // list of pointers to available elements
	mu   sync.Mutex
}

func NewSlicePoolA(size int) *slicePoolA {
	r := &slicePoolA{
		data: make([]int, size),
		free: make([]*int, 0, size),
	}
	// set up initial free list to be all the elements
	for i := 0; i < size; i++ {
		r.free = append(r.free, &r.data[i])
	}
	return r
}

func (p *slicePoolA) Get() *int {
	p.mu.Lock()
	defer p.mu.Unlock()
	size := len(p.free)
	if size == 0 {
		return nil
	}
	r := p.free[size-1]
	p.free = p.free[:size-1]
	return r
}

func (p *slicePoolA) Put(pi *int) {
	p.mu.Lock()
	defer p.mu.Unlock()
	*pi = 0
	p.free = append(p.free, pi)
}

func BenchmarkSlicePoolA(b *testing.B) { // 34 ns/op  0 allocs/op
	b.ReportAllocs()
	p := NewSlicePoolA(100)
	outer := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ptr := p.Get()
		*ptr = i
		outer = *ptr
		p.Put(ptr)
	}
	println(outer)
}
