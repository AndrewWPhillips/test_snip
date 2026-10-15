package utils

import (
	"context"
	"runtime"
	"testing"
	"testing/synctest"
	"time"
)

func MAXPROCSWatcher[T any](ctx context.Context, when <-chan T) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		cpus := runtime.GOMAXPROCS(-1)

		for {
			select {
			case <-ctx.Done():
				return
			case <-when:
				current := runtime.GOMAXPROCS(-1)
				if current != cpus {
					ch <- current
					cpus = current
				}
			}
		}
	}()
	return ch
}

// Test it stops when cancelled

func TestMAXPROCSWatcher(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ch := MAXPROCSWatcher(ctx, time.NewTicker(time.Second).C)
	go func() {
		time.Sleep(2 * time.Second)
		runtime.GOMAXPROCS(7)
		time.Sleep(2 * time.Second)
		runtime.GOMAXPROCS(200)
		time.Sleep(2 * time.Second)
		runtime.GOMAXPROCS(1)
		time.Sleep(2 * time.Second)
		runtime.GOMAXPROCS(2)
		time.Sleep(2 * time.Second)
		runtime.GOMAXPROCS(3)
	}()
	timer := time.NewTimer(5 * time.Second)
	for {
		select {
		case <-timer.C:
			cancel()
		case i, ok := <-ch:
			if !ok {
				return
			}
			println(i)
		}
	}
}

type (
	GoMAXPROCSWatcher struct {
		ch   chan int
		done chan struct{}
		cpus int
	}
)

func NewGoMAXPROCSWatcher() *GoMAXPROCSWatcher {
	return &GoMAXPROCSWatcher{
		ch:   make(chan int),
		done: make(chan struct{}),
		cpus: runtime.GOMAXPROCS(-1),
	}
}

func (w *GoMAXPROCSWatcher) Watch(ctx context.Context, dur time.Duration) <-chan int {
	go w.watch(ctx, dur)
	return w.ch
}

func (w *GoMAXPROCSWatcher) watch(ctx context.Context, dur time.Duration) {
	defer close(w.ch)
	ticker := time.NewTicker(dur)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-w.done:
			return
		case <-ticker.C:
			cpus := runtime.GOMAXPROCS(-1)
			if cpus == w.cpus {
				continue
			}
			w.cpus = cpus
			w.ch <- cpus
		}
	}
}

func (w *GoMAXPROCSWatcher) Stop() {
	close(w.done)
}

func TestGoMAXPROCSWatcher_Stop(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		watcher := NewGoMAXPROCSWatcher()
		ch := watcher.Watch(context.Background(), time.Nanosecond)

		synctest.Wait()
		select {
		case _, ok := <-ch:
			if !ok {
				t.Fatalf("watcher closed on running watcher")
			}
		default:
		}
		//time.Sleep(time.Second)

		watcher.Stop()
		synctest.Wait()

		if _, ok := <-ch; ok {
			t.Fatalf("watcher open on stopped watcher")

		}
	})
}
