package __

import (
	"log"
	"runtime"
	"testing"
)

// TestCaller uses the runtime package to obtain the source code line of the function being executed
func TestCaller(t *testing.T) {
	log.Println(runtime.Caller(0)) // 7096395 E:/Work/andrew/GoProjects/test_snip/StandardLibrary/runtime/caller_test.go 11 true  /// 277 == this line number
	log.Println(runtime.Caller(1)) // 6665519 C:/Program Files/Go1.17.10/src/testing/testing.go 1259 true
	//                                OR: 4377414 C:/Go/src/runtime/proc.go 185 true

	func() {
		log.Println(runtime.Caller(0)) // 7099019 E:/Work/andrew/GoProjects/test_snip/StandardLibrary/runtime/caller_test.go 15 true XXX this line
		log.Println(runtime.Caller(1)) // 7096793 E:/Work/andrew/GoProjects/test_snip/StandardLibrary/runtime/caller_test.go 18 true XXX 2 lines below this
		log.Println(runtime.Caller(2)) // 6665519 C:/Program Files/Go1.17.10/src/testing/testing.go 1259 true
	}()
}
