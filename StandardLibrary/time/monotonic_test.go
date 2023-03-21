package __

import (
	"fmt"
	"testing"
	"time"
)

// TestWallVsMonotonicDiff checks that wall time diff and monotonic time diff are the same to the nanosec
func TestWallVsMonotonicDiff(t *testing.T) {
	tMono := time.Now()
	tWall := tMono.Round(0) // turn off hasMonotonic flag
	<-time.After(1 * time.Second)
	diffWall := time.Since(tWall)
	diffMono := time.Since(tMono)
	if diffWall != diffMono {
		t.Errorf("diffWall = %v, diffMono = %v", diffWall, diffMono)
	}
}

// The following is a test of way to get a "monotonic" time value (see NowTimestamp)
var baseTime time.Time

func init() {
	baseTime = time.Now()
}

// NowTimestamp returns the monotonic time - really just the duration from a base time
// This is useful to:
//  - save memory 8 bytes (time.Duration) vs 24 bytes (time.Time)
//  - ensure it's truly monotonic (time.Time may have its hasMonotonic bit turned off inadvertently eg by call to UTC method)
func NowTimestamp() time.Duration {
	return time.Now().Sub(baseTime)
}

func TestMonotonic(t *testing.T) {
	fmt.Println(NowTimestamp())
	for i := 0; i < 1000; i++ {
		fmt.Println(int(NowTimestamp()))
	}
}

// timeMono encapsulates the above idea in a type (no global variable)
type timeMono struct {
	base time.Time
}

// NewMono creates a timeMono so that (if this is a package) the client cannot create a timeMono directly
func NewMono() timeMono {
	return timeMono{time.Now()}
}

func (t *timeMono) Now() time.Duration {
	return time.Now().Sub(t.base)
}
