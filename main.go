package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/BenChapman/mobile_api/project"
	"github.com/BenChapman/mobile_api/project_sql_repository"
	"github.com/gorilla/mux"
)

var (
	p project.Project
)

func main() {
	db, err := sqlx.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@tcp("+os.Getenv("DB_HOST")+")/"+os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatal(err)
	}

	repo := &project_sql_repository.ProjectSQLRepository{
		Sql: db,
	}

	p = project.Project{
		ProjectRepository: repo,
	}

	r := mux.NewRouter()
	r.HandleFunc("/project", handleListProjects)
	r.HandleFunc("/project/{id}", handleGetProject)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}

func handleListProjects(w http.ResponseWriter, r *http.Request) {
	projects, _ := p.ListProjects()
	json.NewEncoder(w).Encode(projects)
}

func handleGetProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	project, _ := p.GetProject(id)

	json.NewEncoder(w).Encode(project)
}
