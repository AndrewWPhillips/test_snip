package __

import (
	"log"
	"testing"
)

func TestSelectSame(t *testing.T) {
	ch := make(chan int, 5)
	h := 0
	ch <- h
	h++
	ch <- h
	h++

	for j := 0; j < 100; j++ {
		// You can send and receive on the same chan in a select
		select {
		case i := <-ch:
			log.Printf("recv: %v\n", i)
		case n := <-ch:
			log.Printf("recv= %v\n", n)
		case ch <- h:
			log.Printf("sent: %v\n", h)
			h++
		}
	}
}
