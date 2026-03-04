package contract

import (
	"context"

	"github.com/booscaaa/go-paginate/v3/paginate"
	"github.com/booscaaa/locksmith/api/internal/acl/domain"
)

type FetchAclRepository interface {
	Execute(context.Context, paginate.PaginationParams) ([]domain.Acl, error)
}

type CreateActionRepository interface {
	Execute(context.Context, domain.Action) error
}

type CreateModuleRepository interface {
	Execute(context.Context, domain.Module) error
}

type CreateRoleRepository interface {
	Execute(context.Context, domain.Role) error
}

type CreateProjectAclRepository interface {
	Execute(context.Context, domain.ProjectAcl) error
}

type FetchRolesRepository interface {
	Execute(context.Context) ([]domain.Role, error)
}

type FetchModulesRepository interface {
	Execute(context.Context) ([]domain.Module, error)
}

type FetchActionsRepository interface {
	Execute(context.Context) ([]domain.Action, error)
}

type FetchProjectAclRepository interface {
	Execute(context.Context, string) ([]domain.ProjectAcl, error)
}

type DeleteProjectAclByProjectIdRepository interface {
	Execute(context.Context, string) error
}

type GetProjectDomainByProjectIDRepository interface {
	Execute(context.Context, string) (string, error)
}
