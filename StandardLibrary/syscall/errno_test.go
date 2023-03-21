package __

import (
	"log"
	"syscall"
	"testing"
)

func TestErrno(t *testing.T) {
	_, err := syscall.Open("qqqqqq", syscall.O_RDONLY, 0777)
	if err != nil {
		log.Printf("%v\n", err) // The system cannot find the file specified
	}
	log.Println(syscall.Errno(23))
}
