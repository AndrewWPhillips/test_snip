package __

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"reflect"
	"testing"
)

type Mine int

// TestMyType encodes a user-defined type in gob then decodes
func TestMyType(t *testing.T) {
	var m Mine = 42
	var buf bytes.Buffer
	buf.Grow(1024)

	// Encode my type
	enc := gob.NewEncoder(&buf)
	enc.Encode(m)
	enc = nil

	var value Mine = 1
	dec := gob.NewDecoder(&buf)
	dec.Decode(&value)

	println("Mine", value)
}

// TestArbitraryType tests if gob decoding works out the type (it doesn't)
func TestArbitraryType(t *testing.T) {
	var m Mine = 42
	var buf bytes.Buffer
	buf.Grow(1024)

	// Encode my type
	enc := gob.NewEncoder(&buf)
	enc.Encode(m)
	enc = nil

	var value any
	dec := gob.NewDecoder(&buf)
	dec.Decode(&value)

	switch v := value.(type) {
	case Mine:
		println("Mine", v)

	case int:
		println("Int", v)

	default:
		println("Unknown type", v) // always this
	}
}

// TestDecodeReflect checks if DecodeValue works (it doesn't)
func TestDecodeReflect(t *testing.T) {
	var m int = 42
	var buf bytes.Buffer
	buf.Grow(1024)

	// Encode my type
	enc := gob.NewEncoder(&buf)
	enc.Encode(m)
	enc = nil

	var value reflect.Value
	dec := gob.NewDecoder(&buf)
	dec.DecodeValue(value)

	// panic (type is always nil)
	println(value.Type())
	println(value.Kind())

	v := value.Interface()
	fmt.Printf("value %v type %T", v, v)
}

// TestStoreType uses an int before each value to store its type
func TestStoreType(t *testing.T) {
	var m Mine = 42
	var buf bytes.Buffer
	buf.Grow(1024)

	var typ int8
	enc := gob.NewEncoder(&buf)

	typ = 1 // int
	enc.Encode(typ)
	enc.Encode(1)
	typ = 2 // Mine
	enc.Encode(typ)
	enc.Encode(m)
	enc = nil

	var i int
	dec := gob.NewDecoder(&buf)
	for {
		if err := dec.Decode(&typ); err != nil {
			break
		}
		switch typ {
		case 1:
			dec.Decode(&i)
			println(i)
		case 2:
			dec.Decode(&m)
			println(m)
		}
	}
}
