package output

import "github.com/booscaaa/locksmith/api/internal/project/domain"

type Project struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Domain      string `json:"domain"`
}

func NewProject(project domain.Project) Project {
	return Project{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
		Domain:      project.Domain,
	}
}

func NewProjects(projects []domain.Project) []Project {
	var output []Project
	for _, project := range projects {
		output = append(output, NewProject(project))
	}
	return output
}
