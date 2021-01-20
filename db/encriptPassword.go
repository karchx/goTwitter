package db

import "golang.org/x/crypto/bcrypt"

/* EncriptPassword routine that allows me to encrypt the password  */
func EncriptPassword(pass string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
