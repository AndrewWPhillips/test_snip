package __

import (
	"log"
	"testing"

	"github.com/mitchellh/mapstructure"
)

func TestStringMap(t *testing.T) {
	var d struct{ Name, Email string }

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{Result: &d})
	if err != nil {
		t.Fatal(err.Error())
	}

	err = decoder.Decode(map[string]string{"name": "Andrew", "email": "x@y"})
	if err != nil {
		t.Fatal(err.Error())
	}

	log.Printf("%+v", d)
}
