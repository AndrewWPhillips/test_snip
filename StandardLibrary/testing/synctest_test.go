//go:build go1.24

package __

import (
	"fmt"
	"testing"
	"testing/synctest" // requires GOEXPERIMENT=synctest
	"time"
)

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
