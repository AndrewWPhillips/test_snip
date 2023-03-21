package __

import (
	"log"
	"runtime"
	"runtime/pprof"
	"testing"
	"time"
)

// TestSleep stops execution of go-routines using Sleep().  Note that Sleep uses the scheduler so it should not (or no
// longer) hold a thread, so is preferred to time.After (see below). TODO: check number of threads in use
func TestSleep(t *testing.T) {
	const count = 8000 // Note: to use more than 8192 you need to turn off the race detector
	for i := 0; i < count; i++ {
		go func() {
			time.Sleep(time.Second)
		}()
	}
	runtime.Gosched()
	// at this point there are 'count' runnable go-routines
	log.Println("Goroutines:", runtime.NumGoroutine())

	threadProfile := pprof.Lookup("threadcreate")
	log.Println("Threads:", threadProfile.Count())
}

// TestAfter tests time.After (reads from the returned channel).  Effectively does the same
// as time.Sleep, but is *not* recommended as it creates temporary timers which adds to GC load
// and will not be recovered until the next GC.
func TestAfter(t *testing.T) {
	for i := 0; i < 8000; i++ {
		go func() {
			<-time.After(time.Second)
		}()
	}
	runtime.Gosched()
	// at this point there are 100 waiting go-routines (or they will be once all have started)
	log.Println("Goroutines:", runtime.NumGoroutine())

	threadProfile := pprof.Lookup("threadcreate")
	log.Println("Threads:", threadProfile.Count())
}
