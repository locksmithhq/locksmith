package usecase

import (
	"context"

	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/project/contract"
	"github.com/locksmithhq/locksmith/api/internal/project/types/input"
	"github.com/locksmithhq/locksmith/api/internal/project/types/output"
)

type updateProjectUseCase struct {
	updateProjectRepository contract.UpdateProjectRepository
}

// Execute implements contract.UpdateProjectUseCase.
func (usecase *updateProjectUseCase) Execute(ctx context.Context, id string, in input.Project) (output.Project, error) {
	project, err := usecase.updateProjectRepository.Execute(ctx, id, in.ToProjectDomain())
	if err != nil {
		return output.Project{}, stackerror.NewUseCaseError("UpdateProjectUseCase", err)
	}

	return output.NewProject(project), nil

}

func NewUpdateProjectUseCase(
	updateProjectRepository contract.UpdateProjectRepository,
) contract.UpdateProjectUseCase {
	return &updateProjectUseCase{
		updateProjectRepository: updateProjectRepository,
	}
}
