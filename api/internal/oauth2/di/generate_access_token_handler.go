package di

import (
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/handler"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/repository"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/usecase"
)

func NewGenerateAccessTokenHandler() contract.GenerateAccessTokenHandler {
	return handler.NewGenerateAccessTokenHandler(
		NewGenerateAccessTokenUseCase(),
	)
}

func NewGenerateAccessTokenUseCase() contract.GenerateAccessTokenUseCase {
	return usecase.NewGenerateAccessTokenUseCase(
		repository.NewGetAuthCodeByCodeRepository(database.GetConnection()),
		repository.NewUpdateAuthCodeRepository(database.GetConnection()),
		repository.NewGetClientByClientIDRepository(database.GetConnection()),
		repository.NewCreateUserSessionRepository(database.GetConnection()),
		repository.NewCreateRefreshTokenRepository(database.GetConnection()),
		repository.NewGetUserSessionByDeviceRepository(database.GetConnection()),
		repository.NewUpdateUserSessionActivityRepository(database.GetConnection()),
	)
}
