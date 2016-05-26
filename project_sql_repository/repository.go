package project_sql_repository

import "github.com/BenChapman/mobile_api/project"

//import _ "github.com/go-sql-driver/mysql"

type ProjectSQLRepository struct{}

func (p *ProjectSQLRepository) Get(id int) (project.ProjectDetails, error) {
	return project.ProjectDetails{
		Id:          id,
		Name:        "Test Project",
		Description: "This is a really cool test project that is used for testing things.",
		DojoUuid:    "98756a48-d3f5-4951-aef5-2433627a7cbe",
		Category:    "Mobile",
	}, nil
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
