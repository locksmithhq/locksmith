package usecase

import (
	"context"

	"github.com/locksmithhq/locksmith/api/internal/session/contract"
)

type countSessionsByAccountIDUseCase struct {
	countSessionsByAccountIDRepository contract.CountSessionsByAccountIDRepository
}

func (u *countSessionsByAccountIDUseCase) Execute(ctx context.Context, projectID, accountID string) (int64, error) {
	return u.countSessionsByAccountIDRepository.Execute(ctx, projectID, accountID)
}

func NewCountSessionsByAccountIDUseCase(countSessionsByAccountIDRepository contract.CountSessionsByAccountIDRepository) contract.CountSessionsByAccountIDUseCase {
	return &countSessionsByAccountIDUseCase{countSessionsByAccountIDRepository: countSessionsByAccountIDRepository}
}
