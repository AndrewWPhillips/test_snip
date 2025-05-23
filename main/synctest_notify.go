//go:build go1.24 && goexperiment.synctest

package main

import (
	"os"
	"os/signal"
	"testing/synctest" // requires GOEXPERIMENT=synctest
)

// calling signal.Notify in synctest bubble hangs in Windows, panics in Linux (Go 1.24.2)

func main() {
	synctest.Run(func() {
		//signal.NotifyContext(context.Background())
		signal.Notify(make(chan os.Signal, 1))
	})
}
