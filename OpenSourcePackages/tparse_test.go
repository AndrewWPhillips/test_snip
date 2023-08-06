package __

import (
	"log"
	"testing"
	"time"

	"github.com/karrick/tparse"
)

func TestTParse(t *testing.T) {
	t1, err := tparse.ParseNow("2006-01-02T15:04:05", "2023-04-05T06:07:08")
	if err != nil {
		t.Fatal(err.Error())
	}
	log.Printf("%v\n", t1)
}

func TestParseNow(t *testing.T) {
	t1 := time.Now()
	t2, err := tparse.ParseNow(time.RFC822, "now+1h+100m")
	if err != nil {
		t.Fatal(err.Error())
	}
	log.Printf("Now: %v, 2h40m Later: %v\n", t1, t2)
}

func TestParseWithMap(t *testing.T) {
	t1 := time.Now()
	t2, err := tparse.ParseWithMap(time.RFC822, "later+1h",
		map[string]time.Time{
			"soon":  time.Now().Add(time.Minute),
			"later": time.Now().Add(2 * time.Hour),
		})
	if err != nil {
		t.Fatal(err.Error())
	}
	log.Printf("Now: %v, 3h Later: %v\n", t1, t2)
}
