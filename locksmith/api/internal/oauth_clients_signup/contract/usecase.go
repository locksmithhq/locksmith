package contract

import (
	"context"

	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_signup/domain"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_signup/types/input"
)

type GetSignupByClientIDUseCase interface {
	Execute(context.Context, string) (domain.Signup, error)
}

type CreateSignupByClientIDUseCase interface {
	Execute(context.Context, string, input.Signup) error
}

type UpdateSignupByClientIDUseCase interface {
	Execute(context.Context, string, input.Signup) error
}
