package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/acl/contract"
	"github.com/locksmithhq/locksmith/api/internal/acl/domain"
)

type createProjectAclRepository struct {
	database types.Database
}

// Execute implements contract.CreateActionRepository.
func (r *createProjectAclRepository) Execute(ctx context.Context, projectAcl domain.ProjectAcl) error {
	query := `INSERT INTO project_acl (
		role_id,
		module_id,
		action_id,
		project_id
	) VALUES ($1, $2, $3, $4)`
	_, err := r.database.ExecContext(ctx,
		query,
		projectAcl.RoleId,
		projectAcl.ModuleId,
		projectAcl.ActionId,
		projectAcl.ProjectId,
	)
	return err
}

func NewCreateProjectAclRepository(database types.Database) contract.CreateProjectAclRepository {
	return &createProjectAclRepository{database: database}
}
