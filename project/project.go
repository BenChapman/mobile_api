package project

type ProjectDetails struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	DojoUuid    string `json:"dojo_uuid"`
	Category    string `json:"category"`
}

type ProjectRepository interface {
	Get(int) (ProjectDetails, error)
	GetAll() ([]ProjectDetails, error)
}

type Project struct {
	ProjectRepository ProjectRepository
}

func (p *Project) ListProjects() ([]ProjectDetails, error) {
	return p.ProjectRepository.GetAll()
}

func (p *Project) GetProject(id int) (ProjectDetails, error) {
	return p.ProjectRepository.Get(id)
}
