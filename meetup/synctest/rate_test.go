//go:build go1.24

package __test

import (
	"context"
	"golang.org/x/time/rate"
	"testing"
	"testing/synctest"
	"time"
)

func TestRateLimiter(t *testing.T) {
	limiter := rate.NewLimiter(1, 5)
	ctx := context.Background()

	start := time.Now()
	for range 6 {
		limiter.Wait(ctx)
	}
	elapsed := time.Since(start)
	if elapsed != time.Second {
		t.Fatalf("unexpected elapsed time: %v", elapsed)
	}
}

func TestRateLimiterTable(t *testing.T) {
	var table = map[string]struct {
		tokens     int // how many tokens per second
		bucketSize int // how many tokens the bucket holds
		extra      int // how many extra tokens we wait for
	}{
		"simple": {1, 5, 2},
	}
	ctx := context.Background()
	for name, entry := range table {
		t.Run(name, func(t *testing.T) {
			synctest.Run(func() {
				limiter := rate.NewLimiter(rate.Limit(entry.tokens), entry.bucketSize)

				start := time.Now()
				for range entry.bucketSize + entry.extra {
					limiter.Wait(ctx)
				}
				expected := time.Duration(entry.extra) * time.Second
				got := time.Since(start)
				if got != expected {
					t.Fatalf("expected elapsed time: %v but got %v", expected, got)
				}
			})
		})
	}
}
