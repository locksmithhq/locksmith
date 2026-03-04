package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/booscaaa/locksmith/api/internal/acl/contract"
	"github.com/booscaaa/locksmith/api/internal/acl/domain"
	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
)

type fetchProjectAclRepository struct {
	database types.Database
}

func (r *fetchProjectAclRepository) Execute(ctx context.Context, projectId string) ([]domain.ProjectAcl, error) {
	var projectAcls []domain.ProjectAcl
	query := `
		SELECT 
			project_acl.id, 
			project_acl.role_id, 
			roles.title as role_name, 
			project_acl.module_id, 
			modules.title as module_name, 
			project_acl.action_id, 
			actions.title as action_name, 
			project_acl.project_id 
		FROM project_acl 
		INNER JOIN roles ON project_acl.role_id = roles.id
		INNER JOIN modules ON project_acl.module_id = modules.id
		INNER JOIN actions ON project_acl.action_id = actions.id
		WHERE project_acl.project_id = $1
	`
	err := r.database.SelectContext(ctx, &projectAcls, query, projectId)
	if err != nil {
		return nil, stackerror.NewRepositoryError("FetchProjectAclRepository", err)
	}
	return projectAcls, nil
}

func NewFetchProjectAclRepository(database types.Database) contract.FetchProjectAclRepository {
	return &fetchProjectAclRepository{
		database: database,
	}
}
