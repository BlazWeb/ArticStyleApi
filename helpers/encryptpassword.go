package helpers

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(pass string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(bytes), err
}

func EncryptHash(s string) (string, error) {
	cost := 1
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), cost)
	return string(bytes), err
}
