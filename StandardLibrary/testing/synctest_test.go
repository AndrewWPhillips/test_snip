//go:build go1.24

package __

// see also OpenSourcePackages/rate_limiter_test.go:13

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"testing"
	"testing/synctest" // requires GOEXPERIMENT=synctest
	"time"
)

// TestSyncTest just tests 2 goroutines and sleep
// synctest.Run() does not return until the 10 second goroutine exists
// The test runs in a few ms since time.Sleep() calls use the "fake" time.
func TestSyncTest(t *testing.T) {
	start := time.Now()
	synctest.Run(func() {
		go func() {
			fmt.Println("A:", time.Now())
			time.Sleep(10 * time.Second)
			fmt.Println("B:", time.Now())
		}()
		time.Sleep(time.Second)
		fmt.Println("C:", time.Now())
	})
	fmt.Println("took:", time.Now().Sub(start))
}

// TestSyncExternal tests interacting with goroutines running outside
// This test takes about 5 secs to run due to time.Sleep() calls using real time.
func TestSyncExternal(t *testing.T) {
	start := time.Now()

	// Create a channel and start a go-routine to write to it
	ch := make(chan struct{})
	go func() {
		for range 5 {
			time.Sleep(time.Second)
			ch <- struct{}{}
		}
		close(ch)
	}()

	synctest.Run(func() {
		open := true
		// loop till the channel is closed
		for open {
			// get the next one from the channel
			select {
			case _, open = <-ch:
			}
			fmt.Println(time.Now())
		}
	})
	fmt.Println("took:", time.Now().Sub(start))
}

func TestSyncInsideSync(t *testing.T) {
	println("Outside")
	synctest.Run(func() {
		println("Middle")
		synctest.Run(func() { // panic: synctest.Run called from within a synctest bubble
			println("Inside")
		})
	})
}

func TestTimeSleep(t *testing.T) {
	//synctest.Run(func() {
	before := time.Now()
	time.Sleep(time.Second)
	after := time.Now()
	if d := after.Sub(before); d != time.Second {
		t.Fatalf("took %v", d)
	}
	//})
}

func TestSynctestMono(t *testing.T) {
	synctest.Run(func() {
		before := time.Now()
		time.Sleep(5e18)
		after := time.Now()

		fmt.Println(before)
		fmt.Println(after)
		fmt.Println(before.Round(0))
		fmt.Println(after.Round(0))
	})
}

// TestSyncTestNotifyProblem tries (and fails) to reproduce problems seen in main/synctest_notify.go
func TestSyncTestNotifyProblem(t *testing.T) {
	_, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	synctest.Run(func() {
		start := time.Now()
		fmt.Println("Running starting at", start)
		time.Sleep(20 * time.Second)
		now := time.Now()
		duration := now.Sub(start)
		fmt.Println("duration: ", duration)
	})
}
