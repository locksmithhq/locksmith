package usecase

import (
	"context"

	"github.com/locksmithhq/locksmith/api/internal/acl/contract"
	"github.com/locksmithhq/locksmith/api/internal/acl/domain"
)

type fetchActionsUseCase struct {
	fetchActionsRepository contract.FetchActionsRepository
}

func (u *fetchActionsUseCase) Execute(ctx context.Context) ([]domain.Action, error) {
	return u.fetchActionsRepository.Execute(ctx)
}

func NewFetchActionsUseCase(fetchActionsRepository contract.FetchActionsRepository) contract.FetchActionsUseCase {
	return &fetchActionsUseCase{fetchActionsRepository: fetchActionsRepository}
}
