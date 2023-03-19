package __

import (
	"log"
	"strings"
	"testing"

	"github.com/go-yaml/yaml"
)

func TestYAML(t *testing.T) {
	var d struct {
		A int
		B string
	}
	s :=
		`
---
a: 42
b: s
`
	if err := yaml.NewDecoder(strings.NewReader(s)).Decode(&d); err != nil {
		t.Fatal("Error decoding YAML:", err.Error())
	}

	log.Println(d)
}
