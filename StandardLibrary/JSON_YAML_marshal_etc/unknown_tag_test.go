//go:build go1.25 && goexperiment.jsonv2

package __

import (
	json "encoding/json/v2"
	"fmt"
	"testing"
)

const unknownTagTestJSON = `{
	"name": "cake",
	"ingredients": ["eggs", "flour"],
	"image": {
		"src": "images/cake1.jpeg"
	},
	"id": 12345
}`

type (
	recipe struct {
		Name string
	}

	recipe2 struct {
		Name  string
		Other map[string]any `json:",unknown"`
	}
)

func TestDecodeToStruct(t *testing.T) {
	var v struct{ Name string }
	if err := json.Unmarshal([]byte(unknownTagTestJSON), &v); err != nil {
		t.Error(err)
	}
	v.Name = "simple cake"

	if out, err := json.Marshal(v); err != nil {
		t.Error(err)
	} else {
		fmt.Println(string(out))
	}
}

func TestDecodeToAny(t *testing.T) {
	var v any
	if err := json.Unmarshal([]byte(unknownTagTestJSON), &v); err != nil {
		t.Error(err)
	}
	fmt.Printf("%#v %T\n", v, v.(map[string]any)["id"])
	v.(map[string]any)["name"] = "simple cake"

	if out, err := json.Marshal(v); err != nil {
		t.Error(err)
	} else {
		fmt.Println(string(out))
	}
}

func TestDecodeUnknown(t *testing.T) {
	var v recipe2
	if err := json.Unmarshal([]byte(unknownTagTestJSON), &v); err != nil {
		t.Error(err)
	}
	v.Name = "simple cake"

	if out, err := json.Marshal(v); err != nil {
		t.Error(err)
	} else {
		fmt.Println(string(out))
	}
}
