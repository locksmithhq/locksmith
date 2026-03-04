package di

import (
	"github.com/booscaaa/locksmith/api/internal/adapter/database"
	"github.com/booscaaa/locksmith/api/internal/oauth2/contract"
	"github.com/booscaaa/locksmith/api/internal/oauth2/handler"
	"github.com/booscaaa/locksmith/api/internal/oauth2/repository"
	"github.com/booscaaa/locksmith/api/internal/oauth2/usecase"
)

func NewGenerateRefreshTokenHandler() contract.GenerateRefreshTokenHandler {
	return handler.NewGenerateRefreshTokenHandler(
		usecase.NewGenerateRefreshTokenUseCase(
			repository.NewGetRefreshTokenByHashRepository(database.GetConnection()),
			repository.NewUpdateRefreshTokenRepository(database.GetConnection()),
			repository.NewCreateRefreshTokenRepository(database.GetConnection()),
			repository.NewGetClientByClientIDRepository(database.GetConnection()),
			repository.NewGetClientByIDRepository(database.GetConnection()),
		),
	)
}
