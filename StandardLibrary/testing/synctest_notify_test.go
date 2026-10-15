package __

import (
	"os"
	"os/signal"
	"testing"
	"testing/synctest"
)

func TestSyncTestNotify(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		println("hello world")
		signal.Notify(make(chan os.Signal, 1))
		println("hello again")
	})
	println("bye") // never seen
}
