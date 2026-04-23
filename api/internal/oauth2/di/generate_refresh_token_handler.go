package di

import (
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/handler"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/repository"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/usecase"
)

func NewGenerateRefreshTokenHandler() contract.GenerateRefreshTokenHandler {
	return handler.NewGenerateRefreshTokenHandler(
		NewGenerateRefreshTokenUseCase(),
	)
}

func NewGenerateRefreshTokenUseCase() contract.GenerateRefreshTokenUseCase {
	conn := database.GetConnection()
	return usecase.NewGenerateRefreshTokenUseCase(
		repository.NewGetRefreshTokenByHashRepository(conn),
		repository.NewUpdateRefreshTokenRepository(conn),
		repository.NewCreateRefreshTokenRepository(conn),
		repository.NewGetClientByClientIDRepository(conn),
		repository.NewGetClientByIDRepository(conn),
		repository.NewRevokeRefreshTokensBySessionRepository(conn),
		repository.NewRevokeUserSessionRepository(conn),
	)
}
