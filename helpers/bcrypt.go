package helpers

import "golang.org/x/crypto/bcrypt"

func HashPass(p string) string {
	salt := 8

	hash, _ := bcrypt.GenerateFromPassword([]byte(p), salt)
	return string(hash)
}

func ComparePass(h, p []byte) bool {
	// hash, pass := []byte(h), []byte(p)

	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(p))

	return err == nil
}
