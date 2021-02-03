package routers

import (
	"net/http"

	"github.com/KenethSandoval/goTwitter/db"
	"github.com/KenethSandoval/goTwitter/models"
)

/*AltaRelacion graba una alta relacion en la base de datos*/
func AltaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe venir el parametro id", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionId = ID

	status, err := db.InsertoRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrio un error "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar la relacion", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
