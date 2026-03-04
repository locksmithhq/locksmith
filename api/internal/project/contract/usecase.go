package contract

import (
	"context"

	"github.com/booscaaa/go-paginate/v3/paginate"
	"github.com/booscaaa/locksmith/api/internal/project/types/input"
	"github.com/booscaaa/locksmith/api/internal/project/types/output"
)

type CreateProjectUseCase interface {
	Execute(context.Context, input.Project) (output.Project, error)
}

type FetchProjectsUseCase interface {
	Execute(context.Context, paginate.PaginationParams) ([]output.Project, error)
}

type GetProjectByIDUseCase interface {
	Execute(context.Context, string) (output.Project, error)
}

type UpdateProjectUseCase interface {
	Execute(context.Context, string, input.Project) (output.Project, error)
}

type DeleteProjectUseCase interface {
	Execute(context.Context, string) error
}
