package db

import (
	"github.com/KenethSandoval/goTwitter/models"
	"golang.org/x/crypto/bcrypt"
)

func IntentoLogin(email string, password string) (models.User, bool) {
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)
	if encontrado == false {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(usu.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)

	if err != nil {
		return usu, false
	}

	return usu, true
}
