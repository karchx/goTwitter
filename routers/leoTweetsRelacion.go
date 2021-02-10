package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/KenethSandoval/goTwitter/db"
)

/*LeoTweetsUsuarios lee los tweets de todos los seguidores*/
func LeoTweetsUsuarios(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("pagine")) < 1 {
		http.Error(w, "Debe enviar el parámetro página", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))

	if err != nil {
		http.Error(w, "Debe enviar el parametro pagina como entero mayor a 0", http.StatusBadRequest)
		return
	}

	respuesta, correcto := db.LeoTweetsSeguidores(IDUsuario, pagina)
	if correcto == false {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
