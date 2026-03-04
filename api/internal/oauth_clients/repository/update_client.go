package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients/contract"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients/domain"
)

type updateClientRepository struct {
	database types.Database
}

// Execute implements contract.UpdateClientRepository.
func (r *updateClientRepository) Execute(ctx context.Context, projectID string, id string, in domain.Client) (domain.Client, error) {
	var client domain.Client

	query := `UPDATE oauth_clients
		SET name = $1, client_secret = $2, redirect_uris = $3, grant_types = $4, custom_domain = $5
		WHERE project_id = $6 AND id = $7
		RETURNING id, project_id, client_id, client_secret, redirect_uris, grant_types, name, created_at, updated_at, custom_domain`

	err := r.database.QueryRowxContext(ctx,
		query,
		in.Name,
		in.ClientSecret,
		in.RedirectURIs,
		in.GrantTypes,
		in.CustomDomain,
		projectID,
		id,
	).StructScan(&client)

	if err != nil {
		return domain.Client{}, stackerror.NewRepositoryError("UpdateClientRepository", err)
	}
	return client, nil
}

func NewUpdateClientRepository(database types.Database) contract.UpdateClientRepository {
	return &updateClientRepository{database: database}
}
