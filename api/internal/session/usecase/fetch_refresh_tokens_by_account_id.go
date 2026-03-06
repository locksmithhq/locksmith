package usecase

import (
	"context"

	"github.com/locksmithhq/locksmith/api/internal/session/contract"
	"github.com/locksmithhq/locksmith/api/internal/session/types/output"
)

type fetchRefreshTokensByAccountIDUseCase struct {
	fetchRefreshTokensByAccountIDRepository contract.FetchRefreshTokensByAccountIDRepository
}

func (u *fetchRefreshTokensByAccountIDUseCase) Execute(ctx context.Context, projectID, accountID, sessionID string, page, limit int) ([]output.RefreshToken, error) {
	tokens, err := u.fetchRefreshTokensByAccountIDRepository.Execute(ctx, projectID, accountID, sessionID, page, limit)
	if err != nil {
		return nil, err
	}
	return output.NewRefreshTokensFromDomain(tokens), nil
}

func NewFetchRefreshTokensByAccountIDUseCase(
	fetchRefreshTokensByAccountIDRepository contract.FetchRefreshTokensByAccountIDRepository,
) contract.FetchRefreshTokensByAccountIDUseCase {
	return &fetchRefreshTokensByAccountIDUseCase{
		fetchRefreshTokensByAccountIDRepository: fetchRefreshTokensByAccountIDRepository,
	}
}
