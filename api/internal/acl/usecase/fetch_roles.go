package usecase

import (
	"context"

	"github.com/booscaaa/locksmith/api/internal/acl/contract"
	"github.com/booscaaa/locksmith/api/internal/acl/domain"
)

type fetchRolesUseCase struct {
	fetchRolesRepository contract.FetchRolesRepository
}

func (u *fetchRolesUseCase) Execute(ctx context.Context) ([]domain.Role, error) {
	return u.fetchRolesRepository.Execute(ctx)
}

func NewFetchRolesUseCase(fetchRolesRepository contract.FetchRolesRepository) contract.FetchRolesUseCase {
	return &fetchRolesUseCase{fetchRolesRepository: fetchRolesRepository}
}
