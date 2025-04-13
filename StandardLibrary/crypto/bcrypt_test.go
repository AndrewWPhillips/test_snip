package crypto

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestBCrypt(t *testing.T) {
	bcrypt.CompareHashAndPassword(nil, nil)
}

func CheckPasswordNotCracked(hash []byte) bool {
	for pwd := range CrackedPasswords {
		if err := bcrypt.CompareHashAndPassword(hash, pwd); err == nil {
			return true
		}
	}
	return false
}
