package app

import (
	"log"
	"net/http"
)

func InitializeApp() {
	initializeWebServer()
}

func initializeWebServer() {
	log.Fatal(http.ListenAndServe("localhost:8081", nil))
}
