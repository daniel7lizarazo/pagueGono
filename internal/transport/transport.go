package transport

import (
	"net/http"
	"text/template"

	dataBase "github.com/daniel7lizarazo/pagueGono/internal/database"
)

func AddHandlers() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/js"))))
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/loginPage", loginPage)
	http.HandleFunc("/login", login)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles("./web/index.html"))
	templ.Execute(w, nil)
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles("./web/login.html"))
	templ.Execute(w, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	// userInfo := UserLoginInfo{
	// 	user:     r.PostFormValue("user"),
	// 	password: r.PostFormValue("password"),
	// }
	dataBase.TestDb()
}
