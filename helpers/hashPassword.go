package helpers

import "golang.org/x/crypto/bcrypt"

// Hashing password
func HashPassword(password string) string{
	bytes := []byte(password)
	hash, _ := bcrypt.GenerateFromPassword(bytes, 14)

	return string(hash)
}
