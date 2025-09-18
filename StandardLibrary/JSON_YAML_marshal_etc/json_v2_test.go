//go:build go1.25 && goexperiment.jsonv2

package __

import (
	json "encoding/json/v2"
	"fmt"
	"testing"
)

func TestJSONV2(t *testing.T) {
	j := []byte(wedString)

	//var v interface{}
	//if err := json.Unmarshal(j, &v); err != nil {
	//	t.Fatal(err)
	//}
	//showJsonMap(0, v.(map[string]interface{}))

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
