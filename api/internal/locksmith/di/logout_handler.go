package di

import (
	"github.com/locksmithhq/locksmith/api/internal/locksmith/contract"
	"github.com/locksmithhq/locksmith/api/internal/locksmith/handler"
	oauth2Di "github.com/locksmithhq/locksmith/api/internal/oauth2/di"
)

func NewLogoutHandler() contract.LogoutHandler {
	return handler.NewLogoutHandler(
		oauth2Di.NewLogoutUseCase(),
	)
}
