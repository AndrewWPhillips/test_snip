package __

import (
	"log"
	"os"
	"testing"
)

func TestCreateFile(t *testing.T) {
	filename := `.\t.bin`
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalln(err)
	}
	f.Write([]byte("ABCDEFG"))
	f.Close()

	//if f, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666); err == nil {
	//	f.Write([]byte("XYZ"))  // leaves DEFG at end of file since O_TRUNC was not specified
	//	f.Close()
	//}
	f, err = os.Create(filename)
	if err != nil {
		log.Fatalln(err)
	}
	f.Write([]byte("XYZ"))
	f.Close()
}
