package usecase

import (
	"context"

	"github.com/booscaaa/locksmith/api/internal/acl/contract"
	"github.com/booscaaa/locksmith/api/internal/acl/types/input"
)

type createModuleUseCase struct {
	createModuleRepository contract.CreateModuleRepository
}

// Execute implements contract.CreateRoleUseCase.
func (u *createModuleUseCase) Execute(ctx context.Context, in input.Module) error {
	return u.createModuleRepository.Execute(ctx, in.ToModuleDomain())
}

func NewCreateModuleUseCase(createModuleRepository contract.CreateModuleRepository) contract.CreateModuleUseCase {
	return &createModuleUseCase{createModuleRepository: createModuleRepository}
}
