package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients/contract"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients/domain"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients/types/input"
)

type createClientRepository struct {
	database types.Database
}

// Execute implements contract.CreateClientRepository.
func (r *createClientRepository) Execute(ctx context.Context, input input.Client) (domain.Client, error) {
	var client domain.Client

	query := `INSERT INTO oauth_clients (
		project_id, name, client_id, client_secret, redirect_uris, grant_types, custom_domain
	) VALUES ($1, $2, $3, $4, $5, $6, $7) 
	RETURNING id, project_id, client_id, client_secret, redirect_uris, grant_types, name, created_at, updated_at, custom_domain`

	err := r.database.QueryRowxContext(ctx,
		query,
		input.ProjectID,
		input.Name,
		input.ClientID,
		input.ClientSecret,
		input.RedirectURIs,
		input.GrantTypes,
		input.CustomDomain,
	).StructScan(&client)

	if err != nil {
		return domain.Client{}, stackerror.NewRepositoryError("CreateClientRepository", err)
	}
	return client, nil
}

func NewCreateClientRepository(database types.Database) contract.CreateClientRepository {
	return &createClientRepository{database: database}
}
