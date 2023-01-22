package main

import (
	"log"

	"github.com/NicolasCardenas/Posty-master/db"
	"github.com/NicolasCardenas/Posty-master/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("Sin conexion a la DB")
		return
	}
	handlers.Manejadores()
}
