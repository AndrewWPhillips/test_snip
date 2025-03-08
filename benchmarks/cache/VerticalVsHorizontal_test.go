package __

// banchmark traversing 2d array in different directions to see cache effects (if any)

import (
	"golang.org/x/sync/errgroup"
	"math/rand/v2"
	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkArray2D(b *testing.B) {
	//const maxH, maxV = 4096, 4096 // 9 ms/op
	const maxH, maxV = 8192, 8192 // H then V = 40 ms/op, V then H = 38 ms/op
	a := Generate2D(maxH, maxV)

	for b.Loop() {
		odd := 0
		for h := range maxH {
			for v := range maxV {
				if a[h*maxV+v]%2 != 0 {
					odd++
				}
			}
		}
	}
}

func BenchmarkArrayConcurrentGlobal(b *testing.B) {
	const GoRoutines = 1 // 6.7 ms/op, 4 allocs
	//const GoRoutines = 4 // 7.5 ms/op, 10 allocs
	//const GoRoutines = 10 // 7.5 ms/op, 23 allocs

	a := Generate2D(1024, 1024)
	var result uint64

	for b.Loop() {
		var odd atomic.Uint64
		var group errgroup.Group
		chunkSize := len(a) / GoRoutines

		start, end := 0, chunkSize
		for range GoRoutines {
			s, e := start, end
			group.Go(func() error {
				for i := s; i < e; i++ {
					if a[i]%2 == 0 {
						//atomic.AddInt64(&odd, 1)
						odd.Add(1)
					}
				}
				return nil
			})
			start = end
			end += chunkSize
		}

		if err := group.Wait(); err != nil {
			b.Fatal(err)
		}
		result = odd.Load()
	}
	println(result)
}

func BenchmarkArrayConcurrentLocalErrGroup(b *testing.B) {
	//const GoRoutines = 1 // 0.73 ms/op
	//const GoRoutines = 2 // 0.41 ms/op
	const GoRoutines = 4 // 0.18 ms/op
	//const GoRoutines = 7 // 0.22
	//const GoRoutines = 8 // 0.20
	//const GoRoutines = 10 // 0.19 ns/op
	//const GoRoutines = 32 // 0.16

	a := Generate2D(1024, 1024)
	var result uint64

	for b.Loop() {
		var odd atomic.Uint64
		var group errgroup.Group
		chunkSize := len(a) / GoRoutines

		start, end := 0, chunkSize
		for range GoRoutines {
			s, e := start, end
			group.Go(func() error {
				local := 0
				for i := s; i < e; i++ {
					if a[i]%2 == 0 {
						local++
					}
				}
				odd.Add(uint64(local))
				return nil
			})
			start = end
			end += chunkSize
		}

		group.Wait()
		result = odd.Load()
	}
	println(result)
}

func BenchmarkArrayConcurrentLocalWaitGroup(b *testing.B) {
	//const GoRoutines = 1 // 958 ns/op
	const GoRoutines = 4 // 3655 ns/op
	//const GoRoutines = 10 // 6332 ns/op

	a := Generate2D(8192, 8192)
	var result uint64

	for b.Loop() {
		var odd atomic.Uint64
		var wg = &sync.WaitGroup{}
		wg.Add(GoRoutines)
		chunkSize := len(a) / GoRoutines

		start, end := 0, chunkSize
		for range GoRoutines {
			go func(s, e int) {
				local := 0
				for i := s; i < e; i++ {
					if a[i]%2 == 0 {
						local++
					}
				}
				odd.Add(uint64(local))
				wg.Done()
			}(start, end)
			start = end
			end += chunkSize
		}

		wg.Wait()
		result = odd.Load()
	}
	println(result)
}

func BenchmarkArray0(b *testing.B) { // 45000 ns/op, 0 allocs
	a := Generate2D(256, 256)
	var result int

	for b.Loop() {
		var i, local int
		for i = 0; i < len(a); i++ {
			if a[i]%2 == 0 {
				local++
			}
		}
		result = local
	}
	println(result)
}

func BenchmarkArray1(b *testing.B) { // 20_000 ns, 0 allocs
	const Split = 10
	a := Generate2D(256, 256)
	var result int

	for b.Loop() {
		var total int
		chunkSize := len(a) / Split
		start, end := 0, chunkSize
		for range Split {
			func(s, e int) {
				var i, local int
				for i = s; i < e; i++ {
					if a[i]%2 == 0 {
						local++
					}
				}
				// total += local
			}(start, end)
			start = end
			end += chunkSize
		}
		result = total
	}
	println(result)
}

func BenchmarkArray2(b *testing.B) {
	//const Split = 1
	//const Split = 2
	//const Split = 4 // 30_000 ns/op
	const Split = 10
	a := Generate2D(256, 256)
	var result int

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var total int
		ch := make(chan int, Split)
		chunkSize := len(a) / Split
		start, end := 0, chunkSize
		for range Split {
			go func(s, e int) {
				count := 0
				for i := s; i < e; i++ {
					count += a[i]
				}
				ch <- count
			}(start, end)
			start = end
			end += chunkSize
		}
		for range Split {
			total += <-ch
		}
		result = total
	}
	println(result)
}

// --- HELPERS ----

func Generate2D(h, v int) (r []int) {
	r = make([]int, 0, h*v)
	for range h * v {
		r = append(r, int(rand.Int32N(1e6)))
	}
	return
}
