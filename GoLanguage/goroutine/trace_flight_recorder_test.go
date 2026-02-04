//go:build go1.25

package __

import (
	"os"
	"runtime/trace"
	"sync/atomic"
	"testing"
	"time"
)

var global atomic.Int64

// simpleLoad just uses the CPU - use 4e9 for about 1 second of processing on my machine
func simpleLoad(load int) {
	var i int64
	for range load {
		i++
	}
	global.Store(i) // preclude the above loop from being optimised away
}

// gcLoad puts some load on the garbage collector - use 1e8 for more than 1 second of load (my m/c)
func gcLoad(load int) {
	var p *int64
	for i := range load {
		p = new(int64)
		*p = int64(i)
	}
	global.Store(*p) // preclude the above loop from being optimised away
}

func BenchmarkLoad(b *testing.B) {
	for b.Loop() {
		gcLoad(1e8)
	}
}

func TestTraceFlightRecorder(t *testing.T) {
	ptfr := trace.NewFlightRecorder(trace.FlightRecorderConfig{MaxBytes: 10e6})
	if ptfr.Start() != nil {
		t.Fatal("Could not start trace flight recorder")
	}
	defer ptfr.Stop()

	go simpleLoad(4e9)
	go simpleLoad(1e9)
	go gcLoad(1e8)
	time.Sleep(time.Second)

	file, err := os.Create("trace_" + time.Now().UTC().Format("060102150405") + ".out")
	if err != nil {
		t.Fatal("Could not create trace file:", err.Error())
	}
	_, err = ptfr.WriteTo(file)
	if err != nil {
		t.Fatal("Could not write trace file:", err.Error())
	}
	file.Close()
}
