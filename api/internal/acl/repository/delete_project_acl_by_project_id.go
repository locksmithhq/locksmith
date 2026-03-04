package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/acl/contract"
)

type deleteProjectAclByProjectIdRepository struct {
	database types.Database
}

func (r *deleteProjectAclByProjectIdRepository) Execute(ctx context.Context, projectId string) error {
	query := `DELETE FROM project_acl WHERE project_id = $1`
	_, err := r.database.ExecContext(ctx, query, projectId)
	return err
}

func NewDeleteProjectAclByProjectIdRepository(database types.Database) contract.DeleteProjectAclByProjectIdRepository {
	return &deleteProjectAclByProjectIdRepository{database: database}
}
