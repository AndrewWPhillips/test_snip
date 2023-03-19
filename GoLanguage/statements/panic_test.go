package __

import (
	"fmt"
	"log"
	"testing"
)

func TestPanicInPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			// Convert the panic to a string, if possible, and log it
			switch v := r.(type) {
			case string:
				log.Printf("s %q\n", v)
			case error:
				log.Printf("e %q\n", v.Error())
			case fmt.Stringer:
				log.Printf("x %q\n", v.String())
			}
			// rethrow it
			panic(r)
		}
	}()

	panic("panic_text")
}
