package dataBase

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

// var dbUrl = "http://127.0.0.1:8080"
// db, err := sql.Open("libsql", dbUrl)
//
//	if err != nil {
//	    fmt.Fprintf(os.Stderr, "failed to open db %s: %s", dbUrl, err)
//	    os.Exit(1)
//	}

type user struct {
	id       int
	name     string
	email    string
	password string
}

func InitializeDB() {
	log.Print("entró a inicializar")
	db, err := sql.Open("libsql", "http://127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	_, err = db.Query("SELECT 1")
	if err != nil {
		fmt.Println("Hizo el query")
	}

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

	fmt.Println("connected!")
}

func TestDb() {
	db, err := sql.Open("libsql", "http://127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("connected!")
	_, err = db.Query("SELECT 1")
	if err != nil {
		fmt.Println("Algo se daño")
	}

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
