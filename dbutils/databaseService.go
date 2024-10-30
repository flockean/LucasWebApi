package dbutils

import (
	"fmt"
	"log"
)

func getAllProjects(db *DB) ([]Project, error) {
	rows, err := db.Query("SELECT * FROM project")
	if err != nil {
		log.Panicf("Error on query %s", err)
	}
	defer rows.Close()

	var projects []Project
	for rows.Next() {
		var project Project

		err := rows.Scan(&project.id, &project.name, &project.description)
		projects = append(projects, project)
		return nil, err

	}
	return projects, nil
}

func CreateProject(db *DB, project Project) (Project, error) {
	result, err := db.Exec("INSERT INTO project VALUES ($1, $2, $3)",
		project.id, project.name, project.description)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Executed: %s", result)
	return project, err
}

func CreateService(db *DB, service Service) (Service, error) {
	result, err := db.Exec("INSERT INTO service VALUES ($1, $2, $3, $4, $5)",
		service.id, service.name, service.lang, service.focus, service.project)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Executed: %s", result)
	return service, err
}
