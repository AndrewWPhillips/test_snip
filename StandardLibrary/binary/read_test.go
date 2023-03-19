package __

import (
	crand "crypto/rand"
	"encoding/binary"
	"log"
	"testing"
)

func TestBinaryRead(t *testing.T) {
	a := struct {
		A int64
		B int64
	}{1, 2}
	if err := binary.Read(crand.Reader, binary.BigEndian, &a); err != nil {
		t.Fatal(err)
	}
	log.Println(a)
}
