package __

import (
	"log"
	"sync/atomic"
	"testing"
	"time"
)

func TestModifyCapture(t *testing.T) {
	var i int32

	go func() {
		// Withing this i is "captured" as a reference to above i
		time.Sleep(200 * time.Millisecond)

		// Modify j (using atomic op as we are in a diff goroutine)
		atomic.StoreInt32(&i, 2)
	}()
	// print i, wait for change to j, then print i to show it has changed
	log.Println(atomic.LoadInt32(&i))
	time.Sleep(400 * time.Millisecond)
	log.Println(atomic.LoadInt32(&i))
}
