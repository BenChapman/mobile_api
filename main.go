package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/BenChapman/mobile_api/project"
	"github.com/BenChapman/mobile_api/project_sql_repository"
	"github.com/gorilla/mux"
)

var (
	p    project.Project
	repo *project_sql_repository.ProjectSQLRepository
)

func main() {
	repo := new(project_sql_repository.ProjectSQLRepository)

	p = project.Project{
		ProjectRepository: repo,
	}

	r := mux.NewRouter()
	r.HandleFunc("/project", handleListProjects)
	r.HandleFunc("/project/{id}", handleGetProject)
	http.ListenAndServe(":3000", r)
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
