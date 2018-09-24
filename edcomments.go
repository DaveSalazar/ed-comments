package main

import (
	"flag"
	"log"

	"github.com/DaveSalazar/ed-comments/migration"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migracion a la base de datos")
	flag.Parse()
	if migrate == "yes" {
		log.Println("comenzo la migracion....")
		migration.Migrate()
		log.Println("finalizo la migracion")
	}
}
