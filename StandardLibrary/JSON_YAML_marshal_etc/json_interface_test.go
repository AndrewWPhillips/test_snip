package __

import (
	"encoding/json"
	"log"
	"strings"
	"testing"
)

// Tests of unmarshalling into interface variables, incl. map and slice

func showJsonValue(indent int, jvalue interface{}) {
	switch value := jvalue.(type) {
	case map[string]interface{}:
		log.Println("<map>")
		showJsonMap(indent+1, value)

	case []interface{}:
		log.Println("<list>")
		showJsonList(indent+1, value)

	case string:
		log.Printf("<string> %q\n", value)

	case float64:
		log.Println("<float64>", value)

	case bool:
		log.Println("<bool>", value)

	case nil:
		log.Println("nil")

	default:
		log.Printf("<%T> %v\n", jvalue, jvalue)
	}
}

func showJsonList(indent int, jlist []interface{}) {
	log.Printf("%*s[\n", indent*4, "")
	for _, v := range jlist {
		log.Printf("%*s", (indent+1)*4, "")
		showJsonValue(indent, v)
	}
	log.Printf("%*s]\n", indent*4, "")
}

func showJsonMap(indent int, jmap map[string]interface{}) {
	log.Printf("%*s{\n", indent*4, "")
	for k, v := range jmap {
		log.Printf(`%*s"%s": `, (indent+1)*4, "", k) // show key for this element
		showJsonValue(indent, v)
	}
	log.Printf("%*s}\n", indent*4, "")
}

func TestValue(t *testing.T) {
	j := []byte(`"Name"`)
	var v interface{}
	err := json.Unmarshal(j, &v)
	if err == nil {
		showJsonValue(0, v)
	}
}

func TestNumericValue(t *testing.T) {
	j := `6`
	var v interface{}
	err := json.Unmarshal([]byte(j), &v)
	if err == nil {
		showJsonValue(0, v)
	}
}

func TestIntValue(t *testing.T) {
	j := `6`
	var v interface{}
	decoder := json.NewDecoder(strings.NewReader(j))
	decoder.UseNumber() // allows us to distinguish ints from floats (see FixNumberVariables() below)

	err := decoder.Decode(&v)
	if err != nil {
		t.Fatal(err)
	}
	showJsonValue(0, v)

	n, err2 := v.(json.Number).Int64()
	if err2 != nil {
		t.Fatal(err2)
	}
	log.Println(n)
}

func TestFloatValue(t *testing.T) {
	j := `6.666`
	var v interface{}
	decoder := json.NewDecoder(strings.NewReader(j))
	decoder.UseNumber() // allows us to distinguish ints from floats (see FixNumberVariables() below)

	err := decoder.Decode(&v)
	if err != nil {
		t.Fatal(err)
	}
	showJsonValue(0, v)

	n, err2 := v.(json.Number).Float64()
	if err2 != nil {
		t.Fatal(err2)
	}
	log.Println(n)
}

func TestMap(t *testing.T) {
	j := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia", null]}`)
	var v interface{}
	err := json.Unmarshal(j, &v)
	if err == nil {
		showJsonMap(0, v.(map[string]interface{}))
	}
}

func TestSlice(t *testing.T) {
	j := []byte(`[1, 2, "abc"]`)
	var v interface{}
	if err := json.Unmarshal(j, &v); err != nil {
		log.Println("json.Unmarshal error:", err)
		return
	}
	showJsonValue(0, v)
}
