package __test

import (
	"fmt"
	"sync"
	"testing"
)

func BenchmarkNoDefer(b *testing.B) {
	var mu sync.Mutex
	saved := 0
	for i := 0; i < b.N; i++ {
		mu.Lock()
		saved = i
		mu.Unlock()
	}
	fmt.Println(saved)
}

func BenchmarkNoDeferFunc(b *testing.B) {
	var mu sync.Mutex
	saved := 0
	for i := 0; i < b.N; i++ {
		func() {
			mu.Lock()
			saved = i
			mu.Unlock()
		}()
	}
	fmt.Println(saved)
}

func BenchmarkDefer(b *testing.B) { // BenchmarkNoDeferFunc-12         194594614               13.62 ns/op
	var mu sync.Mutex
	saved := 0
	for i := 0; i < b.N; i++ {
		func() {
			mu.Lock()
			defer mu.Unlock()
			saved = i
		}()
	}
	fmt.Println(saved)
}
