//go:build go1.4

package __

import (
	"fmt"
	"testing"
	"testing/synctest" // requires GOEXPERIMENT=synctest
	"time"
)

func TestSyncTest(t *testing.T) {
	synctest.Run(func() {
		fmt.Println(time.Now())
	})
}
