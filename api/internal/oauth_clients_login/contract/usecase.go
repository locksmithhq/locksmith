package contract

import (
	"context"

	"github.com/booscaaa/locksmith/api/internal/oauth_clients_login/domain"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients_login/types/input"
)

type GetLoginByClientIDUseCase interface {
	Execute(context.Context, string) (domain.Login, error)
}

type CreateLoginByClientIDUseCase interface {
	Execute(context.Context, string, input.Login) error
}

type UpdateLoginByClientIDUseCase interface {
	Execute(context.Context, string, input.Login) error
}
