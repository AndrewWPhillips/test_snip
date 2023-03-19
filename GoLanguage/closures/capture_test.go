package __

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"testing"
	"time"
)

// TestCapture demonstrates that captured variables are references
func TestCapture(t *testing.T) {
	a := 1
	f := func() int {
		return a * a
	}
	a = 2
	log.Println(f()) // 4
	a = 3
	log.Println(f()) // 9
	a = 4
}

func TestGoroutineLoop(t *testing.T) {
	var mu sync.Mutex
	a := 2

	start, end := 10, 100
	var wg sync.WaitGroup
	wg.Add(end - start)
	for b := start; b < end; {
		c := b
		go func(d int) {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)
			mu.Lock()
			defer mu.Unlock()
			fmt.Println(a, b, c, d) // b will have values between 10 and 100 (inclusive) but some duplicated and some lost
			// b is incorrect (captured loop variable) - if none of the goroutines execute till after the loop finishes then they will only see the final value 100
			// c and d are captured correctly - we will see every value from 10 to 99
			// * b is the captured loop variable so can have any value from start to end when *any* of the goroutines use the value
			// * c is captured as a different variable on each loop iteration
			// * d is "captured" by being passed as a parameter
			wg.Done()
		}(b) // go vet: loop variable b captured by func literal
		mu.Lock()
		b++ // though b is the "loop" var we need to lock it to prevent a data race
		mu.Unlock()
	}
	wg.Wait()
}
