package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/domain"
)

type getClientByClientIDRepository struct {
	database types.Database
}

func (r *getClientByClientIDRepository) Execute(ctx context.Context, clientID string) (domain.Client, error) {
	var client domain.Client

	query := `
		SELECT id, project_id, client_secret
		FROM oauth_clients
		WHERE client_id = $1
	`
	err := r.database.GetContext(ctx, &client, query, clientID)
	if err != nil {
		return domain.Client{}, stackerror.NewRepositoryError("GetClientByClientIDRepository", err)
	}

	return client, nil
}

func NewGetClientByClientIDRepository(database types.Database) contract.GetClientByClientIDRepository {
	return &getClientByClientIDRepository{
		database: database,
	}
}
