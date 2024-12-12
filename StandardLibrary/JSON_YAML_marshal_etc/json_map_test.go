package __

import (
	"encoding/json"
	"log"
	"testing"
)

type (
	testMapKey   struct{ K string }
	testMapValue struct{ V string }
)

func TestUnmarshalMapKey(t *testing.T) {
	m := make(map[testMapKey]testMapValue)

	testData := map[string]struct {
		in string
	}{
		"map": {`{{"K": "Name"}: {"V": "Wednesday"}}`},
	}

	for name, data := range testData {
		t.Run(name, func(t *testing.T) {
			if err := json.Unmarshal([]byte(data.in), &m); err != nil {
				log.Println(name, err) // t.Fail()
			} else {
				log.Printf("%s: %v {type = %T}\n", name, m, m)
			}
		})
	}
}
