package contract

import (
	"context"

	"github.com/booscaaa/go-paginate/v3/paginate"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/types/input"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/types/output"
)

type CreateClientUseCase interface {
	Execute(context.Context, input.Client) (output.Client, error)
}

type FetchClientsByProjectIDUseCase interface {
	Execute(context.Context, string, paginate.PaginationParams) ([]output.Client, error)
}

type GetClientByIDAndProjectIDUseCase interface {
	Execute(context.Context, string, string) (output.Client, error)
}

type UpdateClientUseCase interface {
	Execute(context.Context, string, string, input.Client) (output.Client, error)
}

type DeleteClientUseCase interface {
	Execute(context.Context, string, string) error
}

type ResolveCustomDomainUseCase interface {
	Execute(ctx context.Context, hostname string) (string, error)
}
