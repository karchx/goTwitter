package jwt

import (
	"time"

	"github.com/KenethSandoval/goTwitter/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/* GeneroJWT genera el encriptado con JWT */
func GeneroJWT(t models.User) (string, error) {
	miClave := []byte("ecinuesecret112019$")

	payload := jwt.MapClaims{
		"email":     t.Email,
		"name":      t.Name,
		"lastname":  t.LastName,
		"birthdate": t.Birthdate,
		"biography": t.Biography,
		"location":  t.Location,
		"website":   t.WebSite,
		"_id":       t.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
