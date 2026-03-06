package usecase

import (
	"context"

	"github.com/locksmithhq/locksmith/api/internal/session/contract"
)

type countRefreshTokensByAccountIDUseCase struct {
	countRefreshTokensByAccountIDRepository contract.CountRefreshTokensByAccountIDRepository
}

func (u *countRefreshTokensByAccountIDUseCase) Execute(ctx context.Context, projectID, accountID, sessionID string) (int64, error) {
	return u.countRefreshTokensByAccountIDRepository.Execute(ctx, projectID, accountID, sessionID)
}

func NewCountRefreshTokensByAccountIDUseCase(
	countRefreshTokensByAccountIDRepository contract.CountRefreshTokensByAccountIDRepository,
) contract.CountRefreshTokensByAccountIDUseCase {
	return &countRefreshTokensByAccountIDUseCase{
		countRefreshTokensByAccountIDRepository: countRefreshTokensByAccountIDRepository,
	}
}
