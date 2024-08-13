package lib

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "github.com/lib/pq"
)

// TODO: Put creds in environment
const (
    host = "localhost"
    port = 5432
    user = "postgres"
    password = "lucas"
    dbname = "postgres"
)

type DB struct{
    *sql.DB
}

func Dbconf() (*DB, error) {
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
    if err != nil { log.Fatalf("Error on reading init.sql: %s", err)
    }

    _, err = db.Exec(string(sqlStmt))
    if err !=nil {log.Fatalf("err exec: %s", err)
    }

    fmt.Println("Database Tables initialized")

    return &DB{db}, nil

}

func getAllProjects(db *DB) ([]Project, error) {
    rows, err := db.Query("SELECT * FROM %s")
    if err != nil {
        log.Panicf("Error on query %s", err)
    }
    defer rows.Close()

    var projects []Project
    for rows.Next() {
        var project Project

        err := rows.Scan(&project.project_id, &project.name, &project.description, &project.services);
        projects = append(projects, project)
            return nil, err

    }
    return projects, nil
}
