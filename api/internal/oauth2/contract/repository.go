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

type UpdateUserSessionActivityRepository interface {
	Execute(ctx context.Context, sessionID string) error
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

type GetSocialProviderByClientRepository interface {
	Execute(ctx context.Context, clientID string, provider string) (domain.SocialProvider, error)
}

type GetEnabledSocialProvidersByClientRepository interface {
	Execute(ctx context.Context, clientID string) ([]string, error)
}

type CreateSocialStateRepository interface {
	Execute(ctx context.Context, state domain.SocialState) error
}

type GetSocialStateByNonceRepository interface {
	Execute(ctx context.Context, nonce string) (domain.SocialState, error)
}

type DeleteSocialStateRepository interface {
	Execute(ctx context.Context, nonce string) error
}

type GetAccountBySocialProviderRepository interface {
	Execute(ctx context.Context, provider string, providerUserID string) (domain.Account, error)
}

type GetAccountByEmailAndProjectRepository interface {
	Execute(ctx context.Context, email string, projectID string) (domain.Account, error)
}

type CreateAccountSocialProviderRepository interface {
	Execute(ctx context.Context, asp domain.AccountSocialProvider) error
}

type CreateAccountRepository interface {
	Execute(ctx context.Context, account domain.Account) (domain.Account, error)
}

type RevokeUserSessionRepository interface {
	Execute(ctx context.Context, sessionID string) error
}

type RevokeRefreshTokensBySessionRepository interface {
	Execute(ctx context.Context, sessionID string) error
}
