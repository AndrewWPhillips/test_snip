//go:build go1.24

package __

import (
	"encoding/json"
	"fmt"
	"testing"
)

type jsonStruct struct {
	Name  string
	Other map[string]any `json:",unknown"`
}

func TestJSONV1(t *testing.T) {
	j := []byte(wedString)

	// First decode it the old way when you don't know what the JSON contains
	var v interface{}
	if err := json.Unmarshal(j, &v); err != nil {
		t.Fatal(err)
	}
	showJsonMap(0, v.(map[string]interface{}))

	var s jsonStruct
	if err := json.Unmarshal(j, &s); err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v\n", s)

	if buf, err := json.Marshal(s); err != nil {
		t.Fatal(err)
	} else {
		fmt.Printf("%s\n", string(buf))
	}
}
