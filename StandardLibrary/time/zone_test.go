package __

import (
	"log"
	"strconv"
	"testing"
	"time"
)

func TestFakeTimeZone(t *testing.T) {
	offset := "-1100"

	tm, err := time.Parse("15:04 GMT-0700", "15:06 GMT"+offset)
	if err != nil {
		log.Println("fail", err)
	}

	log.Println(tm)
	log.Println(tm.UTC())

	if len(offset) == 5 {
		hours, ok1 := strconv.ParseInt(offset[:3], 10, 0)
		mins, ok2 := strconv.ParseInt(offset[3:5], 10, 0)
		if ok1 == nil && ok2 == nil {
			tm = tm.In(time.FixedZone("Fake", int((hours*60+mins)*60)))
			log.Println(tm)
			log.Println(tm.Location())
		}
	}

}
