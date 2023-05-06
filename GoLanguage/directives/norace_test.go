package __

import (
	"log"
	"sync"
	"testing"
)

//go:norace
func racer(pi *int, wg *sync.WaitGroup) {
	*pi++
	wg.Done()
}

// TestNoRace has a race condition on pi in the race() function above
// because we use the "norace" directive the runtime does not detect the problem (using -race option)
func TestNorace(t *testing.T) {
	i := 1
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go racer(&i, wg)
	// This would give a data race except that we used //go:norace
	go racer(&i, wg)
	wg.Wait()
	log.Println(i)
}
