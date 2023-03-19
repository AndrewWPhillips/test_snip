package __

import (
	"log"
	"testing"
	"time"
)

// TestRace has a deliberate data race on b
func TestRace(t *testing.T) {
	const delay = time.Millisecond
	b := byte(1)
	go func() {
		b = 2
	}()
	time.Sleep(delay)
	log.Println(b)
}
