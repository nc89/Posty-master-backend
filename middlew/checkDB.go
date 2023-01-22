package middlew

import (
	"net/http"

	"github.com/NicolasCardenas/Posty-master/db"
)

/*CheckDB es el middlew que me permite conocer el estado de la BD */
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "Conexión perdida con la Base de Datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
