package __

import (
	"log"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	tests := map[string]struct {
		layout, value string
	}{
		"RFC822":     {time.RFC822, "23 Mar 24 21:22 AEDT"},
		"SO75799611": {"2006-01-02T15:04:05Z07:00", "2017-09-04T04:00:00"},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			log.Println(time.Parse(test.layout, test.value))
		})
	}
}
