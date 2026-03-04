package usecase

import (
	"context"

	"github.com/locksmithhq/locksmith/api/internal/session/contract"
)

type countSessionsByProjectIDUseCase struct {
	countSessionsByProjectIDRepository contract.CountSessionsByProjectIDRepository
}

func (u *countSessionsByProjectIDUseCase) Execute(ctx context.Context, projectID string, search string) (int64, error) {
	return u.countSessionsByProjectIDRepository.Execute(ctx, projectID, search)
}

func NewCountSessionsByProjectIDUseCase(countSessionsByProjectIDRepository contract.CountSessionsByProjectIDRepository) contract.CountSessionsByProjectIDUseCase {
	return &countSessionsByProjectIDUseCase{countSessionsByProjectIDRepository: countSessionsByProjectIDRepository}
}
