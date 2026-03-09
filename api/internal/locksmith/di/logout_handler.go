package di

import (
	"github.com/locksmithhq/locksmith/api/internal/locksmith/contract"
	"github.com/locksmithhq/locksmith/api/internal/locksmith/handler"
)

func NewLogoutHandler() contract.LogoutHandler {
	return handler.NewLogoutHandler()
}
