package __

import (
	"os"
	"testing"
)

func TestReadEscape(t *testing.T) {
	f, _ := os.Open("/tmp/foo")
	buf := make([]byte, 4096)
	n, _ := f.Read(buf)
	println(n)
}
