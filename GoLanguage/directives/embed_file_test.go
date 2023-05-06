package __

import (
	"embed"
	"log"
	"testing"
)

// s is a string embedding the contents of the file test.dat
//
//go:embed testdata/test.dat
var s string

func TestEmbedString(t *testing.T) {
	log.Println(s)
}

// bb is a byte slice embedding the contents of test.dat
//
//go:embed testdata/test.dat
var bb []byte

func TestEmbedBytes(t *testing.T) {
	log.Println(string(bb))
}

// fs is an embedded "filesystem" with the file(s) from testdata and a source (.go) file
//
//go:embed testdata embed_file_test.go
var fs embed.FS

func TestEmbedFS(t *testing.T) {
	d, err := fs.ReadFile("testdata/test.dat")
	if err != nil {
		t.Fatal(err)
	}
	log.Println(string(d))

	d, err = fs.ReadFile("embed_file_test.go")
	if err != nil {
		t.Fatal(err)
	}
	log.Println(string(d)[:10])

	d, err = fs.ReadFile("testdata/non-existent")
	if err != nil {
		t.Fatal(err)
	}
	log.Println(string(d))
}

// Note that repeats (as below) are effective ignored - ie. test.dat is only stored once

//go:embed testdata/test.dat testdata/test.dat
var fs2 embed.FS

func TestEmbedTwice(t *testing.T) {
	de, err := fs2.ReadDir(".")
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range de {
		log.Println(entry.Name(), entry.Type())
	}

	// Even though the same file was added twice it was only embedded once
	de, err = fs2.ReadDir("testdata")
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range de {
		log.Println(entry.Name(), entry.Type())
	}
}
