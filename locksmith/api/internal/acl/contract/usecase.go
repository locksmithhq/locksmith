package contract

import (
	"context"

	"github.com/booscaaa/go-paginate/v3/paginate"
	"github.com/locksmithhq/locksmith/api/internal/acl/domain"
	"github.com/locksmithhq/locksmith/api/internal/acl/types/input"
	"github.com/locksmithhq/locksmith/api/internal/acl/types/output"
)

type FetchAclUseCase interface {
	Execute(context.Context, paginate.PaginationParams) ([]domain.Acl, error)
}

type CreateRoleUseCase interface {
	Execute(context.Context, input.Role) error
}

type CreateModuleUseCase interface {
	Execute(context.Context, input.Module) error
}

type CreateActionUseCase interface {
	Execute(context.Context, input.Action) error
}

type FetchRolesUseCase interface {
	Execute(context.Context) ([]domain.Role, error)
}

type FetchModulesUseCase interface {
	Execute(context.Context) ([]domain.Module, error)
}

type FetchActionsUseCase interface {
	Execute(context.Context) ([]domain.Action, error)
}

type CreateProjectAclUseCase interface {
	Execute(context.Context, string, input.ProjectAcl) error
}

type FetchProjectAclUseCase interface {
	Execute(context.Context, string) (output.ProjectAcl, error)
}
