package main

import (
	"log"

	"github.com/ViniciusDJM/jusbrasil-teste/pkg/bootstrap"
)

//	@title			Jusbrasil - Process Finder
//	@version		1.0.0
//	@description	API to helps access process information
//	@contact.name	Jusbrasil
//	@contact.email	example@jusbrasil.com
//	@BasePath		/api
//	@accept			json
func main() {
	app := bootstrap.SetupApp()
	if err := app.Listen(":8000"); err != nil {
		log.Fatalln(err)
	}
}
