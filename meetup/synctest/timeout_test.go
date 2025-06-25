//go:build go1.24

package __test

import (
	"context"
	"testing"
	"testing/synctest"
	"time"
)

// nothing just returns when the context is done
func nothing(ctx context.Context) {
	<-ctx.Done()
	//for range ctx.Done() { /* nothing here */
}

func TestTimeoutQqq(t *testing.T) {
	//synctest.Run(func() {
	delay := 2 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), delay)
	defer cancel()

	start := time.Now()
	nothing(ctx)
	got := time.Since(start)
	//if got < delay {
	if got != delay {
		t.Fatalf("took %v but expected %v", got, delay)
	}
	//})
}

func TestTimeoutTable(t *testing.T) {
	table := map[string]struct {
		delay time.Duration
	}{
		"simple": {delay: 2 * time.Second},
		"day":    {delay: 24 * time.Hour},
	}
	for name, test := range table {
		t.Run(name, func(t *testing.T) {
			synctest.Run(func() {
				ctx, cancel := context.WithTimeout(context.Background(), test.delay)
				defer cancel()
				start := time.Now()
				nothing(ctx)
				got := time.Since(start)
				if got != test.delay {
					t.Fatalf("took %v but expected %v", got, test.delay)
				}
			})
		})
	}
}
