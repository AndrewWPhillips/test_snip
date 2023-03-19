package __

import (
	"log"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestGoID(t *testing.T) {
	go func() {
		b := make([]byte, 64)
		b = b[:runtime.Stack(b, false)]
		s := strings.TrimLeft(strings.TrimPrefix(string(b), "goroutine"), " ")
		end := strings.IndexFunc(s, func(r rune) bool { return r < '0' || r > '9' })
		if end <= 0 {
			t.FailNow()
		}
		id, err := strconv.Atoi(s[:end])
		if err != nil {
			t.FailNow()
		}
		log.Println(id) // Go-routine ID
	}()
	time.Sleep(1 * time.Second)
}
