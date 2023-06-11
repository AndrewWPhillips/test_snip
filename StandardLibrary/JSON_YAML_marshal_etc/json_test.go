package __

// json_test.go has tests of json processing

import (
	"encoding/json"
	"log"
	"math"
	"reflect"
	"testing"
	"time"

	"github.com/dolmen-go/jsonmap"
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
		"full_list":  {`[{"Response": "1b"}, {"Response": "i", "Data": {"Contest": "203244481","BetType": 0 }}]`},
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

func fixup(buf []byte) []byte {
	// Decode the JSON, remove some messages and then re-encode into tmp
	var messages []jsonmap.Ordered
	if err := json.Unmarshal(buf, &messages); err != nil {
		log.Println("ERROR: fixup unmarshal:", err)
		return buf
	}
	foundBetTypes := make(map[int64]struct{})
	for i, message := range messages {
		if message.Data["Response"] == "i" {
			// we found a timestamp message - check if we have seen it before for the bet type
			if d, ok := message.Data["Data"].(map[string]any); ok {
				var betType int64
				if number, ok := d["BetType"].(json.Number); ok {
					betType, _ = number.Int64()
				} else if f, ok := d["BetType"].(float64); ok {
					betType = int64(f)
				}
				if betType == 0 { // not found or Win bet type
					continue
				}
				if _, found := foundBetTypes[betType]; found {
					// we've already seen a timestamp message for this bet type so remove it
					messages[i] = jsonmap.Ordered{}
				} else {
					// remember that we have seen an "i" message for this bet type
					foundBetTypes[betType] = struct{}{}
				}
			}
		}
	}
	messagesCopy := make([]jsonmap.Ordered, 0, len(messages))
	for _, m := range messages {
		if m.Data != nil {
			messagesCopy = append(messagesCopy, m)
		}
	}
	newBuf, err := json.Marshal(messagesCopy)
	if err != nil {
		log.Println("ERROR: fixup marshal:", err)
		return buf
	}
	return newBuf
}

func TestFixup(t *testing.T) {
	tests := map[string]struct {
		messages, expected string
	}{
		"Win         ": {
			messages: `[{"Response":"i", "Data":{"BetType":1}}]`,
			expected: `[{"Response":"i", "Data":{"BetType":1}}]`,
		},
		"Place       ": {
			messages: `[{"Response":"i", "Data":{"BetType":1}}]`,
			expected: `[{"Response":"i", "Data":{"BetType":1}}]`,
		},
		"Win Place   ": {
			messages: `[{"Response":"i", "Data":{"BetType":0}}, {"Response":"i", "Data":{"BetType":1}}]`,
			expected: `[{"Response":"i", "Data":{"BetType":0}}, {"Response":"i", "Data":{"BetType":1}}]`,
		},
		"Two Wins    ": {
			messages: `[{"Response":"i", "Data":{"BetType":0}}, {"Response":"i", "Data":{"BetType":1}}]`,
			expected: `[{"Response":"i", "Data":{"BetType":0}}, {"Response":"i", "Data":{"BetType":1}}]`,
		},
		"WinWinPlace ": {
			messages: `[{"Response":"i", "Data":{"BetType":0}}, {"Response":"i", "Data":{"BetType":0}}, {"Response":"i", "Data":{"BetType":1}}]`,
			expected: `[{"Response":"i", "Data":{"BetType":0}}, {"Response":"i", "Data":{"BetType":1}}]`,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test := test

			// Encode the expected and received JSON so we can compare them reliably (ie. w/o map order problems)
			var expected, got []any
			if json.Unmarshal([]byte(test.expected), &expected) != nil {
				t.Fatal("error encoding expected JSON:", test.expected)
			}
			if json.Unmarshal(fixup([]byte(test.messages)), &got) != nil {
				t.Fatal("error encoding received JSON")
			}
			if !reflect.DeepEqual(expected, got) {
				t.Fatalf("%s: got %v but expected %v\n", name, got, expected)
			}
		})
	}
}

func TestUnmarshalBatch1(t *testing.T) {
	var messages []any
	batch1 := `[{"Response": "1b"}, {"Response": "i", "Data": {"Contest": "203244481","BetType": 0}}]`
	if err := json.Unmarshal([]byte(batch1), &messages); err != nil {
		t.Fatal(err)
	}
	for i, m := range messages {
		log.Printf("%d %T %v\n", i, m, m)
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
