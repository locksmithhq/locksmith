package contract

import (
	"context"

	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_login/domain"
)

type GetLoginByClientIDRepository interface {
	Execute(context.Context, string) (domain.Login, error)
}

type CreateLoginByClientIDRepository interface {
	Execute(context.Context, domain.Login) error
}

type UpdateLoginByClientIDRepository interface {
	Execute(context.Context, domain.Login) error
}
