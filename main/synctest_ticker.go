//go:build go1.24 && goexperiment.synctest

package main

import (
	"testing/synctest"
	"time"
)

func main() {
	synctest.Run(func() {
		_ = time.NewTicker(2 * time.Second)
	})
}
