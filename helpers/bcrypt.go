package helpers

import "golang.org/x/crypto/bcrypt"

func HashPass(password string) string {
	salt := 8

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), salt)
	return string(hash)
}

func ComparePass(hashPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}
