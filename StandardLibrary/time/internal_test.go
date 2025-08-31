package __

import (
	"fmt"
	"testing"
	"time"
	"unsafe"
)

type TimeExposed struct {
	// wall and ext encode the wall time seconds, wall time nanoseconds,
	// and optional monotonic clock reading in nanoseconds.
	//
	// From high to low bit position, wall encodes a 1-bit flag (hasMonotonic),
	// a 33-bit seconds field, and a 30-bit wall time nanoseconds field.
	// The nanoseconds field is in the range [0, 999999999].
	// If the hasMonotonic bit is 0, then the 33-bit field must be zero
	// and the full signed 64-bit wall seconds since Jan 1 year 1 is stored in ext.
	// If the hasMonotonic bit is 1, then the 33-bit field holds a 33-bit
	// unsigned wall seconds since Jan 1 year 1885, and ext holds a
	// signed 64-bit monotonic clock reading, nanoseconds since process start.
	Wall uint64
	Ext  int64

	// loc specifies the Location that should be used to
	// determine the minute, hour, month, day, and year
	// that correspond to this Time.
	// The nil location means UTC.
	// All UTC times are represented with loc==nil, never loc==&utcLoc.
	Loc *time.Location
}

// Show displays the fields of a time.Time for when mono bit is on
func Show(tm time.Time) {
	te := *(*TimeExposed)(unsafe.Pointer(&tm))
	if te.Wall>>63 == 0 {
		fmt.Printf("NO MONO wall:%d.%09d secs since year 1\n",
			te.Ext,
			te.Wall&0x3FFF_FFFF,
		)
	} else {
		fmt.Printf("MONO wall:%v secs, mono:%v.%09d secs since year 1885\n",
			(te.Wall>>30)&0x1_FFFF_FFFF, // wall secs
			te.Ext/1_000_000_000,        // mono clock (nsecs) -> secs
			te.Wall&0x3FFF_FFFF,         // ??
		)

	}
}

// TestTimeInternal checks time internal fields when using synctest
func TestTimeInternal(t *testing.T) {
	//synctest.Run(func() {
	tm := time.Now()
	Show(tm)
	Show(tm.UTC())

	time.Sleep(100_001 * time.Millisecond)
	Show(time.Now())
	Show(time.Now().UTC())
	//})
}
