package utils

import (
	"log"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

func TestNumberOfCPU(t *testing.T) {
	log.Println(runtime.NumCPU(), runtime.NumGoroutine(), runtime.NumCgoCall())
}

// TestGoroutineID gets an ID (number) for a goroutine
func TestGoroutineID(t *testing.T) {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	//log.Println(strconv.Atoi(strings.TrimPrefix(string(b), "goroutine")))
	s := strings.TrimLeft(strings.TrimPrefix(string(b), "goroutine"), " \t")
	offset := strings.IndexFunc(s, func(r rune) bool { return r < '0' || r > '9' })
	log.Println(strconv.Atoi(s[:offset]))
}
