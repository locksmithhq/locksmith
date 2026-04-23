package di

import (
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/repository"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/usecase"
)

func NewLogoutUseCase() contract.LogoutUseCase {
	conn := database.GetConnection()
	return usecase.NewLogoutUseCase(
		repository.NewGetRefreshTokenByHashRepository(conn),
		repository.NewUpdateRefreshTokenRepository(conn),
		repository.NewRevokeUserSessionRepository(conn),
	)
}
