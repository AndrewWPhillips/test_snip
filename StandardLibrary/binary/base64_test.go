package __

import (
	"bytes"
	"encoding/base64"
	"log"
	"testing"
)

func TestEncodeBase64(t *testing.T) {
	data := []byte{35, 9, 98, 78, 34, 155, 1, 2, 3, 5}

	s := base64.StdEncoding.EncodeToString(data)
	log.Println(s)

	out, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Compare(out, data) != 0 {
		t.Fatal("Round trip base64 encode and decode failed")
	}
}
