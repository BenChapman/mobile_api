package project_sql_repository

import (
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
		return project.ProjectDetails{}, err
	}

	return projectDetails, nil
}

func (p *ProjectSQLRepository) GetAll() ([]project.ProjectDetails, error) {
	projectDetails := []project.ProjectDetails{}

	u := p.Sql.Unsafe()
	err := u.Select(&projectDetails, "SELECT * FROM `project` p WHERE p.`project_confirmed` != 'No' or p.`project_confirmed` IS NULL")
	if err != nil {
		return []project.ProjectDetails{}, err
	}

	return projectDetails, nil
}
