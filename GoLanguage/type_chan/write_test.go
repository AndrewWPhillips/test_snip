package __

import (
	"errors"
	"fmt"
	"log"
	"testing"
)

func TestWriteClosed(t *testing.T) {
	f := func() (err error) {
		defer func() {
			if e := recover(); e != nil {
				switch e := e.(type) {
				case error:
					err = fmt.Errorf("Panic with error: %w", e)
				case fmt.Stringer:
					err = errors.New(e.String())
				default:
					err = errors.New("panic sending on chanel")
				}
			}
		}()
		ch := make(chan int, 3)
		ch <- 1
		close(ch)
		ch <- 2
		return
	}
	log.Println(f())
}
