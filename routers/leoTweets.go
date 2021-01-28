package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/KenethSandoval/goTwitter/db"
)

/*LeoTweets lee los tweets*/
func LeoTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parametro página", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))

	if err != nil {
		http.Error(w, "Debe enviar el parametro página con un valor mayor a 0", http.StatusBadRequest)
		return
	}

	pag := int64(pagina)
	respuesta, correcto := db.LeoTweets(ID, pag)
	if correcto == false {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(respuesta)
}
