package __

import (
	"encoding/json/v2"
	"log"
	"testing"
	"time"
)

func TestMarshallTime(t *testing.T) {
	v1 := time.Now()
	d, _ := json.Marshal(v1)
	log.Println(string(d))

	var v2 time.Time
	json.Unmarshal(d, &v2)
	println(v1.String())
	println(v2.String())
	println(v1 == v2)
	println(v1.Equal(v2))
}

// https://youtu.be/xzaS8gXlzqM?t=1438
