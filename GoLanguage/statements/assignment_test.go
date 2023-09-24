package __

import (
	"log"
	"testing"
)

// TestMultAssignWithField test mult. assigning including to struct field
func TestMultAssignWithField(t *testing.T) {
	f := func() (int, error) { return 3, nil }
	i := 2
	i, err := f()
	log.Println(i, err)

	var err2 error
	a := struct{ b int }{2}
	a.b, err2 = f()
	log.Println(a, err2)
}
