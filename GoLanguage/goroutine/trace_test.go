package __

import (
	"io/ioutil"
	"log"
	"math"
	"runtime"
	"runtime/trace"
	"sync"
	"testing"
	"time"
)

func TestGC(t *testing.T) {
	f, err := ioutil.TempFile("", "gc-test-")
	if err != nil {
		t.Fatal(err.Error())
	}
	err = trace.Start(f)
	if err != nil {
		t.Fatal(err.Error())
	}

	// Run GC while goroutines are executing
	wg := startGoroutines(8)
	runtime.GC()
	wg.Wait()

	trace.Stop()
	log.Println("Trace written to", f.Name())
}

// TestSTW checks what happens if a go-routine blocks the stop the world required for GC
func TestSTW(t *testing.T) {
	f, err := ioutil.TempFile("", "stw-test-")
	if err != nil {
		t.Fatal(err.Error())
	}
	err = trace.Start(f)
	if err != nil {
		t.Fatal(err.Error())
	}

	// Start some goroutines
	wg := startGoroutines(8)

	// Start a goroutine that should block STW
	var wg2 sync.WaitGroup
	wg2.Add(1)
	go func() {
		for j := 0; j < 1e12; j++ {
			if j == 1e9 {
				wg2.Done()
			}
		}
	}()
	wg2.Wait()
	time.Sleep(100 * time.Millisecond)

	log.Println("Starting GC")
	runtime.GC()
	wg.Wait()
	trace.Stop()
	log.Println("Trace written to", f.Name())
}

// startGoRoutines runs the specified number of go routines.
// Each goroutine does some "busy" work which should keep one CPU busy for
// a significant fraction of a second (but does not allocate memory).
// It returns after all goroutines have been started running.
// It returns a WaitGroup, so you can wait for all the go-routines finish.
func startGoroutines(num int) *sync.WaitGroup {
	var wgStart, wgStop sync.WaitGroup // wait until they are all started
	wgStart.Add(num)
	wgStop.Add(num)
	for i := 0; i < num; i++ {
		go func(i int) {
			wgStart.Done()
			k := 0
			for j := 0; j < math.MaxInt32; j++ {
				if j%1e6 == 0 {
					runtime.Gosched()
				}
				k = j
			}
			_ = k // make sure the loop is not optimised away
			//log.Println(i, k)
			wgStop.Done()
		}(i)
	}
	// wait till they've all started and return WaitGroup allowing caller to wait for them to finish
	wgStart.Wait()
	return &wgStop
}
