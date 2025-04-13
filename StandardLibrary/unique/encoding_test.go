package __

import (
	"encoding/json"
	"fmt"
	"testing"
	"unique"
)

// also see StandardLibrary/JSON_YAML_marshal_etc/unique_handle_test.go:16

func TestUniqueJSONEncoding(t *testing.T) {
	r, err := json.Marshal(unique.Make("abc"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(r))

	var v unique.Handle[string]
	json.Unmarshal(r, &v)
	fmt.Printf("%T %q\n", v, v.Value())
}
