package di

import (
	"github.com/locksmithhq/locksmith/api/internal/locksmith/contract"
	"github.com/locksmithhq/locksmith/api/internal/locksmith/handler"
	oauth2Di "github.com/locksmithhq/locksmith/api/internal/oauth2/di"
)

func NewCallbackHandler() contract.CallbackHandler {
	return handler.NewCallbackHandler(
		oauth2Di.NewGenerateAccessTokenUseCase(),
	)
}
