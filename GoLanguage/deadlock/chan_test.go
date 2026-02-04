package __

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestVenkatramanaOrig(t *testing.T) {
	tasks := make(chan int, 10)
	results := make(chan int, 10)

	for w := 0; w < 3; w++ {
		go func() {
			for task := range tasks {
				time.Sleep(10 * time.Millisecond)
				results <- task * 2
			}
		}()
	}

	for i := 0; i < 1000; i++ {
		tasks <- i
	}

	for i := 0; i < 1000; i++ {
		fmt.Println(<-results)
	}

	fmt.Println("END")
}

func TestVenkatramanaSoln(t *testing.T) {
	tasks := make(chan int, 10)
	results := make(chan int, 10)

	for w := 0; w < 3; w++ {
		go func() {
			for task := range tasks {
				time.Sleep(10 * time.Millisecond)
				results <- task * 2
			}
		}()
	}

	go func() {
		for i := 0; i < 1000; i++ {
			tasks <- i
		}
	}()

	for i := 0; i < 1000; i++ {
		fmt.Println(<-results)
	}

	fmt.Println("END")
}

func TestVenkatramanaBetter(t *testing.T) {
	tasks := make(chan int)
	results := make(chan int)
	wg := sync.WaitGroup{}

	// Start the workers
	const workers = 3
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func() {
			defer wg.Done()
			for task := range tasks {
				results <- task * 2
			}
		}()
	}

	// Start goroutine to assign the work
	go func() {
		for i := 0; i < 10; i++ {
			tasks <- i
		}
		close(tasks)

		// close output chan when finished
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Println(r)
	}
}
