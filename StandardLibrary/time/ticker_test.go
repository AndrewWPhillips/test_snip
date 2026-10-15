package __

// tickers and timers

import (
	"testing"
	"time"
)

func TestNewTimer(t *testing.T) {
	<-time.NewTimer(time.Second).C
}

func TestTickerStop(t *testing.T) {
	ticker := time.NewTicker(100 * time.Millisecond)
	for range ticker.C {
		ticker.Stop() // does not close the ticker.C
		println("bye")
		break // otherwise deadlock
	}
	println("bye")
}
