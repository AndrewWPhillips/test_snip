package __

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"testing"
)

// TestTypeSwitch tests a type switch with scalar (string) and interface (error, Stringer) cases
func TestTypeSwitch(t *testing.T) {
	for _, a := range []interface{}{
		"abc", errors.New("def"), bytes.NewBufferString("ghi"), [...]int{1}, [...]int{1, 2},
	} {
		switch v := a.(type) {
		case string:
			log.Printf("s: %s", v)
		case error:
			log.Printf("e: %s", v.Error())
		case fmt.Stringer:
			log.Printf("fs: %s", v.String())
		case [2]int:
			log.Printf("a: %v", v)
		default:
			log.Printf("dv: %v", v)
		}
	}
}
