package transport

import (
	"fmt"
	"net/http"
	"text/template"

	dataBase "github.com/daniel7lizarazo/pagueGono/internal/dataBase"
	"github.com/daniel7lizarazo/pagueGono/internal/models"
)

func AddHandlers() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/loginPage", loginPage)
	http.HandleFunc("/login", login)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles("./web/html/index.html"))
	templ.Execute(w, nil)
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles("./web/html/login.html"))
	templ.Execute(w, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	userInfo := models.UserLoginInfo{
		Email:    r.PostFormValue("email"),
		Password: r.PostFormValue("password"),
	}
	fmt.Println(userInfo)
	dataBase.TestDb()
}
