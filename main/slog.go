package main

import (
	"errors"
	"log"
	"log/slog"
)

func main() {
	err := errors.New("TEST")
	log.Printf("test error=%q\n", err) // old logging
	slog.Info("test", "error", err)    // new logging
	slog.Info("test", slog.Any("error", err))
}
