package usecase

import (
	"context"

	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
	"github.com/booscaaa/locksmith/api/internal/project/contract"
	"github.com/booscaaa/locksmith/api/internal/project/types/output"
)

type getProjectByIDUseCase struct {
	getProjectByIDRepository contract.GetProjectByIDRepository
}

// Execute implements contract.GetProjectByIDUseCase.
func (u *getProjectByIDUseCase) Execute(ctx context.Context, id string) (output.Project, error) {
	project, err := u.getProjectByIDRepository.Execute(ctx, id)
	if err != nil {
		return output.Project{}, stackerror.NewUseCaseError("GetProjectByIDUseCase", err)
	}
	return output.NewProject(project), nil
}

func NewGetProjectByIDUseCase(getProjectByIDRepository contract.GetProjectByIDRepository) contract.GetProjectByIDUseCase {
	return &getProjectByIDUseCase{
		getProjectByIDRepository: getProjectByIDRepository,
	}
}
