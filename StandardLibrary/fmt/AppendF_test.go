package __

import (
	"fmt"
	"log"
	"testing"
)

func TestAppendF(t *testing.T) {
	var b []byte
	b = fmt.Appendf(b, "%d", 42)
	log.Print(string(b))
}
