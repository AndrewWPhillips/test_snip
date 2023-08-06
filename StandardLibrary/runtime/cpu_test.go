package __

import (
	"log"
	"runtime"
	"testing"
	"time"
)

// TestCpuCount tests calling runtime.NumCPU() while the number of available CPUs is adjusted (externally) using the
// OS (Linux or Windows) thread affinity capabilities.  Conclusion: return of NumCPU() does not change even when
// thread affinity is used.
func TestCpuCount(t *testing.T) {
	ticker := time.NewTicker(2 * time.Second)
	i := 0
	for {
		select {
		case <-ticker.C:
			log.Println(runtime.NumCPU())

		default:
			i++
		}
	}
}
