package contract

import (
	"context"

	"github.com/booscaaa/go-paginate/v3/paginate"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/domain"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/types/input"
)

type CreateClientRepository interface {
	Execute(context.Context, input.Client) (domain.Client, error)
}

type GetClientByProjectIDAndClientIDRepository interface {
	Execute(context.Context, string, string) (domain.Client, error)
}

type FetchClientsByProjectIDRepository interface {
	Execute(context.Context, string, paginate.PaginationParams) ([]domain.Client, error)
}

type GetClientByIDAndProjectIDRepository interface {
	Execute(context.Context, string, string) (domain.Client, error)
}

type UpdateClientRepository interface {
	Execute(context.Context, string, string, domain.Client) (domain.Client, error)
}

type DeleteClientRepository interface {
	Execute(context.Context, string, string) error
}

type GetClientByCustomDomainRepository interface {
	Execute(ctx context.Context, domain string) (clientID string, redirectURI string, err error)
}

type GetClientByClientIDRepository interface {
	Execute(ctx context.Context, clientID string) (domain.Client, error)
}
