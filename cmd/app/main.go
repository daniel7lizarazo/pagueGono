package main

import (
	"fmt"
	"log"

	"github.com/daniel7lizarazo/pagueGono/internal/app"
	"github.com/daniel7lizarazo/pagueGono/internal/transport"
)

func main() {
	transport.AddHandlers()
	app.InitializeApp()
	log.Print("Corriendo from logger")
	fmt.Println("Corriendo from fmt")
}
