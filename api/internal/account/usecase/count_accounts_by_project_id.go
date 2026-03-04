package usecase

import (
	"context"

	"github.com/booscaaa/go-paginate/v3/paginate"
	"github.com/booscaaa/locksmith/api/internal/account/contract"
)

type countAccountsByProjectIDUseCase struct {
	countAccountsByProjectIDRepository contract.CountAccountsByProjectIDRepository
}

// Execute implements contract.CountAccountsByProjectIDUseCase.
func (u *countAccountsByProjectIDUseCase) Execute(ctx context.Context, projectID string, params paginate.PaginationParams) (int64, error) {
	return u.countAccountsByProjectIDRepository.Execute(ctx, projectID, params)
}

func NewCountAccountsByProjectIDUseCase(countAccountsByProjectIDRepository contract.CountAccountsByProjectIDRepository) contract.CountAccountsByProjectIDUseCase {
	return &countAccountsByProjectIDUseCase{countAccountsByProjectIDRepository: countAccountsByProjectIDRepository}
}
