//go:build go1.24 && goexperiment.synctest

package __test

import (
	"fmt"
	"math/rand/v2"
	"testing"
	"testing/synctest"
	"time"
)

func query(address string) int {
	// simulate delay of 1 to 5 seconds
	time.Sleep(time.Duration(rand.IntN(4)+1) * time.Second)
	return len(address)
}

func fastestQuery(addresses ...string) int {
	ch := make(chan int)
	for _, address := range addresses {
		go func() {
			ch <- query(address)
		}()
	}
	return <-ch
}

func TestFastest(t *testing.T) {
	synctest.Run(func() {
		start := time.Now()
		_ = fastestQuery("ntp.org", "ntp.org.au", "time.google.com")
		elapsed := time.Since(start)
		if elapsed < time.Second || elapsed > time.Hour {
			t.FailNow()
		}
		fmt.Printf("took %v\n", elapsed)
	})
}

func TestAnotherDeadlock(t *testing.T) {
	synctest.Run(func() {
		ch := make(chan struct{})
		go func() {
			<-ch
		}()
	})
}
