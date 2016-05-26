package project_sql_repository

import (
	"log"

	"github.com/BenChapman/mobile_api/project"
	"github.com/jmoiron/sqlx"
)

type ProjectSQLRepository struct {
	Sql *sqlx.DB
}

func (p *ProjectSQLRepository) Get(id int) (project.ProjectDetails, error) {
	projectDetails := project.ProjectDetails{}

	u := p.Sql.Unsafe()
	err := u.Get(&projectDetails, "SELECT * FROM `project` p WHERE p.`id`=?", id)

	if err != nil {
		log.Fatal(err)
	}

	return projectDetails, nil
}

func (p *ProjectSQLRepository) GetAll() ([]project.ProjectDetails, error) {
	return []project.ProjectDetails{{
		Id:          1,
		Name:        "Test Project",
		Description: "This is a really cool test project that is used for testing things.",
		DojoUuid:    "98756a48-d3f5-4951-aef5-2433627a7cbe",
		Category:    "Mobile",
	}}, nil
}
