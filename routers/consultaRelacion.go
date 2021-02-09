package routers

import (
	"encoding/json"
	"net/http"

	"github.com/KenethSandoval/goTwitter/db"
	"github.com/KenethSandoval/goTwitter/models"
)

/*ConsultaRelacion consulta la relacion que hay entre 2 usuarios*/
func ConsultaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionId = ID

	var resp models.RespuestaConsultaRelacion

	status, err := db.ConsultoRelacion(t)
	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
