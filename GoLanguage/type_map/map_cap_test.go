package maps

import (
	"log"
	"runtime"
	"testing"
	"unsafe"
)

// TestCapacity tests the capacity parameter to make
func TestCapacity(t *testing.T) {
	m := make(map[int]int, 10)
	log.Println(len(m), unsafe.Sizeof(m))
	m[1] = 2
	m[2]++
	log.Println(len(m), unsafe.Sizeof(m))
}

func TestMemoryRelease(t *testing.T) {
	const mapSize = 3e6
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	log.Printf("Alloc = %v MiB", ms.Alloc/1e6)

	m := make(map[int]int, mapSize)
	for i := 0; i < mapSize; i++ {
		m[i] = i
	}
	runtime.ReadMemStats(&ms)
	log.Printf("Alloc = %v MiB", ms.Alloc/1e6)

	for i := range m {
		delete(m, i)
	}
	runtime.ReadMemStats(&ms)
	log.Printf("Alloc = %v MiB", ms.Alloc/1e6)

	m = nil
	runtime.GC()
	runtime.ReadMemStats(&ms)
	log.Printf("Alloc = %v MiB", ms.Alloc/1e6)
}

