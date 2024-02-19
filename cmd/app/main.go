package main

import (
	"fmt"
	"log"

	"github.com/daniel7lizarazo/pagueGono/internal/app"
	dataBase "github.com/daniel7lizarazo/pagueGono/internal/dataBase"
	"github.com/daniel7lizarazo/pagueGono/internal/transport"
)

func main() {
	dataBase.InitializeDB()
	transport.AddHandlers()
	app.InitializeApp()
	log.Print("Corriendo from logger")
	fmt.Println("Corriendo from fmt")
}
