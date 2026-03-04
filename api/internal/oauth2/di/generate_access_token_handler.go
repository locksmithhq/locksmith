package di

import (
	"github.com/booscaaa/locksmith/api/internal/adapter/database"
	"github.com/booscaaa/locksmith/api/internal/oauth2/contract"
	"github.com/booscaaa/locksmith/api/internal/oauth2/handler"
	"github.com/booscaaa/locksmith/api/internal/oauth2/repository"
	"github.com/booscaaa/locksmith/api/internal/oauth2/usecase"
)

func NewGenerateAccessTokenHandler() contract.GenerateAccessTokenHandler {
	return handler.NewGenerateAccessTokenHandler(
		usecase.NewGenerateAccessTokenUseCase(
			repository.NewGetAuthCodeByCodeRepository(database.GetConnection()),
			repository.NewUpdateAuthCodeRepository(database.GetConnection()),
			repository.NewGetClientByClientIDRepository(database.GetConnection()),
			repository.NewCreateUserSessionRepository(database.GetConnection()),
			repository.NewCreateRefreshTokenRepository(database.GetConnection()),
			repository.NewGetUserSessionByDeviceRepository(database.GetConnection()),
		),
	)
}
