package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	//"os/signal"
	"testing/synctest" // requires GOEXPERIMENT=synctest
	"time"
)

const fakeTimeFormat = "2006-01-02 15:04:05"

var fakeTime = flag.String("fakeTime", "2030-01-02 00:00:00", "run using a fake time eg: "+fakeTimeFormat)

func main() {
	flag.Parse()
	if *fakeTime != "" {
		startTime, err := time.Parse(fakeTimeFormat, *fakeTime)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if startTime.Year() < 2000 {
			fmt.Fprintln(os.Stderr, "fake time must be this century")
			os.Exit(1)
		}
		synctest.Run(func() {
			time.Sleep(startTime.Sub(time.Now()))
			if err := sleep20secs(context.Background(), os.Stdout, os.Args); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		})
	} else {
		if err := sleep20secs(context.Background(), os.Stdout, os.Args); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}

func sleep20secs(ctx context.Context, out io.Writer, args []string) error {
	//ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	//defer cancel()

	start := time.Now()
	fmt.Println("Running starting at", start)
	time.Sleep(20 * time.Second)
	now := time.Now()
	duration := now.Sub(start)
	fmt.Fprintf(out, "duration: %s\n", duration)

	return nil
}
