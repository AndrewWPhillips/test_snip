package __

import (
	"crypto/rand"
	"runtime"
	"testing"
	"time"
)

func TestHoldLargerString(t *testing.T) {
	var subString string
	{
		longString := rand.Text()
		runtime.AddCleanup(&longString, func(int) {
			println("longString cleanup")
		}, 0)
		subString = longString[1:2]
	}
	println("before GC")
	runtime.GC()
	time.Sleep(time.Microsecond) // allow for GC to finish
	println("after GC")
	runtime.KeepAlive(subString)
}
