package __

import (
	"context"
	"testing"
	"time"
)

func TestAfterFunc(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	calledCh := make(chan struct{}) // closed when AfterFunc is called
	context.AfterFunc(ctx, func() {
		close(calledCh)
	})

	if funcCalled(t, calledCh) {
		t.Fatalf("AfterFunc function called before context is canceled")
	}

	cancel()

	if !funcCalled(t, calledCh) {
		t.Fatalf("AfterFunc function not called after context is canceled")
	}
}

func funcCalled(t *testing.T, calledCh <-chan struct{}) bool {
	select {
	case <-calledCh:
		return true
	case <-time.After(10 * time.Nanosecond):
		return false
	}
}
