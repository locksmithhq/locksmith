package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/booscaaa/locksmith/api/internal/acl/contract"
	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
)

type getProjectDomainByProjectIDRepository struct {
	database types.Database
}

// Execute implements contract.GetProjectDomainByProjectIDRepository.
func (r *getProjectDomainByProjectIDRepository) Execute(ctx context.Context, projectID string) (string, error) {
	var domain string
	query := `
		SELECT 
			domain 
		FROM projects 
		WHERE id = $1
	`
	err := r.database.GetContext(ctx, &domain, query, projectID)
	if err != nil {
		return "", stackerror.NewRepositoryError("GetProjectDomainByProjectIDRepository", err)
	}

	return domain, nil
}

func NewGetProjectDomainByProjectIDRepository(database types.Database) contract.GetProjectDomainByProjectIDRepository {
	return &getProjectDomainByProjectIDRepository{
		database: database,
	}
}
