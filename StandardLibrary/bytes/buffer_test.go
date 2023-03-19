package __test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestByteBuffer(t *testing.T) {
	//var b bytes.Buffer
	//b.WriteString("abc")
	//log.Println(b.String())

	var b bytes.Buffer
	b.Write([]byte("Hello"))         // add bytes to buffer
	fmt.Fprintf(&b, "%s", " world!") // formatted printing to buffer

	// write the buffer to an io.Writer
	var i io.Writer = os.Stderr
	_, _ = b.WriteTo(i) // os.Stdout implements io.Writer
}
