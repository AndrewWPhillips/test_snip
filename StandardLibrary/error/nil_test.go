package __

import (
	"log"
	"os"
	"testing"
)

func TestNil(t *testing.T) {
	var err error
	log.Println("<", os.IsNotExist(err), ">")
}
