package dataBase

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type user struct {
	id       int
	name     string
	email    string
	password string
}

func InitializeDB() {
	db, err := sql.Open("libsql", "http://127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
		return
	}

	fmt.Println("connected!")
}

func TestDb() {
	db, err := sql.Open("libsql", "http://127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	res, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println("algo se puteo")
	}

	for res.Next() {
		var userRes user
		if err := res.Scan(
			&userRes.id,
			&userRes.name,
			&userRes.email,
			&userRes.password,
		); err != nil {
			log.Fatal("unable to read from db")
		}
		fmt.Printf("Usuario #: %v, nombre: %s, email: %s, pass: %s", userRes.id, userRes.name, userRes.email, userRes.password)
	}
}

func UserLogin(user) {

}
