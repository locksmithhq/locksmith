package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/locksmithhq/locksmith/api/internal/project/contract"
	"github.com/locksmithhq/locksmith/api/internal/project/types/input"
	"github.com/locksmithhq/locksmith/api/internal/project/types/output"
	"gopkg.in/yaml.v2"
)

type configDefaultProjectCMD struct {
	createProjectUseCase contract.CreateProjectUseCase
}

func (c *configDefaultProjectCMD) Execute(ctx context.Context) (output.Project, error) {
	filePath := "/etc/locksmith/config/seeder.yaml"
	data, err := os.ReadFile(filePath)
	if err != nil {
		return output.Project{}, fmt.Errorf("failed to read seeder config file %s: %w", filePath, err)
	}

	var seederConfig struct {
		DefaultProject struct {
			Name        string `yaml:"name"`
			Description string `yaml:"description"`
			Domain      string `yaml:"domain"`
		} `yaml:"default_project"`
	}

	if err := yaml.Unmarshal(data, &seederConfig); err != nil {
		return output.Project{}, fmt.Errorf("failed to unmarshal seeder config from %s: %w", filePath, err)
	}

	in := input.Project{
		Name:        seederConfig.DefaultProject.Name,
		Description: seederConfig.DefaultProject.Description,
		Domain:      seederConfig.DefaultProject.Domain,
	}

	project, err := c.createProjectUseCase.Execute(ctx, in)
	if err != nil {
		return output.Project{}, err
	}

	return project, nil
}

func NewConfigDefaultProject(createProjectUseCase contract.CreateProjectUseCase) contract.ConfigDefaultProjectCMD {
	return &configDefaultProjectCMD{createProjectUseCase: createProjectUseCase}
}
