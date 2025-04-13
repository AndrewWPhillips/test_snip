package __

import (
	"crypto/rand"
	"golang.org/x/sync/errgroup"
	"sync/atomic"
	"testing"
	"unique"
)

func TestName(t *testing.T) {
	const repeat = 1_000_000
	var eg errgroup.Group
	var count atomic.Int32

	for range 10 {
		eg.Go(func() error {
			for range repeat {
				unique.Make(rand.Text())
			}
			return nil
		})
	}
	eg.Go(func() error {
		var length int
		for range repeat {
			h := unique.Make("abcdefghijklmnopqrstuvwxyz")
			length = len(h.Value())
		}
		count.Add(int32(length))
		return nil
	})
	eg.Wait()
	println(count.Load())
}
