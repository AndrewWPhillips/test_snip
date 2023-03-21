package __

import (
	"log"
	"testing"
)

func TestMapOfMap(t *testing.T) {
	var msgmap map[string]map[string]int
	//msgmap["a"]["b"] = 42  // panic: assignment to entry in nil map
	msgmap = make(map[string]map[string]int)
	//msgmap["a"]["b"] = 42 // panic: assignment to entry in nil map
	msgmap["a"] = make(map[string]int)
	msgmap["a"]["b"] = 42 // OK

	log.Println(msgmap["a"]["b"])
}

func TestMapOfMapOfMap(t *testing.T) {
	type s struct {
		i int
	}

	/* this way causes panics
	var m map[string]map[string]map[string]s
	m["topic"] = make(map[string]map[string]s)  // panic
	m["topic"]["type"] = make(map[string]s)
	m["topic"]["type"]["stream_id"] = s{}
	*/

	strTopic := "topic" // map key need not be constant
	forty2 := 42

	m1 := map[string]s{strTopic: {i: forty2}}
	log.Println(m1["topic"].i)

	m2 := map[string]map[string]s{"topic": {"type": {i: 43}}}
	log.Println(m2["topic"]["type"].i)

	m3 := map[string]map[string]map[string]s{"topic": {"type": {"stream_id": {i: 44}}}}
	log.Println(m3["topic"]["type"]["stream_id"].i)

	//m3["topic"]["type"]["stream_id"].i = 47       // cannot assign to struct field in map
	var tmp = m3["topic"]["type"]["stream_id"] // workaround using tmp struct
	tmp.i = 47
	m3["topic"]["type"]["stream_id"] = tmp
	log.Println(m3["topic"]["type"]["stream_id"].i)

}
