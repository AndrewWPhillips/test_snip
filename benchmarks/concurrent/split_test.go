package __

import (
	"golang.org/x/sync/errgroup"
	"sync"
	"testing"
)

func BenchmarkSplitChan(b *testing.B) {
	//const Size, Split = 4096 * 4096, 1 // 5.2 ms/op, 3 allocs
	//const Size, Split = 4096 * 4096, 2 // 2.6 ms/op, 5 allocs
	//const Size, Split = 4096 * 4096, 3 // 1.5 ms/op, 7 allocs
	//const Size, Split = 4096 * 4096, 4 // 1.2 ms/op, 9 allocs
	//const Size, Split = 4096 * 4096, 5 // 0.97 ms/op, 11 allocs
	//const Size, Split = 4096 * 4096, 6 // 0.86 ms/op, 13 allocs
	//const Size, Split = 4096 * 4096, 7 // 1.4 ms/op, 15 allocs
	//const Size, Split = 4096 * 4096, 8 // 1.2 ms/op, 17 allocs
	//const Size, Split = 4096 * 4096, 9 // 1.0 ms/op
	//const Size, Split = 4096 * 4096, 10 // 0.95 ms/op
	//const Size, Split = 4096 * 4096, 11 // 0.87 ms/op
	//const Size, Split = 4096 * 4096, 12 // 0.86

	//const Size, Split = 8192 * 8192, 2 // 11.3 ms/op
	const Size, Split = 8192 * 8192, 4 // 4.7 ms/op
	//const Size, Split = 8192 * 8192, 6 // 3.6 ms/op
	//const Size, Split = 8192 * 8192, 7 // 5.2 ms/op
	//const Size, Split = 8192 * 8192, 8 // 4.9 ms/op
	var result int

	for b.Loop() {
		var total int
		ch := make(chan int, Split)
		chunkSize := Size / Split
		start, end := 0, chunkSize
		for range Split {
			go func(s, e int) {
				count := 0
				for i := s; i < e; i++ {
					count++
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

func BenchmarkSplitWaitgroup(b *testing.B) {
	const Size, Split = 4096 * 4096, 12

	for b.Loop() {
		var wg sync.WaitGroup
		chunkSize := Size / Split
		start, end := 0, chunkSize
		wg.Add(Split)
		for range Split {
			go func(s, e int) {
				count := 0
				for i := s; i < e; i++ {
					count++
				}
				wg.Done()
			}(start, end)
			start = end
			end += chunkSize
		}
		wg.Wait()
	}
}

func BenchmarkSplitErrgroup(b *testing.B) {
	const Size, Split = 4096 * 4096, 12

	for b.Loop() {
		var group errgroup.Group
		chunkSize := Size / Split
		start, end := 0, chunkSize
		for range Split {
			s, e := start, end
			group.Go(func() error {
				count := 0
				for i := s; i < e; i++ {
					count++
				}
				return nil
			})
			start = end
			end += chunkSize
		}
		group.Wait()
	}
}
