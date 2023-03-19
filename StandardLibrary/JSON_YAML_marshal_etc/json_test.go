package __

// json_test.go has tests of json processing

import (
	"encoding/json"
	"log"
	"math"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	testData := map[string]struct {
		in string
	}{
		"bad":        {"?"},
		"empty":      {"{}"},
		"map":        {`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia", null]}`},
		"test2":      {`{ "a": [1, true] }`},
		"test4":      {`{ "a": ["A", [2,3,4], null ] }`},
		"emptyList":  {"[]"},
		"numberList": {"[-1, 42, 7]"},
		"bad2":       {`{ "1": 1, }`},
	}

	for name, data := range testData {
		t.Run(name, func(t *testing.T) {
			var result interface{}
			if err := json.Unmarshal([]byte(data.in), &result); err != nil {
				log.Println(name, err) // t.Fail()
			} else {
				log.Printf("%s: %v %t\n", name, result, result)
			}
		})
	}
}

func TestMarshal(t *testing.T) {
	type jt struct {
		N float64
		M float64 `json:",omitempty"`
	}

	testData := map[string]struct {
		data interface{}
		exp  string
	}{
		"empty":   {data: nil, exp: `null`},
		"string":  {data: "error", exp: `"error"`},
		"int":     {data: 42, exp: `42`},
		"slice":   {data: []int{1, 2}, exp: `[1,2]`},
		"struct0": {data: jt{}, exp: `{"N":0}`},
		"struct":  {data: jt{-1.99999, 1.99999}, exp: `{"N":-1.99999,"M":1.99999}`},
		"rounded": {
			data: jt{-1.9999999999999999999999999, 1.99999999999999999999999999999999999},
			exp:  `{"N":-2,"M":2}`,
		},
		"inf": {data: priceType(math.NaN()), exp: `null`},
		//"nan": {
		//	data: jt{math.NaN(), math.NaN()}, exp: `{"N":???,"M":???}`,
		//}, // runtime error = json: unsupported value: NaN
		//"inf": {
		//	data: jt{math.Inf(1), math.Inf(-11)}, exp: `{"N":???,"M":???}`,
		//}, // runtime error = json: unsupported value: +Inf
		"map": {data: map[string]string{"a": "b"}, exp: `{"a":"b"}`},
	}

	for name, data := range testData {
		t.Run(name, func(t *testing.T) {
			b, err := json.Marshal(data.data)
			if err != nil {
				t.Fatalf("%10s: marshall failed with error %v\n", name, err)
			}
			if string(b) != data.exp {
				t.Fatalf("%10s: expected %s but got %s\n", name, data.exp, string(b))
			}
		})
	}
}

type priceType float64

func (p priceType) MarshalJSON() ([]byte, error) {
	f := float64(p)
	if math.IsNaN(f) {
		return []byte("null"), nil
	}
	return json.Marshal(f)
}

func TestCustomMarshal(t *testing.T) {
	b, err := json.Marshal(priceType(1.999))
	if err == nil {
		log.Println("<", string(b), ">")
	} else {
		log.Println(err)
	}
	b, err = json.Marshal(priceType(math.NaN()))
	if err == nil {
		log.Println("<", string(b), ">")
	} else {
		log.Println(err)
	}
}
