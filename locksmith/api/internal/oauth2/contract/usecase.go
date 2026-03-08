package contract

import (
	"context"

	"github.com/locksmithhq/locksmith/api/internal/oauth2/types/input"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/types/output"
)


type AuthorizeClient interface {
	Execute(context.Context, input.Authorization) (output.Client, error)
}

type GenerateAccessTokenUseCase interface {
	Execute(context.Context, input.AccessToken) (output.AccessToken, error)
}

type GenerateRefreshTokenUseCase interface {
	Execute(context.Context, string) (output.AccessToken, error)
}

type GetUserInfoUseCase interface {
	Execute(context.Context, int) (map[string]interface{}, error)
}

type LoginUseCase interface {
	Execute(context.Context, input.Login) (output.Login, error)
}

type CheckTokenStatusUseCase interface {
	Execute(context.Context, string) error
}

type RegisterUseCase interface {
	Execute(context.Context, input.Register) (output.Login, error)
}
