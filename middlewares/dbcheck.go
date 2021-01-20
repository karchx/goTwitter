package middleware

import (
	"net/http"

	"github.com/KenethSandoval/goTwitter/db"
)

/* DBCheck it is the middleware that allows me to know the state of the database */
func DBCheck(nex http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.ConnectionCheck() == 0 {
			http.Error(w, "Conexión perdida con la base de datos", 500)
			return
		}
		next.ServeHttp(w, r)
	}
}
