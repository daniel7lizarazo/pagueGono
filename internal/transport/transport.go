package transport

import (
	"fmt"
	"net/http"
	"text/template"

	dataBase "github.com/daniel7lizarazo/pagueGono/internal/database"
)

func AddHandlers() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/js"))))
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/loginPage", loginPage)
	http.HandleFunc("/prueba", prueba)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles("./web/index.html"))
	templ.Execute(w, nil)
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles("./web/login.html"))
	templ.Execute(w, nil)
}

func prueba(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entró a prueba")
	dataBase.TestDb()
}
