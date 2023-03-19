package __

import (
	"log"
	"testing"
)

type hh struct{ ii int }

func (p *hh) cc() int { return 42 }

type dd struct {
	hh
	ee float64
}

func (p *dd) cc() int { return 73 }

func TestOverride(t *testing.T) {
	var d1 dd
	log.Println(d1.cc()) // overrides hh version
	log.Println(d1.hh.cc())
}
