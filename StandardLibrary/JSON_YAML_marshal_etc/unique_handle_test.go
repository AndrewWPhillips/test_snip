//go:build go1.23

// Test how handles from the unique package (Go 1.23) are encoded

package __

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"log"
	"testing"
	"unique"
)

type uniq struct {
	Handle  unique.Handle[string]
	String  string
	Pointer *string
}

func TestUniqueJSON(t *testing.T) {
	var s = "ABCD"

	got, err := json.Marshal(
		uniq{
			Handle:  unique.Make(s), // encodes as {}
			String:  s,
			Pointer: &s,
		})
	if err != nil {
		log.Fatalln("json.Marshall error:", err)
	}
	log.Printf("got JSON %q\n", got)

	var result uniq
	json.Unmarshal(got, &result)
	log.Printf("decoded to %v\n", result)
}

type (
	uniqJSON[T comparable] unique.Handle[T]

	toJSON struct {
		Handle  uniqJSON[string]
		String  string
		Pointer *string
	}
)

// MarshalJSON method mean uniqJSON implements json.Marshaler
func (uj uniqJSON[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(unique.Handle[T](uj).Value())
}

// UnmarshalJSON method means uniqJSON implements json.Unmarshaler
func (uj *uniqJSON[T]) UnmarshalJSON(data []byte) error {
	var v T
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*uj = uniqJSON[T](unique.Make[T](v))
	return nil
}

func TestUniqueCustomJSON(t *testing.T) {
	var s = "ABCD"

	got, err := json.Marshal(
		toJSON{
			Handle:  uniqJSON[string](unique.Make(s)),
			String:  s,
			Pointer: &s,
		})
	if err != nil {
		log.Fatalln("json.Marshall error:", err)
	}
	log.Printf("got JSON %q\n", got)

	var result toJSON
	json.Unmarshal(got, &result)
	log.Printf("decoded to %v\n", result)
}

func TestUniqueGOB(t *testing.T) {
	var s = "ABCD"
	var buf bytes.Buffer

	err := gob.NewEncoder(&buf).Encode(
		uniq{
			Handle:  unique.Make(s), // gob.Encoder.Encode error: gob: type unique.Handle[string] has no exported fields
			String:  s,
			Pointer: &s,
		})
	if err != nil {
		log.Fatalln("gob.Encoder.Encode error:", err)
	}
	log.Printf("got gob %q\n", buf.String())
}
