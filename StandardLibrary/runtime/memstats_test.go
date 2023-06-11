package __

import (
	"log"
	"runtime"
	"testing"
)

func TestName(t *testing.T) {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	log.Printf("%v\n", ms)
}
