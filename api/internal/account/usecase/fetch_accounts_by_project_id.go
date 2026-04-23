package usecase

import (
	"context"

	"github.com/booscaaa/go-paginate/v3/paginate"
	"github.com/locksmithhq/locksmith/api/internal/account/contract"
	"github.com/locksmithhq/locksmith/api/internal/account/types/output"
)

type fetchAccountsByProjectIDUseCase struct {
	fetchAccountsByProjectIDRepository         contract.FetchAccountsByProjectIDRepository
	fetchSocialProvidersByAccountIDsRepository contract.FetchSocialProvidersByAccountIDsRepository
}

// Execute implements contract.FetchAccountsByProjectIDUseCase.
func (u *fetchAccountsByProjectIDUseCase) Execute(ctx context.Context, projectID string, params paginate.PaginationParams) ([]output.Account, error) {
	accounts, err := u.fetchAccountsByProjectIDRepository.Execute(ctx, projectID, params)
	if err != nil {
		return nil, err
	}

	ids := make([]string, len(accounts))
	for i, a := range accounts {
		ids[i] = a.ID
	}

	socialMap, _ := u.fetchSocialProvidersByAccountIDsRepository.Execute(ctx, ids)
	for i := range accounts {
		if providers, ok := socialMap[accounts[i].ID]; ok {
			accounts[i].SocialProviders = providers
		} else {
			accounts[i].SocialProviders = []string{}
		}
	}

	return output.NewAccountsFromDomain(accounts), nil
}

func NewFetchAccountsByProjectIDUseCase(
	fetchAccountsByProjectIDRepository contract.FetchAccountsByProjectIDRepository,
	fetchSocialProvidersByAccountIDsRepository contract.FetchSocialProvidersByAccountIDsRepository,
) contract.FetchAccountsByProjectIDUseCase {
	return &fetchAccountsByProjectIDUseCase{
		fetchAccountsByProjectIDRepository:         fetchAccountsByProjectIDRepository,
		fetchSocialProvidersByAccountIDsRepository: fetchSocialProvidersByAccountIDsRepository,
	}
}
