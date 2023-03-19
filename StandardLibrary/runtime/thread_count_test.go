package __

import (
	"net"
	"runtime"
	"runtime/pprof"
	"sync"
	"testing"
)

// TestThreadCount tests how to get the number of threads used by the Go runtime and how CGO does not release them
func TestThreadCount(t *testing.T) {
	threadProfile := pprof.Lookup("threadcreate") // "goroutine" "threadcreate" "heap" "allocs" "block" "mutex"
	println(threadProfile.Count())

	//debug.SetMaxThreads(15)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			net.LookupHost("www.google.com")
		}()
	}
	wg.Wait()
	// According to https://www.sobyte.net/post/2021-06/golang-number-of-threads-in-the-running-program/
	// there may be more than 100 threads at this point because CGO is used for LookupHost (in this case)
	// and threads are never returned to the OS
	println(threadProfile.Count())
}

// TestThreadKill shows how to return one thread to the OS
// Note that the Go runtime does not return threads yet (Go 1.18) but the runtime and std lib do their best to not
// allocate threads.  It can be a problem if you use CGO or call a function that does.
func TestThreadKill(t *testing.T) {
	threadProfile := pprof.Lookup("threadcreate") // "goroutine" "threadcreate" "heap" "allocs" "block" "mutex"
	println(threadProfile.Count())
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		runtime.LockOSThread()
		return
		// when this go-routine end the thread is returned
	}()
	wg.Wait()

	println(threadProfile.Count())
}
