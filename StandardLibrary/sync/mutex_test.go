package __

import (
	"sync"
	"testing"
)

func TestMutexCompare(t *testing.T) {
	var m, n sync.Mutex

	println(m == n)
}
