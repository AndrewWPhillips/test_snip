package __

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestStringReader(t *testing.T) {
	const input = "abc"
	b, err := io.ReadAll(strings.NewReader(input))
	if err != nil || string(b) != input {
		t.FailNow()
	}
}

func TestBufReader(t *testing.T) {
	var buf io.Reader = bytes.NewBuffer([]byte{'a', 'b', 'c'})
	b, err := io.ReadAll(buf)
	if err != nil || bytes.Compare(b, []byte{'a', 'b', 'c'}) != 0 {
		t.FailNow()
	}
}
