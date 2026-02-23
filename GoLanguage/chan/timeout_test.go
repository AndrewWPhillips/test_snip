package _chan

import (
	"fmt"
	"testing"
	"time"
)

func SendChanTest(ch chan<- int) {
	ch <- 1
	time.Sleep(time.Millisecond)
	ch <- 2
	time.Sleep(time.Millisecond * 3)
	ch <- 3
	time.Sleep(time.Millisecond)
	ch <- 4
	time.Sleep(time.Second)
	close(ch)
}

// TestChanTimeout demonstrates how to use a timeout when reading from a channel
// Prints: 1 2 3 4
func TestChanNoTimeout(t *testing.T) {
	ch := make(chan int)
	go SendChanTest(ch)

	// loop reading from the channel till closed
	for v := range ch {
		fmt.Println(v)
	}
}

// TestChanTimeout demonstrates how to use a timeout when reading from a channel
// Prints: 1 2 timeout 3 4 timeout timeout timeout
func TestChanTimeout(t *testing.T) {
	ch := make(chan int)
	go SendChanTest(ch)

	// Loop 8 times reading from the channel with timeout
	for range 8 {
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-time.After(time.Millisecond * 2):
			fmt.Println("timeout")
		}
	}
}

// TestChanTimeoutClose demonstrates reading from a chan with timeout and handling close
// Prints: 1 2 timeout 3 4 timeout timeout timeout timeout timeout timeout ... 0
func TestChanTimeoutClose(t *testing.T) {
	ch := make(chan int)
	go SendChanTest(ch)

	// Loop reading from the channel with timeout until the chan is closed
	for ch != nil {
		select {
		case v, ok := <-ch:
			if !ok {
				ch = nil // indicate it's closed
			}
			fmt.Println(v)
		case <-time.After(time.Millisecond * 2):
			fmt.Println("timeout")
		}
	}
}
