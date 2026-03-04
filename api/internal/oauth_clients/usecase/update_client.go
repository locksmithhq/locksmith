package usecase

import (
	"context"

	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/types/input"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/types/output"
)

type updateClientUseCase struct {
	updateClientRepository contract.UpdateClientRepository
}

// Execute implements [contract.UpdateClientUseCase].
func (u *updateClientUseCase) Execute(ctx context.Context, projectID string, clientID string, in input.Client) (output.Client, error) {
	client, err := u.updateClientRepository.Execute(ctx, projectID, clientID, in.ToDomain())
	if err != nil {
		return output.Client{}, err
	}

	return output.NewClientFromDomain(client), nil
}

func NewUpdateClientUseCase(updateClientRepository contract.UpdateClientRepository) contract.UpdateClientUseCase {
	return &updateClientUseCase{updateClientRepository: updateClientRepository}
}
