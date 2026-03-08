package usecase

import (
	"context"

	"github.com/locksmithhq/locksmith/api/internal/session/contract"
	"github.com/locksmithhq/locksmith/api/internal/session/types/output"
)

type fetchSessionsByProjectIDUseCase struct {
	fetchSessionsByProjectIDRepository contract.FetchSessionsByProjectIDRepository
}

func (u *fetchSessionsByProjectIDUseCase) Execute(ctx context.Context, projectID string, page, limit int, search string) ([]output.Session, error) {
	sessions, err := u.fetchSessionsByProjectIDRepository.Execute(ctx, projectID, page, limit, search)
	if err != nil {
		return nil, err
	}

	return output.NewSessionsFromDomain(sessions), nil
}

func NewFetchSessionsByProjectIDUseCase(fetchSessionsByProjectIDRepository contract.FetchSessionsByProjectIDRepository) contract.FetchSessionsByProjectIDUseCase {
	return &fetchSessionsByProjectIDUseCase{fetchSessionsByProjectIDRepository: fetchSessionsByProjectIDRepository}
}
