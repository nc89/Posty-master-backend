package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/NicolasCardenas/Posty-master/db"
	"github.com/NicolasCardenas/Posty-master/models"
)

/*GraboPost permite grabar el post en la base de datos */
func GraboPost(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Post
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboPost{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := db.InsertoPost(registro)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el registro, reintente nuevamente"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el Post", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
