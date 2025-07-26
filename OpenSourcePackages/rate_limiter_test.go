//go:build go1.24

package __

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"testing"
	"testing/synctest"
	"time"
)

func TestLimiter(t *testing.T) {
	synctest.Run(func() {
		const bucketSize = 5
		const limit = 5 // 5 per second
		//lim := rate.NewLimiter(1, bucketSize)
		lim := rate.NewLimiter(limit, bucketSize)
		start := time.Now()
		ctx := context.Background()

		// Empty the bucket then wait for 1 more
		for range bucketSize + 1 {
			lim.Wait(ctx)
		}
		took := time.Since(start)
		if took < time.Second/limit {
			t.Fail()
		}
		fmt.Println(took)
	})
}
