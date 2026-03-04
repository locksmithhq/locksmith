package usecase

import (
	"context"

	"github.com/booscaaa/go-paginate/v3/paginate"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients/contract"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients/types/output"
)

type fetchClientsByProjectIDUseCase struct {
	fetchClientsByProjectIDRepository contract.FetchClientsByProjectIDRepository
}

// Execute implements contract.FetchClientsByProjectIDUseCase.
func (u *fetchClientsByProjectIDUseCase) Execute(
	ctx context.Context,
	projectID string,
	params paginate.PaginationParams,
) ([]output.Client, error) {
	clients, err := u.fetchClientsByProjectIDRepository.Execute(ctx, projectID, params)
	if err != nil {
		return nil, err
	}

	return output.NewClientsFromDomain(clients), nil
}

func NewFetchClientsByProjectIDUseCase(
	fetchClientsByProjectIDRepository contract.FetchClientsByProjectIDRepository,
) contract.FetchClientsByProjectIDUseCase {
	return &fetchClientsByProjectIDUseCase{
		fetchClientsByProjectIDRepository: fetchClientsByProjectIDRepository,
	}
}
