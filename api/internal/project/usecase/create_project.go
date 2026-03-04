package usecase

import (
	"context"
	"errors"

	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
	"github.com/booscaaa/locksmith/api/internal/project/contract"
	"github.com/booscaaa/locksmith/api/internal/project/types/input"
	"github.com/booscaaa/locksmith/api/internal/project/types/output"
)

type createProjectUseCase struct {
	createProjectRepository    contract.CreateProjectRepository
	getProjectByNameRepository contract.GetProjectByNameRepository
}

// Execute implements contract.CreateProjectUseCase.
func (usecase *createProjectUseCase) Execute(ctx context.Context, in input.Project) (output.Project, error) {
	project, err := usecase.getProjectByNameRepository.Execute(ctx, in.Name)
	if err != nil && !errors.Is(err, stackerror.ErrNotFound) {
		return output.Project{}, stackerror.NewUseCaseError("CreateProjectUseCase", err)
	}

	if project.ID != "" {
		return output.NewProject(project), nil
	}

	project, err = usecase.createProjectRepository.Execute(ctx, in.ToProjectDomain())
	if err != nil {
		return output.Project{}, stackerror.NewUseCaseError("CreateProjectUseCase", err)
	}

	return output.NewProject(project), nil

}

func NewCreateProjectUseCase(
	createProjectRepository contract.CreateProjectRepository,
	getProjectByNameRepository contract.GetProjectByNameRepository,
) contract.CreateProjectUseCase {
	return &createProjectUseCase{
		createProjectRepository:    createProjectRepository,
		getProjectByNameRepository: getProjectByNameRepository,
	}
}
