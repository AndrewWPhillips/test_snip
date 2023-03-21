package __

import (
	"log"
	"testing"
	"time"
)

// TestTimeAfter tests using time.After to break from a for-select loop. See also TestAfter
func TestTimeAfter(t *testing.T) {
	ch := make(chan int)
	go func() {
		ch <- 1
		time.Sleep(time.Second)
		ch <- 2
	}()

loop:
	for {
		select {
		// Note that time.After() returns a buffered channel of size 1 so that even if the channel is never
		// read (eg: case <- ch is taken in switch below) the sender will not block and cause a leak.
		case <-time.After(500 * time.Microsecond):
			log.Println("After")
			break loop

		case i := <-ch:
			log.Println(i)
		}
	}
}
