package __

import (
	"testing"
	"time"
)

func send(ch chan bool) {
	time.Sleep(120 * time.Millisecond)
	ch <- true
	close(ch)
	println("Sent and closed")
}

func TestClose(t *testing.T) {
	ch := make(chan bool)
	timeout := time.After(600 * time.Millisecond)
	go send(ch)
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				return
			}
			println("Got message.")
		case <-timeout:
			println("Time out")
			return
		default:
			println("*yawn*")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
