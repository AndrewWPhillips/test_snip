package __

import (
	"math/rand/v2"
	"testing"
	"testing/synctest"
	"time"
)

func request(address string) int {
	// simulate getting a response in 1 to 10 seconds
	time.Sleep(time.Duration(1+rand.IntN(10)) * time.Second)
	return len(address) // dummy return value
}

// GetFirstResponse simulates getting the fastest response from several addresses
// WARNING: this code has (deliberate) bugs
func GetFirstResponse(addresses ...string) int {
	responses := make(chan int)
	for _, address := range addresses {
		go func(a string) { responses <- request(a) }(address)
	}
	return <-responses
}

// TestSyncTestLeak shows how synctest.Run() handles a goroutine leak
func TestSyncTestLeak(t *testing.T) {
	synctest.Run(func() {
		v := GetFirstResponse("ntp.gov", "time.google.com", "au.pool.ntp.org")
		if v < 0 {
			t.FailNow()
		}
	})
}

// An example of goroutine leak from:
// https://alenkacz.medium.com/an-example-of-a-goroutine-leak-and-how-to-debug-one-a0697cf677a3

func waitReady() bool {
	deadlineChan := time.NewTimer(time.Duration(1) * time.Second).C
	checkPodTicker := time.NewTicker(500 * time.Millisecond)
	doneChan := make(chan bool)

	defer checkPodTicker.Stop()

	go func() {
		for range checkPodTicker.C {
			verified := verify()
			if verified {
				doneChan <- true
				break
			}
		}
	}()

	for {
		select {
		case <-deadlineChan:
			return false
		case <-doneChan:
			return true
		}
	}
}

// instead of this run a code that will periodically verify something
func verify() bool {
	time.Sleep(10 * time.Second)
	return true
}

func TestWaitReady(t *testing.T) {
	synctest.Run(func() {
		waitReady()
	})
}
