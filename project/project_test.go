package project_test

import (
	"github.com/BenChapman/mobile_api/project"
	"github.com/BenChapman/mobile_api/project/projectfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Project", func() {
	var (
		fakeRepo       *projectfakes.FakeProjectRepository
		projectHandler project.Project
	)

	BeforeEach(func() {
		fakeRepo = new(projectfakes.FakeProjectRepository)

		fakeRepo.GetStub = func(id int) (project.ProjectDetails, error) {
			return project.ProjectDetails{
				Id:          id,
				Name:        "Test Project",
				Description: "This is a really cool test project that is used for testing things.",
				DojoUuid:    "98756a48-d3f5-4951-aef5-2433627a7cbe",
				Category:    "Mobile",
			}, nil
		}
		fakeRepo.GetAllStub = func() ([]project.ProjectDetails, error) {
			return []project.ProjectDetails{{
				Id:          1,
				Name:        "Test Project",
				Description: "This is a really cool test project that is used for testing things.",
				DojoUuid:    "98756a48-d3f5-4951-aef5-2433627a7cbe",
				Category:    "Mobile",
			}}, nil
		}

		projectHandler = project.Project{
			ProjectRepository: fakeRepo,
		}
	})

	Describe("ListProjects", func() {
		It("returns a list of projects", func() {
			projectDetails, err := projectHandler.ListProjects()
			Expect(err).ToNot(HaveOccurred())
			Expect(projectDetails).To(Equal([]project.ProjectDetails{{
				Id:          1,
				Name:        "Test Project",
				Description: "This is a really cool test project that is used for testing things.",
				DojoUuid:    "98756a48-d3f5-4951-aef5-2433627a7cbe",
				Category:    "Mobile",
			}}))
		})
	})

	Describe("GetProject", func() {
		It("returns a project", func() {
			projectDetails, err := projectHandler.GetProject(1)
			Expect(err).ToNot(HaveOccurred())
			Expect(projectDetails).To(Equal(project.ProjectDetails{
				Id:          1,
				Name:        "Test Project",
				Description: "This is a really cool test project that is used for testing things.",
				DojoUuid:    "98756a48-d3f5-4951-aef5-2433627a7cbe",
				Category:    "Mobile",
			}))
		})
	})
})
