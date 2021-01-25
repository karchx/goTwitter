package routers

import (
	"encoding/json"
	"net/http"

	"github.com/KenethSandoval/goTwitter/db"
	"github.com/KenethSandoval/goTwitter/models"
)

/* Registry it is the function to believe the user record in the database */
func Registry(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar un password de al menos 6 caracters", 400)
		return
	}

	_, encontrado, _ := db.ChequeoYaExisteUsuario(t.Email)

	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return
	}

	_, status, err := db.InsertRegistry(t)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro de Usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
