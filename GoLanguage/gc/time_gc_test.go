package __

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestTimeGC(t *testing.T) {
	const MAX = int(1e9)
	//var stats debug.GCStats
	var ms runtime.MemStats

	//debug.ReadGCStats(&stats)
	//fmt.Println(stats)
	runtime.ReadMemStats(&ms)
	fmt.Println(ms)

	all := make([]time.Time, 0, MAX)
	for i := 0; i < MAX; i++ {
		all = append(all, time.Now().UTC())
	}
	//debug.ReadGCStats(&stats)
	//fmt.Println(stats)
	runtime.ReadMemStats(&ms)
	fmt.Println(ms)

	runtime.GC()
	//debug.ReadGCStats(&stats)
	//fmt.Println(stats)
	runtime.ReadMemStats(&ms)
	fmt.Println(ms)
}
