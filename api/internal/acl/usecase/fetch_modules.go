package usecase

import (
	"context"

	"github.com/booscaaa/locksmith/api/internal/acl/contract"
	"github.com/booscaaa/locksmith/api/internal/acl/domain"
)

type fetchModulesUseCase struct {
	fetchModulesRepository contract.FetchModulesRepository
}

func (u *fetchModulesUseCase) Execute(ctx context.Context) ([]domain.Module, error) {
	return u.fetchModulesRepository.Execute(ctx)
}

func NewFetchModulesUseCase(fetchModulesRepository contract.FetchModulesRepository) contract.FetchModulesUseCase {
	return &fetchModulesUseCase{fetchModulesRepository: fetchModulesRepository}
}
