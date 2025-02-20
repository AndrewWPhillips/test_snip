package __

import (
	"sync"
	"testing"
)

func TestLock(t *testing.T) {
	mu := sync.Mutex{}
	mu.Lock()
	defer mu.Unlock()
}
