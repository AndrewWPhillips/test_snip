package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"log"
	"testing"
)

func TestAES(t *testing.T) {
	blockCipher16, err := aes.NewCipher([]byte("1234567812345678"))
	if err != nil {
		t.Fatalf("Error creating AES cipher %v\n", err)
	}
	log.Printf("Block size %v\n", blockCipher16.BlockSize())

	blockCipher32, err2 := aes.NewCipher([]byte("12345678123456781234567812345678"))
	if err2 != nil {
		t.Fatalf("Error creating AES cipher %v\n", err2)
	}
	log.Printf("Block size %v\n", blockCipher32.BlockSize())

	stream := cipher.NewCTR(blockCipher16, []byte("1234567812345678"))
	log.Printf("%T\n", stream)
}
