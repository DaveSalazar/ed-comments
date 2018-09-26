package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/urfave/negroni"

	"github.com/DaveSalazar/ed-comments/commons"
	"github.com/DaveSalazar/ed-comments/migration"
	"github.com/DaveSalazar/ed-comments/routes"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migracion a la base de datos")
	flag.IntVar(&commons.Port, "port", 8080, "Puerto para el servidor web")
	flag.Parse()
	if migrate == "yes" {
		log.Println("comenzo la migracion....")
		migration.Migrate()
		log.Println("finalizo la migracion")
	}

	//inicia las rutas
	router := routes.InitRoutes()

	//iniciar los middlewares
	n := negroni.Classic()
	n.UseHandler(router)

	// inicia el servidor

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", commons.Port),
		Handler: n,
	}

	log.Printf("iniciado el servidor en http://localhost:%d", commons.Port)

	log.Println(server.ListenAndServe())

	log.Println("Fin de la ejecucion")
}
