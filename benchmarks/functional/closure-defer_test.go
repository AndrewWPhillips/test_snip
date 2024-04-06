package __

import (
	"sync"
	"testing"
)

func BenchmarkNoFunc(b *testing.B) { // BenchmarkNoFunc-12      217909747               10.86 ns/op
	r := 0
	var rwMutex sync.RWMutex
	for i := 0; i < b.N; i++ {
		var value int
		rwMutex.RLock()
		value = 42
		rwMutex.RUnlock()
		r = value
	}
	println(r)
}

// BenchmarkYesFunc is the same code as BenchmarkNoFunc except we put the inner code in a nested
// function, so we can defer the call to RUnlock
func BenchmarkYesFunc(b *testing.B) { // BenchmarkYesFunc-12     148215850               15.70 ns/op
	r := 0
	var rwMutex sync.RWMutex
	for i := 0; i < b.N; i++ {
		var value int
		func() {
			rwMutex.RLock()
			defer rwMutex.RUnlock()
			value = 42
		}()
		r = value
	}
	println(r)
}
