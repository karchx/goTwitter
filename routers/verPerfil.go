package routers

import (
	"encoding/json"
	"net/http"

	"github.com/KenethSandoval/goTwitter/db"
)

/*VerPerfil permite extraer los valores del Perfil*/
func VerPerfil(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	perfil, err := db.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Ocurrion un error al intentar buscar el registro"+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}
