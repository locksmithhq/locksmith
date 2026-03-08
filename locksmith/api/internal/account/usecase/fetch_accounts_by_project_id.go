package usecase

import (
	"context"

	"github.com/booscaaa/go-paginate/v3/paginate"
	"github.com/locksmithhq/locksmith/api/internal/account/contract"
	"github.com/locksmithhq/locksmith/api/internal/account/types/output"
)

type fetchAccountsByProjectIDUseCase struct {
	fetchAccountsByProjectIDRepository contract.FetchAccountsByProjectIDRepository
}

// Execute implements contract.FetchAccountsByProjectIDUseCase.
func (u *fetchAccountsByProjectIDUseCase) Execute(ctx context.Context, projectID string, params paginate.PaginationParams) ([]output.Account, error) {
	accounts, err := u.fetchAccountsByProjectIDRepository.Execute(ctx, projectID, params)
	if err != nil {
		return nil, err
	}

	return output.NewAccountsFromDomain(accounts), nil
}

func NewFetchAccountsByProjectIDUseCase(fetchAccountsByProjectIDRepository contract.FetchAccountsByProjectIDRepository) contract.FetchAccountsByProjectIDUseCase {
	return &fetchAccountsByProjectIDUseCase{fetchAccountsByProjectIDRepository: fetchAccountsByProjectIDRepository}
}
