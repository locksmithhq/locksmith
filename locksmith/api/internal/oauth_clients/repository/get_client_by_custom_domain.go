package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/contract"
)

type getClientByCustomDomainRepository struct {
	database types.Database
}

// Execute implements contract.GetClientByCustomDomainRepository.
func (r *getClientByCustomDomainRepository) Execute(ctx context.Context, domain string) (string, error) {
	var clientID string

	query := `
		SELECT client_id FROM oauth_clients
		WHERE custom_domain = $1
		LIMIT 1
	`
	err := r.database.GetContext(ctx, &clientID, query, domain)
	if err != nil {
		return "", stackerror.NewRepositoryError("GetClientByCustomDomainRepository", err)
	}

	return clientID, nil
}

func NewGetClientByCustomDomainRepository(database types.Database) contract.GetClientByCustomDomainRepository {
	return &getClientByCustomDomainRepository{
		database: database,
	}
}
