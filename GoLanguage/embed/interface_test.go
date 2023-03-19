package __

import (
	"io"
	"log"
	"strings"
	"testing"
)

func TestInterface(t *testing.T) {
	// Can call methods of an interface if an element of the interface type is embedded
	var a = struct {
		b string
		io.Reader
	}{"abc", strings.NewReader("def")}

	c := [10]byte{}
	d, e := a.Read(c[:]) // calls Read() on io.Reader interface which is a strings.Reader.Read()

	log.Println(a, c, d, e) // {abc <ptr>} [100 101 102 0 0 0 0 0 0 0] 3 nil
}
