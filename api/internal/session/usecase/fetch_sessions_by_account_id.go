package usecase

import (
	"context"

	"github.com/locksmithhq/locksmith/api/internal/session/contract"
	"github.com/locksmithhq/locksmith/api/internal/session/types/output"
)

type fetchSessionsByAccountIDUseCase struct {
	fetchSessionsByAccountIDRepository contract.FetchSessionsByAccountIDRepository
}

func (u *fetchSessionsByAccountIDUseCase) Execute(ctx context.Context, projectID, accountID string, page, limit int) ([]output.Session, error) {
	sessions, err := u.fetchSessionsByAccountIDRepository.Execute(ctx, projectID, accountID, page, limit)
	if err != nil {
		return nil, err
	}
	return output.NewSessionsFromDomain(sessions), nil
}

func NewFetchSessionsByAccountIDUseCase(fetchSessionsByAccountIDRepository contract.FetchSessionsByAccountIDRepository) contract.FetchSessionsByAccountIDUseCase {
	return &fetchSessionsByAccountIDUseCase{fetchSessionsByAccountIDRepository: fetchSessionsByAccountIDRepository}
}
