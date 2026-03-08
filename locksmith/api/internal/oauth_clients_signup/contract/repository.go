package contract

import (
	"context"

	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_signup/domain"
)

type GetSignupByClientIDRepository interface {
	Execute(context.Context, string) (domain.Signup, error)
}

type CreateSignupByClientIDRepository interface {
	Execute(context.Context, domain.Signup) error
}

type UpdateSignupByClientIDRepository interface {
	Execute(context.Context, domain.Signup) error
}
