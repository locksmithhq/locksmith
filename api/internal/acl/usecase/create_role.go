package usecase

import (
	"context"

	"github.com/locksmithhq/locksmith/api/internal/acl/contract"
	"github.com/locksmithhq/locksmith/api/internal/acl/types/input"
)

type createRoleUseCase struct {
	createRoleRepository contract.CreateRoleRepository
}

// Execute implements contract.CreateRoleUseCase.
func (u *createRoleUseCase) Execute(ctx context.Context, in input.Role) error {
	return u.createRoleRepository.Execute(ctx, in.ToRoleDomain())
}

func NewCreateRoleUseCase(createRoleRepository contract.CreateRoleRepository) contract.CreateRoleUseCase {
	return &createRoleUseCase{createRoleRepository: createRoleRepository}
}
