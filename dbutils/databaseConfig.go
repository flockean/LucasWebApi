package dbutils

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

// TODO: Put creds in environment
const (
	host     = "0.0.0.0"
	port     = 5432
	user     = "postgres"
	password = "rocket123"
	dbname   = "projectDb"
)

type DB struct {
	*sql.DB
}

func ConnectDatabase() (*DB, error) {
	// TODO: password needs to be... well not visible
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	fmt.Println(psqlInfo)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("DB connected")

	sqlStmt, err := os.ReadFile("./db/init.sql")
	if err != nil {
		log.Fatalf("Error on reading init.sql: %s", err)
	}

	_, err = db.Exec(string(sqlStmt))
	if err != nil {
		log.Fatalf("err exec: %s", err)
	}

	fmt.Println("Database Tables initialized")

	return &DB{db}, nil

}
