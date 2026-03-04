package usecase

import (
	"context"

	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
	"github.com/booscaaa/locksmith/api/internal/project/contract"
)

type deleteProjectUseCase struct {
	deleteProjectRepository contract.DeleteProjectRepository
}

func (u *deleteProjectUseCase) Execute(ctx context.Context, id string) error {
	err := u.deleteProjectRepository.Execute(ctx, id)
	if err != nil {
		return stackerror.NewUseCaseError("DeleteProjectUseCase", err)
	}

	return nil
}

func NewDeleteProjectUseCase(
	deleteProjectRepository contract.DeleteProjectRepository,
) contract.DeleteProjectUseCase {
	return &deleteProjectUseCase{
		deleteProjectRepository: deleteProjectRepository,
	}
}
