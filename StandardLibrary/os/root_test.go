//go:build go1.24

package __

import (
	"os"
	"testing"
)

func TestOSRoot(t *testing.T) {
	r, err := os.OpenRoot("E:\\work\\tmp\\Daily")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()

	//f, err2 := r.Open("E:\\work\\tmp\\Daily\\go.mod") // path escapes from parent
	f, err2 := r.Open("go.mod")
	if err2 != nil {
		t.Fatal(err2)
	}
	defer f.Close()

	buf := make([]byte, 1024)
	n, err3 := f.Read(buf)
	if err3 != nil {
		t.Fatal(err3)
	}
	println(n)
}
