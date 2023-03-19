package __

import (
	"io"
	"strings"
	"testing"
)

func TestNopCloser(t *testing.T) {
	const input = "abc"
	var r io.Reader = strings.NewReader(input)
	var rc io.ReadCloser = io.NopCloser(r)

	b := make([]byte, 2)
	n, err := rc.Read(b)
	if err != nil || n != 3 {
		t.FailNow()
	}
	err = rc.Close()
	if err != nil {
		t.FailNow()
	}
}
