package __

import (
	"log"
	"testing"
)

func TestNestedMap(t *testing.T) {
	var msmss map[string]map[string]string
	msmss = map[string]map[string]string{"m": {"a": "b", "c": "d"}}
	log.Println(msmss)
}
