package __

// json_test.go has tests of json processing

import (
	"encoding/json"
	"log"
	"math"
	"testing"
	"time"
)

// TestUnmarshal tests decoding various JSON string into any type (interface{})
func TestUnmarshal(t *testing.T) {
	testData := map[string]struct {
		in string
	}{
		"bool":       {"true"}, // type = bool
		"int":        {"4"},    // type = float64
		"float":      {"3.14"}, // type = float64
		"null":       {"null"}, // type = nil
		"empty_obj":  {"{}"},   // type = map[string]interface{}
		"empty_list": {"[]"},   // type = []interface{}
		"map":        {`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia", null]}`},
		"test2":      {`{ "a": [1, true] }`},
		"test4":      {`{ "a": ["A", [2,3,4], null ] }`},
		"numberList": {"[-1, 42, 7]"},

		"bad":       {"?"},
		"bad_comma": {`{ "1": 1, }`}, // extra comma
		"bad_key":   {`{ 1: 1 }`},    // key must be a string
	}

	for name, data := range testData {
		t.Run(name, func(t *testing.T) {
			var result any
			if err := json.Unmarshal([]byte(data.in), &result); err != nil {
				log.Println(name, err) // t.Fail()
			} else {
				log.Printf("%s: %v {type = %T}\n", name, result, result)
			}
		})
	}
}

func TestUnmarshalArray(t *testing.T) {
	testData := map[string]struct {
		in string
	}{
		"empty": {"[]"},                // [0 0 0]
		"short": {"[42]"},              // [42 0 0]
		"same":  {"[-1, 42, 7]"},       // [-1 42 7]
		"long":  {"[-1, 42, 7, 8, 9]"}, // extras silently ignored
	}

	for name, data := range testData {
		t.Run(name, func(t *testing.T) {
			var result [3]int
			if err := json.Unmarshal([]byte(data.in), &result); err != nil {
				log.Println(name, err) // t.Fail()
			} else {
				log.Printf("%s: %v\n", name, result)
			}
		})
	}
}

func TestMarshal(t *testing.T) {
	type mt struct {
		S string `json:",omitempty"`
	}
	type jt struct {
		N float64
		M float64 `json:",omitempty"`
	}
	type cost struct {
		CostUSD string `json:"cost $"`
		CostEUR string `json:"cost €"` // This contains the non-ASCII punctuation character
	}
	type raw struct {
		R json.RawMessage
		S string
	}

	testData := map[string]struct {
		data any
		exp  string
	}{
		"empty":  {data: nil, exp: `null`},
		"string": {data: "error", exp: `"error"`},
		"int":    {data: 42, exp: `42`},
		// Test some "zero" (unassigned) values
		"pointer0": {data: (*int)(nil), exp: `null`},
		"time0":    {data: time.Time{}, exp: `"0001-01-01T00:00:00Z"`}, // default time is not null or {}
		// An empty slice and a nil slice are encoded differently
		"nil_slice":   {data: []int(nil), exp: `null`},
		"empty_slice": {data: []int{}, exp: `[]`},
		"slice":       {data: []int{1, 2}, exp: `[1,2]`},
		// A struct with all empty fields is not encoded as null
		"mt_struct0":     {data: mt{}, exp: `{}`},
		"mt_struct":      {data: mt{S: "a"}, exp: `{"S":"a"}`},
		"nil_struct_ptr": {data: (*mt)(nil), exp: `null`},

		"struct0": {data: jt{}, exp: `{"N":0}`}, // only has fields without omitempty
		"struct":  {data: jt{-1.99999, 1.99999}, exp: `{"N":-1.99999,"M":1.99999}`},
		"rounded": {
			data: jt{-1.9999999999999999999999999, 1.99999999999999999999999999999999999},
			exp:  `{"N":-2,"M":2}`,
		},
		"trailing_zeroes": {
			// obviously json.Marshall can't know of trailing zeroes of compile-time literal
			data: jt{123.0, 456.10000000000000000},
			exp:  `{"N":123,"M":456.1}`,
		},
		"inf": {data: priceType(math.NaN()), exp: `null`},
		//"nan": {
		//	data: jt{math.NaN(), math.NaN()}, exp: `{"N":???,"M":???}`,
		//}, // runtime error = json: unsupported value: NaN
		//"inf": {
		//	data: jt{math.Inf(1), math.Inf(-11)}, exp: `{"N":???,"M":???}`,
		//}, // runtime error = json: unsupported value: +Inf

		"map":          {data: map[string]string{"a": "b"}, exp: `{"a":"b"}`},
		"int_key_map":  {data: map[int]string{1: "b"}, exp: `{"1":"b"}`}, // int keys are converted to json strings
		"time_key_map": {data: map[time.Time]string{time.Unix(42, 0).UTC(): "b"}, exp: `{"1970-01-01T00:00:42Z":"b"}`},

		// A few oddities
		"non_ascii_tag": {data: cost{}, exp: `{"cost $":"","CostEUR":""}`}, // tag with € is ignored
		// Marshall escapes special HTML chars which avoids problems in some browsers - see SetEscapeHTML
		"html": {data: "Q&A", exp: `"Q\u0026A"`}, // avoid possible display as "Q&amp;A"
		"raw":  {data: raw{R: json.RawMessage(`"a"`), S: `"a"`}, exp: `{"R":"a","S":"\"a\""}`},
		"raw2": {
			data: raw{R: json.RawMessage(`{"a": "b"}`), S: `{"a": "b"}`},
			exp:  `{"R":{"a":"b"},"S":"{\"a\": \"b\"}"}`,
		},
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
