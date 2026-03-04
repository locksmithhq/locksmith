package contract

import (
	"context"

	"github.com/locksmithhq/locksmith/api/internal/oauth2/domain"
)

type GetClientByClientIDRepository interface {
	Execute(context.Context, string) (domain.Client, error)
}

type GetClientByIDRepository interface {
	Execute(context.Context, string) (domain.Client, error)
}

type CreateAuthCodeRepository interface {
	Execute(context.Context, domain.AuthCode) (domain.AuthCode, error)
}

type GetAuthCodeByCodeRepository interface {
	Execute(context.Context, string) (domain.AuthCode, error)
}

type UpdateAuthCodeRepository interface {
	Execute(context.Context, domain.AuthCode) error
}

type CreateUserSessionRepository interface {
	Execute(context.Context, domain.UserSession) (domain.UserSession, error)
}

type GetUserSessionByDeviceRepository interface {
	Execute(context.Context, string, string, string) (domain.UserSession, error)
}

type CreateRefreshTokenRepository interface {
	Execute(context.Context, domain.RefreshToken) (domain.RefreshToken, error)
}

type GetRefreshTokenByHashRepository interface {
	Execute(context.Context, string) (domain.RefreshToken, error)
}

type UpdateRefreshTokenRepository interface {
	Execute(context.Context, domain.RefreshToken) error
}

type GetAccountByEmailPasswordAndProjectIDRepository interface {
	Execute(context.Context, string, string, string) (domain.Account, error)
}

type GetLoginByClientIDRepository interface {
	Execute(context.Context, string) (domain.OAuthClientLogin, error)
}

type GetSignupByClientIDRepository interface {
	Execute(context.Context, string) (domain.OAuthClientSignup, error)
}
