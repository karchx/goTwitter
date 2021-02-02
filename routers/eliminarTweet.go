package routers

import (
	"net/http"

	"github.com/KenethSandoval/goTwitter/db"
)

/*EliminarTweet permite borrar un Tweet determinado*/
func EliminarTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Deber enviar el parametro ID", http.StatusBadRequest)
		return
	}
	err := db.BorroTweet(ID, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar borrar el tweet", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
