package __

import (
	"encoding/json"
	"testing"
)

func TestKeep(t *testing.T) {
	testData := map[string]struct {
		in  string
		add []any
	}{
		"simple": {`{"xxx": "yyy", "groups": ["a", "b"]}`, []any{"c", "d"}},
	}

	for _, data := range testData {
		var v any

		err := json.Unmarshal([]byte(data.in), &v)
		if err != nil {
			t.Fatal("got error:" + err.Error())
		}
		m, ok := v.(map[string]any)
		if !ok {
			t.Fatal("m is not map[string]any")
		}
		g, ok2 := m["groups"]
		if !ok2 {
			t.Fatal("groups not found")
		}
		gg, ok3 := g.([]any)
		if !ok3 {
			t.Fatal("gg is not []any")
		}
		gg = append(gg, data.add...)
		m["groups"] = gg
		got, err2 := json.Marshal(v)
		if err2 != nil {
			t.Fatal("got error:" + err2.Error())
		}
		println(string(got))
	}
}
