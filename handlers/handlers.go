package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/NicolasCardenas/Posty-master/middlew"
	"github.com/NicolasCardenas/Posty-master/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*seteo del puerto, el handler y poner a escuchar el servidor*/
func Manejadores() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/registro", middlew.CheckDB(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.CheckDB(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlew.CheckDB(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/Post", middlew.CheckDB(middlew.ValidoJWT(routers.GraboPost))).Methods("POST")
	router.HandleFunc("/leoPosts", middlew.CheckDB(middlew.ValidoJWT(routers.LeoPosts))).Methods("GET")
	router.HandleFunc("/eliminarPost", middlew.CheckDB(middlew.ValidoJWT(routers.EliminarPost))).Methods("DELETE")

	router.HandleFunc("/subirAvatar", middlew.CheckDB(middlew.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middlew.CheckDB(routers.ObtenerAvatar)).Methods("GET")
	router.HandleFunc("/subirBanner", middlew.CheckDB(middlew.ValidoJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerBanner", middlew.CheckDB(routers.ObtenerBanner)).Methods("GET")

	router.HandleFunc("/altaRelacion", middlew.CheckDB(middlew.ValidoJWT(routers.AltaRelacion))).Methods("POST")
	router.HandleFunc("/bajaRelacion", middlew.CheckDB(middlew.ValidoJWT(routers.BajaRelacion))).Methods("DELETE")
	router.HandleFunc("/consultaRelacion", middlew.CheckDB(middlew.ValidoJWT(routers.ConsultaRelacion))).Methods("GET")

	router.HandleFunc("/listaUsuarios", middlew.CheckDB(middlew.ValidoJWT(routers.ListaUsuarios))).Methods("GET")
	router.HandleFunc("/leoPostsSeguidores", middlew.CheckDB(middlew.ValidoJWT(routers.LeoPostsSeguidores))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
