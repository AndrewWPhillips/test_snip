package crypto

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestBCrypt(t *testing.T) {
	bcrypt.CompareHashAndPassword(nil, nil)
}

func CheckPasswordNotCracked(hash []byte) bool {
	var CrackedPasswords = []string{"abc", "def"}
	for _, pwd := range CrackedPasswords {
		if err := bcrypt.CompareHashAndPassword(hash, []byte(pwd)); err == nil {
			return true
		}
	}
	return false
}
