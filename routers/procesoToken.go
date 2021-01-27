package routers

import (
	"errors"
	"strings"

	"github.com/KenethSandoval/goTwitter/db"
	"github.com/KenethSandoval/goTwitter/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*Email valor de Email usado en todos los EndPoints*/
var Email string

/*IDUsuario es el ID devuelto del modelo, que se usara en todos los EndPoints*/
var IDUsuario string

/*ProcesoToken proces token para extraer sus valores*/
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("ecinuesecret112019$")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, _ := db.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}

		return claims, encontrado, IDUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}
	return claims, false, string(""), err
}
