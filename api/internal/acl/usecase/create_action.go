package usecase

import (
	"context"

	"github.com/locksmithhq/locksmith/api/internal/acl/contract"
	"github.com/locksmithhq/locksmith/api/internal/acl/types/input"
)

type createActionUseCase struct {
	createActionRepository contract.CreateActionRepository
}

// Execute implements contract.CreateActionUseCase.
func (u *createActionUseCase) Execute(ctx context.Context, in input.Action) error {
	return u.createActionRepository.Execute(ctx, in.ToActionDomain())
}

func NewCreateActionUseCase(createActionRepository contract.CreateActionRepository) contract.CreateActionUseCase {
	return &createActionUseCase{createActionRepository: createActionRepository}
}
