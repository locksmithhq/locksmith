package contract

import (
	"context"

	"github.com/booscaaa/go-paginate/v3/paginate"
	"github.com/locksmithhq/locksmith/api/internal/project/domain"
)

type CreateProjectRepository interface {
	Execute(context.Context, domain.Project) (domain.Project, error)
}

type GetProjectByNameRepository interface {
	Execute(context.Context, string) (domain.Project, error)
}

type FetchProjectsRepository interface {
	Execute(context.Context, paginate.PaginationParams) ([]domain.Project, error)
}

type GetProjectByIDRepository interface {
	Execute(context.Context, string) (domain.Project, error)
}

type UpdateProjectRepository interface {
	Execute(context.Context, string, domain.Project) (domain.Project, error)
}

type DeleteProjectRepository interface {
	Execute(context.Context, string) error
}
