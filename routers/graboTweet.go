package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/KenethSandoval/goTwitter/db"
	"github.com/KenethSandoval/goTwitter/models"
)

/*GraboTweet permite grbar el tweet en la DB*/
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := db.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar publicar el tweet "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se logro insertar el tweet ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
